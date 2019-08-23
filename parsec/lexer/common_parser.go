package lexer

// 解析空白字符
var ParseWhite = ToParser("ParseWhite", func(s *TextScanner) []Token {
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
})

// 解析标志符
var ParseID = ToParser("ParseID", func(s *TextScanner) []Token {
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
})

// 解析关键词
func GenParseKeyword(keywords []string) Parser {
	return ToParser("ParseKeyword", func(s *TextScanner) []Token {
		tokenList := ParseID.Call(s)
		if len(tokenList) != 1 {
			return nil
		}
		if !IsInStringArray(keywords, tokenList[0].TokenText) {
			return nil
		}
		tokenList[0].TokenType = "Keyword"
		return tokenList
	})
}

// 解析整数字面量
var ParseIntLiteral = ToParser("ParseIntLiteral", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	offset := s.Offset()
	c := s.Current()
	if !IsDigit(c) {
		return nil
	}
	chars = append(chars, c)
	for {
		c := s.Next()
		if IsDigit(c) || c == '_' || c == 'e' || c == 'E' {
			chars = append(chars, c)
		} else {
			break
		}
	}
	return []Token{{TokenType: "IntLiteral", TokenText: string(chars), Offset: offset}}
})

// 解析浮点数字面量
var ParseFloatLiteral = ToParser("ParseFloatLiteral", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	offset := s.Offset()
	c := s.Current()
	if !IsDigit(c) {
		return nil
	}
	chars = append(chars, c)
	for {
		c := s.Next()
		if IsDigit(c) || c == '.' || c == '_' || c == 'e' || c == 'E' {
			chars = append(chars, c)
		} else {
			break
		}
	}
	return []Token{{TokenType: "FloatLiteral", TokenText: string(chars), Offset: offset}}
})

// 解析数值字面量
var ParseNumberLiteral = OneOf(ParseFloatLiteral, ParseIntLiteral)

// 生成解析固定连续字符
func GenParseToken(tokenType string, text string) Parser {
	chars := []rune(text)
	size := len(chars)
	if size < 1 {
		panic("GenParseToken failed: text is empty")
	}
	return ToParser("ParseToken'"+text+"'", func(s *TextScanner) []Token {
		offset := s.Offset()
		c := s.Current()
		if c != chars[0] {
			return nil
		}
		i := 1
		for i < size {
			c := s.Next()
			if c != chars[i] {
				return nil
			}
			i++
		}
		s.Next()
		return []Token{{TokenType: tokenType, TokenText: text, Offset: offset}}
	})
}
