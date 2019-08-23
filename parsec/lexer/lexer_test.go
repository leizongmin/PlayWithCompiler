package lexer

import (
	"fmt"
	"testing"
)

func TestRunParser(t *testing.T) {
	element := OneOf(
		ParseID,
		ParseNumberLiteral,
	)
	whitespace := DiscardResult(ParseWhite)
	expression := Sequence(
		GenParseToken("OpenParent", "("),
		ZeroOrMore(OneOf(element, whitespace)),
		GenParseToken("CloseParent", ")"),
	)
	test := func(source string) {
		tokenList, err := RunParser(expression, source)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(source)
		for _, t := range tokenList {
			fmt.Printf("\t%+v\n", t)
		}
		fmt.Println("")
	}

	test("()")
	test("(hello)")
	test("(hello world)")
	test("(hello 123 456_789e100 123.5)")
}
