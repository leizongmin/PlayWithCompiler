package lexer

// 解析空白字符
var ParseWhite = ToParser("ParseWhite", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	position := s.Position()
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
	return []Token{{TokenType: "White", TokenText: string(chars), Position: position}}
})

// 解析标志符
var ParseID = ToParser("ParseID", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	position := s.Position()
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
	return []Token{{TokenType: "ID", TokenText: string(chars), Position: position}}
})

// 解析关键词
func GenParseKeyword(keywords []string) *Parser {
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
	position := s.Position()
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
	return []Token{{TokenType: "IntLiteral", TokenText: string(chars), Position: position}}
})

// 解析浮点数字面量
var ParseFloatLiteral = ToParser("ParseFloatLiteral", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	position := s.Position()
	c := s.Current()
	if !IsDigit(c) {
		return nil
	}
	chars = append(chars, c)
	hasPoint := false
	for {
		c := s.Next()
		if IsDigit(c) || c == '.' || c == '_' || c == 'e' || c == 'E' {
			chars = append(chars, c)
			if c == '.' {
				hasPoint = true
			}
		} else {
			break
		}
	}
	if !hasPoint {
		return nil
	}
	return []Token{{TokenType: "FloatLiteral", TokenText: string(chars), Position: position}}
})

// 解析数值字面量
var ParseNumberLiteral = OneOf(ParseFloatLiteral, ParseIntLiteral)

// 解析字符串字面量
var ParseStringLiteral = ToParser("ParseStringLiteral", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	position := s.Position()
	c := s.Current()
	if c != '"' {
		return nil
	}
	chars = append(chars, c)
	lc := c
	for {
		c := s.Next()
		if c == '"' && lc != '\\' {
			chars = append(chars, c)
			s.Next()
			break
		} else {
			chars = append(chars, c)
			lc = c
		}
	}
	return []Token{{TokenType: "StringLiteral", TokenText: string(chars), Position: position}}
})

// 解析字符字面量
var ParseCharLiteral = ToParser("ParseCharLiteral", func(s *TextScanner) []Token {
	chars := make([]rune, 0)
	position := s.Position()
	c := s.Current()
	if c != '\'' {
		return nil
	}
	chars = append(chars, c)
	c2 := s.Next()
	if c2 == '\\' {
		c3 := s.Next()
		c4 := s.Next()
		if c4 != '\'' {
			return nil
		}
		chars = append(chars, c2, c3, c4)
	} else {
		c3 := s.Next()
		if c3 != '\'' {
			return nil
		}
		chars = append(chars, c2, c3)
	}
	s.Next()
	return []Token{{TokenType: "CharLiteral", TokenText: string(chars), Position: position}}
})

// 生成解析固定连续字符
func GenParseToken(tokenType string, text string) *Parser {
	chars := []rune(text)
	size := len(chars)
	if size < 1 {
		panic("GenParseToken failed: text is empty")
	}
	return ToParser("ParseToken'"+text+"'", func(s *TextScanner) []Token {
		position := s.Position()
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
		return []Token{{TokenType: tokenType, TokenText: text, Position: position}}
	})
}
