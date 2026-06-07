package evaluator

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

func Eval(node antlr.ParserRuleContext, env *object.Environment) object.Object {
	switch node := node.(type) {

	case parser.IProgContext:
		return evalProgram(node, env)

	case *parser.ReturnStatementContext:
		val := Eval(node.Expression(), env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *parser.ExpressionStatementContext:
		return Eval(node.Expression(), env)

	case *parser.NegatedExpressionContext:
		right := Eval(node.Expression(), env)
		if isError(right) {
			return right
		}
		return evalBangOperatorExpression(right)

	case *parser.StarSlashExpressionContext:
		return evalInfixExpression(node, env)
	case *parser.PlusMinusExpressionContext:
		return evalInfixExpression(node, env)
	case *parser.RelationExpressionContext:
		return evalInfixExpression(node, env)

	case *parser.LetStatementContext:
		val := Eval(node.Expression(), env)
		if isError(val) {
			return val
		}
		env.Set(node.Identifier().GetText(), val)

	case *parser.FunctionExpressionContext:
		lit := node.Function_literal()
		params := lit.Params()
		body := lit.Statement()
		return &object.Function{Parameters: params, Env: env, Body: body}

	case *parser.CallExpressionContext:
		call := node.Call_expression()
		ident := call.Identifier()
		function := evalIdentifier(ident, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(call.Expression_list(), env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		return applyFunction(function, args)

	case *parser.IdentifierExpressionContext:
		return evalIdentifier(node.Identifier(), env)

	case *parser.CompoundStatementContext:
		return evalBlockStatement(node, env)

	case *parser.IfExpressionContext:
		return evalIfExpression(node, env)

	case *parser.LiteralExpressionContext:
		return evalLiteralExpression(node)

	case *parser.WhileStatementContext:
		return evalWhileLoop(node, env)

	case *parser.ArrayLiteralExpressionContext:
		elements := evalExpressions(node.Array_literal().Expression_list(), env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}

	case *parser.HashLiteralExpressionContext:
		return evalHashLiteral(node, env)

	case *parser.IndexExpressionContext:
		left := Eval(node.Expression(0), env)
		if isError(left) {
			return left
		}
		index := Eval(node.Expression(1), env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index)

	case *parser.ParenExpressionContext:
		return Eval(node.Expression(), env)

	case nil:
		return NULL
	}

	return nil
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)

	case *object.Builtin:
		return fn.Fn(args...)

	default:
		return newError("not a function: %s", fn.Type())
	}
}

func evalLiteralExpression(le *parser.LiteralExpressionContext) object.Object {
	lit := le.Literal()
	booleanLit := lit.BooleanLiteral()
	if booleanLit != nil {
		return nativeBoolToBooleanObject(booleanLit.GetText() == "true")
	}
	integerLit := lit.IntegerLiteral()
	if integerLit != nil {
		value, err := strconv.ParseInt(integerLit.GetText(), 0, 64)
		if err != nil {
			return newError("could not parse %s as integer", integerLit.GetText())
		}
		return &object.Integer{Value: value}
	}
	floatLit := lit.FloatLiteral()
	if floatLit != nil {
		value, err := strconv.ParseFloat(floatLit.GetText(), 64)
		if err != nil {
			return newError("could not parse %q as float", floatLit.GetText())
		}
		return &object.Float{Value: value}
	}
	stringLit := lit.StringLiteral()
	if stringLit != nil {
		rawText := stringLit.GetText()
		return &object.String{Value: rawText[1 : len(rawText)-1]}
	}
	return newError("unexpected literal: %s", lit.GetText())
}

func extendFunctionEnv(
	fn *object.Function,
	args []object.Object,
) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters.AllIdentifier() {
		env.Set(param.GetText(), args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func evalExpressions(
	exps parser.IExpression_listContext,
	env *object.Environment,
) []object.Object {
	var result []object.Object
	for _, e := range exps.AllExpression() {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}

func evalIdentifier(
	node parser.IIdentifierContext,
	env *object.Environment,
) object.Object {
	identifier := node.GetText()
	if val, ok := env.Get(identifier); ok {
		return val
	}

	if builtin, ok := builtins[identifier]; ok {
		return builtin
	}

	return newError("identifier not found: %s", identifier)
}

func evalProgram(program parser.IProgContext, env *object.Environment) object.Object {
	var result object.Object
	for _, statement := range program.AllStatement() {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}
	return result
}

func evalBlockStatement(block *parser.CompoundStatementContext, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.AllStatement() {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalBangOperatorExpression(right object.Object) object.Object {
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

func evalInfixExpression(
	node InfixExpression,
	env *object.Environment,
) object.Object {
	left := Eval(node.Expression(0), env)
	if isError(left) {
		return left
	}
	right := Eval(node.Expression(1), env)
	if isError(right) {
		return right
	}
	operator := node.GetOp().GetText()
	switch {
	case operator == "+":
		leftPlus, ok := left.(object.OperatorPlus)
		if ok {
			return errOrResult(leftPlus.Plus(right))
		}
		return object.UnsupportedOperation(left.Type(), operator, right.Type())
	case operator == "-":
		leftMinus, ok := left.(object.OperatorMinus)
		if ok {
			return errOrResult(leftMinus.Minus(right))
		}
		return object.UnsupportedOperation(left.Type(), operator, right.Type())
	case operator == "*":
		leftStar, ok := left.(object.OperatorStar)
		if ok {
			return errOrResult(leftStar.Star(right))
		}
		return object.UnsupportedOperation(left.Type(), operator, right.Type())
	case operator == "/":
		leftSlash, ok := left.(object.OperatorSlash)
		if ok {
			return errOrResult(leftSlash.Slash(right))
		}
		return object.UnsupportedOperation(left.Type(), operator, right.Type())
	case operator == "==":
		leftEquals, ok := left.(object.OperatorEquals)
		if ok {
			return nativeBoolToBooleanObject(left == right || leftEquals.Equals(right))
		}
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		leftEquals, ok := left.(object.OperatorEquals)
		if ok {
			return nativeBoolToBooleanObject(left != right && !leftEquals.Equals(right))
		}
		return nativeBoolToBooleanObject(left != right)
	case operator == ">=":
		return evalComparable(left, right, operator, compareGreaterEquals)
	case operator == ">":
		return evalComparable(left, right, operator, compareGreater)
	case operator == "<=":
		return evalComparable(left, right, operator, compareLessEquals)
	case operator == "<":
		return evalComparable(left, right, operator, compareLess)
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s",
			left.Type(), operator, right.Type())
	default:
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
}

func compareGreater(cmp object.CompareResult) bool {
	return cmp == object.GREATER_THAN
}

func compareGreaterEquals(cmp object.CompareResult) bool {
	return cmp == object.EQUAL || cmp == object.GREATER_THAN
}

func compareLess(cmp object.CompareResult) bool {
	return cmp == object.LESS_THAN
}

func compareLessEquals(cmp object.CompareResult) bool {
	return cmp == object.EQUAL || cmp == object.LESS_THAN
}

func evalComparable(
	left, right object.Object,
	operator string,
	comparison func(cmp object.CompareResult) bool,
) object.Object {
	leftComparable, ok := left.(object.Comparable)
	if ok {
		result, err := leftComparable.CompareTo(right)
		if err != nil {
			return err
		}
		return nativeBoolToBooleanObject(comparison(result))
	}
	return object.UnsupportedOperation(left.Type(), operator, right.Type())
}

func evalStringInfixExpression(
	operator string,
	left, right object.Object,
) object.Object {
	if operator != "+" {
		return newError("unknown operator: %s %s %s",
			left.Type(), operator, right.Type())
	}
	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value
	return &object.String{Value: leftVal + rightVal}
}

func evalWhileLoop(wl *parser.WhileStatementContext, env *object.Environment) object.Object {
	keepGoing := true
	var retVal object.Object = NULL
	for keepGoing {
		condition := Eval(wl.Expression(), env)
		if isError(condition) {
			return condition
		}
		keepGoing = isTruthy(condition)
		if !keepGoing {
			break
		}
		retVal = Eval(wl.Statement(), env)
		if retVal.Type() == object.RETURN_VALUE_OBJ {
			break
		}
	}
	return retVal
}

func evalIfExpression(ie *parser.IfExpressionContext, env *object.Environment) object.Object {
	if_exp := ie.If_expression()
	condition := Eval(if_exp.Expression(), env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(if_exp.Statement(0), env)
	} else if if_exp.ELSE() != nil {
		return Eval(if_exp.Statement(1), env)
	} else {
		return NULL
	}
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.HASH_OBJ:
		return evalHashIndexExpression(left, index)
	default:
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalHashIndexExpression(hash, index object.Object) object.Object {
	hashObject := hash.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return newError("unusable as hash key: %s", index.Type())
	}
	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}
	return pair.Value
}

func evalHashLiteral(
	node *parser.HashLiteralExpressionContext,
	env *object.Environment,
) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for _, pairNode := range node.Hash_literal().AllExpression_pair() {
		key := Eval(pairNode.Expression(0), env)
		if isError(key) {
			return key
		}
		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}
		val := Eval(pairNode.Expression(1), env)
		if isError(val) {
			return val
		}
		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: val}
	}
	return &object.Hash{Pairs: pairs}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string, a ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func errOrResult(res object.Object, err *object.Error) object.Object {
	if err != nil {
		return err
	}
	if res != nil {
		return res
	}
	return NULL
}

type InfixExpression interface {
	Expression(i int) parser.IExpressionContext
	GetOp() antlr.Token
}
