package lexer

import (
	"fmt"
	"strings"
	"testing"
)

func TestRunParser(t *testing.T) {
	element := OneOf(
		ParseID,
		ParseNumberLiteral,
		ParseStringLiteral,
		ParseCharLiteral,
	)
	whitespace := DiscardResult(ParseWhite)
	expression := Sequence(
		GenParseToken("OpenParent", "("),
		ZeroOrMore(OneOf(element, whitespace)),
		GenParseToken("CloseParent", ")"),
	)
	test := func(source string) {
		scanner := NewTextScanner(source)
		tokenList, err := RunParser(expression, scanner)
		if err != nil {
			pos := scanner.Position()
			line, _, _ := scanner.GetLine(pos.Offset)
			fmt.Printf("%s\n%s^\n", line, strings.Repeat(" ", pos.Column))
			t.Error(err)
		} else {
			fmt.Println(source)
			for _, t := range tokenList {
				fmt.Printf("\t%+v\n", t)
			}
		}
		fmt.Println("")
	}

	test(`()`)
	test(`(hello)`)
	test(`(hello world)`)
	test(`(hello "abc" "a\n\nb" 'a' '\\' 123 456_789e100 123.5)`)
	test("(define hello (lambda (a b) (println a b)))")
}
