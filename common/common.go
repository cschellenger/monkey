package common

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/parser"
)

type InfixExpression interface {
	Expression(i int) parser.IExpressionContext
	GetOp() antlr.Token
}

type MonkeyErrListener struct {
	Errors []string
}

func NewErrListener() *MonkeyErrListener {
	return &MonkeyErrListener{Errors: make([]string, 0)}
}

func (mel *MonkeyErrListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	mel.Errors = append(mel.Errors, msg)
}

func (mel *MonkeyErrListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	//fmt.Printf("Got ambiguity\n")
}
func (mel *MonkeyErrListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	//fmt.Printf("Got attempting full context\n")
}
func (mel *MonkeyErrListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	//fmt.Printf("Got context sensitivity\n")
}
