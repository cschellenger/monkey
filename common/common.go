package common

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

type InfixExpression interface {
	Expression(i int) parser.IExpressionContext
	GetOp() antlr.Token
}

func EvalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func NewError(format string, a ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
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
