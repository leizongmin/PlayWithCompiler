package main

import (
	"fmt"

	"github.com/leizongmin/PlayWithCompiler/parsec/lexer"
)

func main() {
	line := "int x = 999;\nresult = x + 123 + 456 * 789;"
	tokenList, err := Parse(line)
	if err != nil {
		panic(err)
	}
	fmt.Println(line)
	for _, t := range tokenList {
		fmt.Printf("\t%+v\n", t)
	}
	fmt.Println("")
}

func Parse(source string) ([]lexer.Token, error) {
	expression := lexer.OneOrMore(lexer.OneOf(
		lexer.ParseNumberLiteral,
		lexer.ParseID,
		lexer.GenParseToken("EQ", "="),
		lexer.GenParseToken("PLUS", "+"),
		lexer.GenParseToken("MINUS", "-"),
		lexer.GenParseToken("MULTIPLE", "*"),
		lexer.GenParseToken("DIVIDE", "/"),
		lexer.GenParseToken("SEMICOLON", ";"),
		lexer.DiscardResult(lexer.ParseWhite),
	))
	// lexer.Debug(true, 0)
	return lexer.RunParser(expression, lexer.NewTextScanner(source))
}
