// Code generated from Hello.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 23, 123,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 6, 4, 54, 10, 4, 13, 4, 14, 4,
	55, 3, 5, 3, 5, 7, 5, 60, 10, 5, 12, 5, 14, 5, 63, 11, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 75, 10, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	19, 3, 19, 3, 20, 3, 20, 7, 20, 103, 10, 20, 12, 20, 14, 20, 106, 11, 20,
	3, 21, 6, 21, 109, 10, 21, 13, 21, 14, 21, 110, 3, 21, 3, 21, 3, 22, 3,
	22, 5, 22, 117, 10, 22, 3, 22, 5, 22, 120, 10, 22, 3, 22, 3, 22, 3, 61,
	2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39,
	21, 41, 22, 43, 23, 3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124,
	6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 4, 2, 11, 11, 34, 34, 2, 131, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 3, 45, 3, 2, 2, 2, 5, 48, 3, 2, 2, 2, 7, 53,
	3, 2, 2, 2, 9, 57, 3, 2, 2, 2, 11, 66, 3, 2, 2, 2, 13, 74, 3, 2, 2, 2,
	15, 76, 3, 2, 2, 2, 17, 78, 3, 2, 2, 2, 19, 80, 3, 2, 2, 2, 21, 82, 3,
	2, 2, 2, 23, 84, 3, 2, 2, 2, 25, 86, 3, 2, 2, 2, 27, 88, 3, 2, 2, 2, 29,
	90, 3, 2, 2, 2, 31, 92, 3, 2, 2, 2, 33, 94, 3, 2, 2, 2, 35, 96, 3, 2, 2,
	2, 37, 98, 3, 2, 2, 2, 39, 100, 3, 2, 2, 2, 41, 108, 3, 2, 2, 2, 43, 119,
	3, 2, 2, 2, 45, 46, 7, 107, 2, 2, 46, 47, 7, 104, 2, 2, 47, 4, 3, 2, 2,
	2, 48, 49, 7, 107, 2, 2, 49, 50, 7, 112, 2, 2, 50, 51, 7, 118, 2, 2, 51,
	6, 3, 2, 2, 2, 52, 54, 9, 2, 2, 2, 53, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2,
	2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 8, 3, 2, 2, 2, 57, 61, 7,
	36, 2, 2, 58, 60, 11, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 63, 3, 2, 2, 2,
	61, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 64, 3, 2, 2, 2, 63, 61, 3,
	2, 2, 2, 64, 65, 7, 36, 2, 2, 65, 10, 3, 2, 2, 2, 66, 67, 7, 63, 2, 2,
	67, 12, 3, 2, 2, 2, 68, 75, 7, 64, 2, 2, 69, 70, 7, 64, 2, 2, 70, 75, 7,
	63, 2, 2, 71, 75, 7, 62, 2, 2, 72, 73, 7, 62, 2, 2, 73, 75, 7, 63, 2, 2,
	74, 68, 3, 2, 2, 2, 74, 69, 3, 2, 2, 2, 74, 71, 3, 2, 2, 2, 74, 72, 3,
	2, 2, 2, 75, 14, 3, 2, 2, 2, 76, 77, 7, 44, 2, 2, 77, 16, 3, 2, 2, 2, 78,
	79, 7, 45, 2, 2, 79, 18, 3, 2, 2, 2, 80, 81, 7, 37, 2, 2, 81, 20, 3, 2,
	2, 2, 82, 83, 7, 61, 2, 2, 83, 22, 3, 2, 2, 2, 84, 85, 7, 48, 2, 2, 85,
	24, 3, 2, 2, 2, 86, 87, 7, 46, 2, 2, 87, 26, 3, 2, 2, 2, 88, 89, 7, 93,
	2, 2, 89, 28, 3, 2, 2, 2, 90, 91, 7, 95, 2, 2, 91, 30, 3, 2, 2, 2, 92,
	93, 7, 125, 2, 2, 93, 32, 3, 2, 2, 2, 94, 95, 7, 127, 2, 2, 95, 34, 3,
	2, 2, 2, 96, 97, 7, 42, 2, 2, 97, 36, 3, 2, 2, 2, 98, 99, 7, 43, 2, 2,
	99, 38, 3, 2, 2, 2, 100, 104, 9, 3, 2, 2, 101, 103, 9, 4, 2, 2, 102, 101,
	3, 2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2,
	2, 2, 105, 40, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 109, 9, 5, 2, 2,
	108, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 108, 3, 2, 2, 2, 110,
	111, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 113, 8, 21, 2, 2, 113, 42,
	3, 2, 2, 2, 114, 116, 7, 15, 2, 2, 115, 117, 7, 12, 2, 2, 116, 115, 3,
	2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 120, 7, 12, 2,
	2, 119, 114, 3, 2, 2, 2, 119, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	122, 8, 22, 2, 2, 122, 44, 3, 2, 2, 2, 11, 2, 55, 61, 74, 102, 104, 110,
	116, 119, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'if'", "'int'", "", "", "'='", "", "'*'", "'+'", "'#'", "';'", "'.'",
	"','", "'['", "']'", "'{'", "'}'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "If", "Int", "IntLiteral", "StringLiteral", "AssignmentOP", "RelationalOP",
	"Star", "Plus", "Sharp", "SemiColon", "Dot", "Comm", "LeftBracket", "RightBracket",
	"LeftBrace", "RightBrace", "LeftParen", "RightParen", "Id", "Whitespace",
	"Newline",
}

var lexerRuleNames = []string{
	"If", "Int", "IntLiteral", "StringLiteral", "AssignmentOP", "RelationalOP",
	"Star", "Plus", "Sharp", "SemiColon", "Dot", "Comm", "LeftBracket", "RightBracket",
	"LeftBrace", "RightBrace", "LeftParen", "RightParen", "Id", "Whitespace",
	"Newline",
}

type Hello struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewHello produces a new lexer instance for the optional input antlr.CharStream.
//
// The *Hello instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewHello(input antlr.CharStream) *Hello {
	l := new(Hello)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Hello.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Hello tokens.
const (
	HelloIf            = 1
	HelloInt           = 2
	HelloIntLiteral    = 3
	HelloStringLiteral = 4
	HelloAssignmentOP  = 5
	HelloRelationalOP  = 6
	HelloStar          = 7
	HelloPlus          = 8
	HelloSharp         = 9
	HelloSemiColon     = 10
	HelloDot           = 11
	HelloComm          = 12
	HelloLeftBracket   = 13
	HelloRightBracket  = 14
	HelloLeftBrace     = 15
	HelloRightBrace    = 16
	HelloLeftParen     = 17
	HelloRightParen    = 18
	HelloId            = 19
	HelloWhitespace    = 20
	HelloNewline       = 21
)
