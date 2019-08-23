package lexer

import (
	"fmt"
)

func RunParser(parser Parser, source string) ([]Token, error) {
	scanner := NewTextScanner(source)
	tokenList := parser.Call(scanner)
	if tokenList == nil {
		pos := scanner.Position()
		return nil, fmt.Errorf("parse failed: unexpected character '%s' (line:%d,column:%d)", string(scanner.Current()), pos.Line+1, pos.Column+1)
	}
	return tokenList, nil
}
