package main

// 符号
type Token struct {
	TokenType string // 符号类型
	TokenText string // 符号文本
	Offset    int    // 位置
}

type ParserFunc = func(s *Scanner) []Token

// 是否为字母
func IsAlpha(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// 是否为数字
func IsDigit(c rune) bool {
	return c >= '0' && c <= '9'
}

// 是否为标点符号
func IsPunctuation(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/' || c == '=' || c == '_' || c == '.'
}

// 是否为空白字符
func IsWhite(c rune) bool {
	return c == ' ' || c == '\n' || c == '\r' || c == '\t'
}

// 字符串是否在数组内
func isInStringArray(arr []string, s string) bool {
	for _, a := range arr {
		if s == a {
			return true
		}
	}
	return false
}

// 解析空白字符
func ParseWhite(s *Scanner) []Token {
	chars := make([]rune, 0)
	offset := s.Offset()
	c := s.Current()
	if !IsWhite(c) {
		return nil
	}
	chars = append(chars, c)
	for {
		c := s.Next()
		if IsWhite(c) {
			chars = append(chars, c)
		} else {
			break
		}
	}
	return []Token{{TokenType: "White", TokenText: string(chars), Offset: offset}}
}

// 关键词列表
var Keywords = []string{"int", "float", "bool"}

// 解析关键词
func ParseKeyword(s *Scanner) []Token {
	tokenList := ParseID(s)
	if len(tokenList) != 1 {
		return nil
	}
	if !isInStringArray(Keywords, tokenList[0].TokenText) {
		return nil
	}
	tokenList[0].TokenType = "Keyword"
	return tokenList
}

// 解析标志符
func ParseID(s *Scanner) []Token {
	chars := make([]rune, 0)
	offset := s.Offset()
	c := s.Current()
	if !IsAlpha(c) {
		return nil
	}
	chars = append(chars, c)
	for {
		c := s.Next()
		if IsAlpha(c) || IsDigit(c) || c == '_' {
			chars = append(chars, c)
		} else {
			break
		}
	}
	return []Token{{TokenType: "ID", TokenText: string(chars), Offset: offset}}
}

// 解析整数字面量
func ParseIntLiteral(s *Scanner) []Token {
	chars := make([]rune, 0)
	offset := s.Offset()
	c := s.Current()
	if !IsDigit(c) {
		return nil
	}
	chars = append(chars, c)
	for {
		c := s.Next()
		if IsDigit(c) {
			chars = append(chars, c)
		} else {
			break
		}
	}
	return []Token{{TokenType: "IntLiteral", TokenText: string(chars), Offset: offset}}
}

// 解析标点符号
func ParsePunctuation(s *Scanner) []Token {
	c := s.Current()
	if !IsPunctuation(c) {
		return nil
	}
	s.Next()
	return []Token{{TokenType: "Punctuation", TokenText: string(c), Offset: s.Offset()}}
}

// 尝试顺序执行解析器
func And(parsers ...ParserFunc) ParserFunc {
	return func(s *Scanner) []Token {
		result := make([]Token, 0)
		for _, parser := range parsers {
			tokenList := parser(s)
			if tokenList == nil {
				return nil
			}
			result = append(result, tokenList...)
		}
		return result
	}
}

// 尝试顺序执行解析器，每个解析器中间加入可选的空白字符解析
func AndNoWhiteToken(parsers ...ParserFunc) ParserFunc {
	return func(s *Scanner) []Token {
		newParsers := make([]ParserFunc, 0)
		for _, parser := range parsers {
			newParsers = append(newParsers, parser, Discard(Optional(ParseWhite)))
		}
		newParsers = newParsers[0 : len(newParsers)-1]
		return And(newParsers...)(s)
	}
}

// 尝试匹配列表中任意一个token
func OneOfToken(parser ParserFunc, tokenTexts ...string) ParserFunc {
	return func(s *Scanner) []Token {
		tokenList := parser(s)
		if len(tokenList) != 1 {
			return nil
		}
		if !isInStringArray(tokenTexts, tokenList[0].TokenText) {
			return nil
		}
		return tokenList
	}
}

// 尝试执行任意一个解析器，如果成功则返回
func Try(parsers ...ParserFunc) ParserFunc {
	return func(s *Scanner) []Token {
		for _, parser := range parsers {
			s.StackPush()
			tokenList := parser(s)
			if tokenList != nil {
				s.StackPop(false)
				return tokenList
			}
			s.StackPop(true)
		}
		return nil
	}
}

// 尝试可选的解析器
func Optional(parser ParserFunc) ParserFunc {
	return func(s *Scanner) []Token {
		s.StackPush()
		tokenList := parser(s)
		if tokenList != nil {
			s.StackPop(false)
			return tokenList
		} else {
			s.StackPop(true)
			return []Token{}
		}
	}
}

// 丢弃解析到的结果
func Discard(parser ParserFunc) ParserFunc {
	return func(s *Scanner) []Token {
		tokenList := parser(s)
		if tokenList == nil {
			return nil
		}
		return []Token{}
	}
}
