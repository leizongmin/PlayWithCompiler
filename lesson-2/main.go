package main

import (
	"fmt"

	"github.com/leizongmin/PlayWithCompiler/parsec/lexer"
)

func main() {
	line := "int a = 123"
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
	assignStatement := lexer.SequenceWithOptionalWhitespace(
		lexer.GenParseKeyword([]string{"int", "float"}),
		lexer.ParseID,
		lexer.GenParseToken("EQ", "="),
		lexer.ParseNumberLiteral,
	)
	return lexer.RunParser(assignStatement, lexer.NewTextScanner(source))
}
