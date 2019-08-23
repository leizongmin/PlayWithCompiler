package lexer

type ParserFunc = func(s *TextScanner) []Token

type Parser struct {
	Name string
	Func ParserFunc
}

func ToParser(name string, parser ParserFunc) Parser {
	return Parser{Name: name, Func: parser}
}

func (p Parser) Call(s *TextScanner) []Token {
	debugf("Call %s", p.Name)
	ret := p.Func(s)
	if ret == nil {
		debugf("\t Fail")
	} else {
		debugf("\t OK: %+v", ret)
	}
	return ret
}

type Token struct {
	TokenType string // 符号类型
	TokenText string // 符号文本
	Offset    int    // 位置
}
