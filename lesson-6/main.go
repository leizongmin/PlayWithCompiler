package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/leizongmin/PlayWithCompiler/lesson-6/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream(`int age = 45;
if (age >= 17+8+20){
  printf("Hello old man!");
}`)

	// Create the Lexer
	lexer := parser.NewHello(is)

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n", lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}
