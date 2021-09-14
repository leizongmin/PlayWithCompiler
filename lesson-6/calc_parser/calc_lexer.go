// Code generated from Calc.g4 by ANTLR 4.9.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 70, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 5, 8, 41, 10, 8, 3, 9, 3, 9, 7, 9, 45, 10, 9, 12, 9, 14,
	9, 48, 11, 9, 3, 10, 3, 10, 3, 10, 6, 10, 53, 10, 10, 13, 10, 14, 10, 54,
	3, 11, 3, 11, 3, 11, 6, 11, 60, 10, 11, 13, 11, 14, 11, 61, 3, 12, 6, 12,
	65, 10, 12, 13, 12, 14, 12, 66, 3, 12, 3, 12, 2, 2, 13, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 3, 2, 8,
	3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 4, 2, 81, 81, 113, 113, 4, 2, 90, 90,
	122, 122, 5, 2, 50, 59, 67, 72, 99, 104, 5, 2, 11, 12, 14, 15, 34, 34,
	2, 75, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25,
	3, 2, 2, 2, 5, 27, 3, 2, 2, 2, 7, 29, 3, 2, 2, 2, 9, 31, 3, 2, 2, 2, 11,
	33, 3, 2, 2, 2, 13, 35, 3, 2, 2, 2, 15, 40, 3, 2, 2, 2, 17, 42, 3, 2, 2,
	2, 19, 49, 3, 2, 2, 2, 21, 56, 3, 2, 2, 2, 23, 64, 3, 2, 2, 2, 25, 26,
	7, 44, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 49, 2, 2, 28, 6, 3, 2, 2, 2,
	29, 30, 7, 45, 2, 2, 30, 8, 3, 2, 2, 2, 31, 32, 7, 47, 2, 2, 32, 10, 3,
	2, 2, 2, 33, 34, 7, 42, 2, 2, 34, 12, 3, 2, 2, 2, 35, 36, 7, 43, 2, 2,
	36, 14, 3, 2, 2, 2, 37, 41, 5, 17, 9, 2, 38, 41, 5, 19, 10, 2, 39, 41,
	5, 21, 11, 2, 40, 37, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 39, 3, 2, 2,
	2, 41, 16, 3, 2, 2, 2, 42, 46, 9, 2, 2, 2, 43, 45, 9, 3, 2, 2, 44, 43,
	3, 2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2,
	47, 18, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 50, 2, 2, 50, 52, 9,
	4, 2, 2, 51, 53, 9, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54,
	52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 20, 3, 2, 2, 2, 56, 57, 7, 50,
	2, 2, 57, 59, 9, 5, 2, 2, 58, 60, 9, 6, 2, 2, 59, 58, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 22, 3, 2, 2, 2,
	63, 65, 9, 7, 2, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64, 3,
	2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 69, 8, 12, 2, 2, 69,
	24, 3, 2, 2, 2, 8, 2, 40, 46, 54, 61, 66, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "NUMBER", "DECIMALNUMBER", "OCTONARYNUMBER",
	"HEXADECIMALNUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "NUMBER", "DECIMALNUMBER",
	"OCTONARYNUMBER", "HEXADECIMALNUMBER", "WHITESPACE",
}

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	l := new(CalcLexer)
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
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0              = 1
	CalcLexerT__1              = 2
	CalcLexerT__2              = 3
	CalcLexerT__3              = 4
	CalcLexerT__4              = 5
	CalcLexerT__5              = 6
	CalcLexerNUMBER            = 7
	CalcLexerDECIMALNUMBER     = 8
	CalcLexerOCTONARYNUMBER    = 9
	CalcLexerHEXADECIMALNUMBER = 10
	CalcLexerWHITESPACE        = 11
)
