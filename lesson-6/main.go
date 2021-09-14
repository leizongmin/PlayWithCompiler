package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/leizongmin/PlayWithCompiler/lesson-6/calc_parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream(`0x123 + 456_780 * 5 / 2`)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)

	// Read all tokens
	//for {
	//	t := lexer.NextToken()
	//	if t.GetTokenType() == antlr.TokenEOF {
	//		break
	//	}
	//	fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	//}

	s := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(s)
	prog := p.Prog()
	fmt.Printf("%+v\n", prog)
	fmt.Println(prog.GetText())
	fmt.Printf("%+v\n", prog.GetPayload())
}
