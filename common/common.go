package common

import (
	"fmt"
	"strconv"

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

func EvalLiteralExpression(le *parser.LiteralExpressionContext) object.Object {
	lit := le.Literal()
	booleanLit := lit.BooleanLiteral()
	if booleanLit != nil {
		return NativeBoolToBooleanObject(booleanLit.GetText() == "true")
	}
	integerLit := lit.IntegerLiteral()
	if integerLit != nil {
		value, err := strconv.ParseInt(integerLit.GetText(), 0, 64)
		if err != nil {
			return NewError("could not parse %s as integer", integerLit.GetText())
		}
		return &object.Integer{Value: value}
	}
	floatLit := lit.FloatLiteral()
	if floatLit != nil {
		value, err := strconv.ParseFloat(floatLit.GetText(), 64)
		if err != nil {
			return NewError("could not parse %q as float", floatLit.GetText())
		}
		return &object.Float{Value: value}
	}
	stringLit := lit.StringLiteral()
	if stringLit != nil {
		rawText := stringLit.GetText()
		return &object.String{Value: rawText[1 : len(rawText)-1]}
	}
	return NewError("unexpected literal: %s", lit.GetText())
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

func NativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

type MonkeyErrListener struct {
	Errors []string
}

func NewErrListener() *MonkeyErrListener {
	return &MonkeyErrListener{Errors: make([]string, 0)}
}

func (mel *MonkeyErrListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
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
