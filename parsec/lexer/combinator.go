package lexer

import (
	"strings"
)

func getParserNames(parsers []Parser) string {
	names := make([]string, 0)
	for _, p := range parsers {
		names = append(names, p.Name)
	}
	return strings.Join(names, ",")
}

// 尝试顺序执行解析器
func Sequence(parsers ...Parser) Parser {
	return ToParser("Sequence("+getParserNames(parsers)+")", func(s *TextScanner) []Token {
		result := make([]Token, 0)
		for _, parser := range parsers {
			tokenList := parser.Call(s)
			if tokenList == nil {
				return nil
			}
			result = append(result, tokenList...)
		}
		return result
	})
}

// 尝试顺序执行解析器，每个解析器中间加入可选的空白字符解析
func SequenceWithOptionalWhitespace(parsers ...Parser) Parser {
	newParsers := make([]Parser, 0)
	for _, parser := range parsers {
		newParsers = append(newParsers, parser, DiscardResult(Optional(ParseWhite)))
	}
	newParsers = newParsers[0 : len(newParsers)-1]
	return Sequence(newParsers...)
}

// 尝试执行任意一个解析器，如果成功则返回
func OneOf(parsers ...Parser) Parser {
	return ToParser("OneOf("+getParserNames(parsers)+")", func(s *TextScanner) []Token {
		for _, parser := range parsers {
			s.StackPush()
			tokenList := parser.Call(s)
			if tokenList != nil {
				s.StackPop(false)
				return tokenList
			}
			s.StackPop(true)
		}
		return nil
	})
}

// 尝试可选的解析器
func Optional(parser Parser) Parser {
	return ToParser("Optional("+parser.Name+")", func(s *TextScanner) []Token {
		s.StackPush()
		tokenList := parser.Call(s)
		if tokenList != nil {
			s.StackPop(false)
			return tokenList
		} else {
			s.StackPop(true)
			return []Token{}
		}
	})
}

func Many(parser Parser) Parser {
	return ToParser("Many("+parser.Name+")", func(s *TextScanner) []Token {
		tokenList := parser.Call(s)
		if tokenList == nil {
			return nil
		}
		for {
			s.StackPush()
			ret := parser.Call(s)
			if ret == nil {
				s.StackPop(true)
				break
			} else {
				s.StackPop(false)
				tokenList = append(tokenList, ret...)
			}
		}
		return tokenList
	})
}

// 重复一个或多个的解析器
func OneOrMore(parser Parser) Parser {
	return Sequence(parser, Optional(Many(parser)))
}

// 重复零个或多个的解析器
func ZeroOrMore(parser Parser) Parser {
	return Optional(Many(parser))
}

// 丢弃解析到的结果
func DiscardResult(parser Parser) Parser {
	return ToParser("DiscardResult("+parser.Name+")", func(s *TextScanner) []Token {
		tokenList := parser.Call(s)
		if tokenList == nil {
			return nil
		}
		return []Token{}
	})
}
