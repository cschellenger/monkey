package common

import (
	"encoding/gob"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"
)

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

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
	fmt.Printf("Syntax error: %v, msg: %s\n", offendingSymbol, msg)
	mel.Errors = append(mel.Errors, msg)
	//panic(errors.New(msg))
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

func RegisterGobTypes() {
	gob.Register(&object.Integer{})
	gob.Register(&object.Float{})
	gob.Register(&object.Boolean{})
	gob.Register(&object.String{})
	gob.Register(&object.Null{})
	gob.Register(&object.ReturnValue{})
	gob.Register(&object.Error{})
	gob.Register(&object.Function{})
	gob.Register(&object.Builtin{})
	gob.Register(&object.Array{})
	gob.Register(&object.Hash{})
	gob.Register(&object.HashPair{})
	gob.Register(&object.CompiledFunction{})
	gob.Register(&object.Closure{})
}
