package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"strings"

	"github.com/cschellenger/monkey/ast"
)

type ObjectType string

const (
	INTEGER_OBJ      ObjectType = "INTEGER"
	FLOAT_OBJ        ObjectType = "FLOAT"
	BOOLEAN_OBJ      ObjectType = "BOOLEAN"
	STRING_OBJ       ObjectType = "STRING"
	NULL_OBJ         ObjectType = "NULL"
	RETURN_VALUE_OBJ ObjectType = "RETURN_VALUE"
	ERROR_OBJ        ObjectType = "ERROR"
	FUNCTION_OBJ     ObjectType = "FUNCTION"
	BUILTIN_OBJ      ObjectType = "BUILTIN"
	ARRAY_OBJ        ObjectType = "ARRAY"
	HASH_OBJ         ObjectType = "HASH"
)

type CompareResult int

const (
	LESS_THAN    CompareResult = -1
	EQUAL        CompareResult = 0
	GREATER_THAN CompareResult = 1
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type OperatorEquals interface {
	// Equality
	Equals(Object) bool
}

type OperatorPlus interface {
	// Usually "add"
	Plus(Object) (Object, *Error)
}

type OperatorMinus interface {
	// Usually "subtract"
	Minus(Object) (Object, *Error)
}

type OperatorStar interface {
	// Usually "multiply"
	Star(Object) (Object, *Error)
}

type OperatorSlash interface {
	// Usually "divide"
	Slash(Object) (Object, *Error)
}

type Comparable interface {
	CompareTo(Object) (CompareResult, *Error)
}

type Number interface {
	Object
	OperatorEquals
	OperatorPlus
	OperatorMinus
	OperatorStar
	OperatorSlash
	Comparable

	IntValue() int64
	FloatValue() float64
	Add(Number) Number
	Subtract(Number) Number
	Multiply(Number) Number
	Divide(Number) (Number, *Error)
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i *Integer) IntValue() int64 {
	return i.Value
}

func (i *Integer) FloatValue() float64 {
	return float64(i.Value)
}

func (i *Integer) Plus(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return i.Add(otherNum), nil
	}
	return nil, UnsupportedOperation(i.Type(), "+", other.Type())
}

func (i *Integer) Minus(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return i.Subtract(otherNum), nil
	}
	return nil, UnsupportedOperation(i.Type(), "-", other.Type())
}

func (i *Integer) Star(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return i.Multiply(otherNum), nil
	}
	return nil, UnsupportedOperation(i.Type(), "*", other.Type())
}

func (i *Integer) Slash(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return i.Divide(otherNum)
	}
	return nil, UnsupportedOperation(i.Type(), "/", other.Type())
}

func (i *Integer) Add(other Number) Number {
	if other.Type() == FLOAT_OBJ {
		return &Float{Value: i.FloatValue() + other.FloatValue()}
	}
	return &Integer{Value: i.Value + other.IntValue()}
}

func (i *Integer) Subtract(other Number) Number {
	if other.Type() == FLOAT_OBJ {
		return &Float{Value: i.FloatValue() - other.FloatValue()}
	}
	return &Integer{Value: i.Value - other.IntValue()}
}

func (i *Integer) Multiply(other Number) Number {
	if other.Type() == FLOAT_OBJ {
		return &Float{Value: i.FloatValue() * other.FloatValue()}
	}
	return &Integer{Value: i.Value * other.IntValue()}
}

func (i *Integer) Divide(other Number) (Number, *Error) {
	if other.Type() == FLOAT_OBJ {
		otherFloat := other.(*Float)
		if otherFloat.Value == 0.0 {
			return nil, &Error{Message: "division by zero"}
		}
		return &Float{Value: i.FloatValue() / otherFloat.Value}, nil
	}
	otherInt := other.(*Integer)
	if otherInt.Value == 0 {
		return nil, &Error{Message: "division by zero"}
	}
	return &Integer{Value: i.Value / otherInt.Value}, nil
}

func (i *Integer) CompareTo(other Object) (CompareResult, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return compare(i, otherNum), nil
	}
	return EQUAL, UnsupportedOperation(i.Type(), "CompareTo", other.Type())
}

func (i *Integer) Equals(other Object) bool {
	cmp, err := i.CompareTo(other)
	if err != nil {
		return false
	}
	return cmp == EQUAL
}

type Float struct {
	Value float64
}

func (f *Float) Inspect() string {
	return fmt.Sprintf("%.2f", f.Value)
}

func (f *Float) Type() ObjectType {
	return FLOAT_OBJ
}

func (f *Float) IntValue() int64 {
	return int64(f.Value)
}

func (f *Float) FloatValue() float64 {
	return f.Value
}

func (f *Float) Plus(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return f.Add(otherNum), nil
	}
	return nil, UnsupportedOperation(f.Type(), "+", other.Type())
}

func (f *Float) Minus(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return f.Subtract(otherNum), nil
	}
	return nil, UnsupportedOperation(f.Type(), "-", other.Type())
}

func (f *Float) Star(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return f.Multiply(otherNum), nil
	}
	return nil, UnsupportedOperation(f.Type(), "*", other.Type())
}

func (f *Float) Slash(other Object) (Object, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return f.Divide(otherNum)
	}
	return nil, UnsupportedOperation(f.Type(), "/", other.Type())
}

func (f *Float) Add(other Number) Number {
	return &Float{Value: f.Value + other.FloatValue()}
}

func (f *Float) Subtract(other Number) Number {
	return &Float{Value: f.Value - other.FloatValue()}
}

func (f *Float) Multiply(other Number) Number {
	return &Float{Value: f.Value * other.FloatValue()}
}

func (f *Float) Divide(other Number) (Number, *Error) {
	if other.Type() == FLOAT_OBJ {
		otherFloat := other.(*Float)
		if otherFloat.Value == 0.0 {
			return nil, &Error{Message: "division by zero"}
		}
		return &Float{Value: f.Value / otherFloat.Value}, nil
	}
	// Convert integer divisor to float and perform division
	return &Float{Value: f.Value / float64(other.(*Integer).Value)}, nil
}

func (f *Float) CompareTo(other Object) (CompareResult, *Error) {
	otherNum, ok := other.(Number)
	if ok {
		return compare(f, otherNum), nil
	}
	return EQUAL, UnsupportedOperation(f.Type(), "CompareTo", other.Type())
}

func (f *Float) Equals(other Object) bool {
	cmp, err := f.CompareTo(other)
	if err != nil {
		return false
	}
	return cmp == EQUAL
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b *Boolean) Equals(other Object) bool {
	if other.Type() == BOOLEAN_OBJ {
		return b.Value == other.(*Boolean).Value
	}
	return false
}

type String struct {
	Value string
}

func (s *String) Inspect() string {
	return s.Value
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}

func (s *String) Equals(other Object) bool {
	if other.Type() == STRING_OBJ {
		return s.Value == other.(*String).Value
	}
	return false
}

func (s *String) Plus(other Object) (Object, *Error) {
	if other.Type() == STRING_OBJ {
		otherString := other.(*String)
		return &String{s.Value + otherString.Value}, nil
	}
	return nil, UnsupportedOperation(s.Type(), "+", other.Type())
}

type Null struct{}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Equals(other Object) bool {
	return other.Type() == NULL_OBJ
}

func compare(left, right Number) CompareResult {
	result := left.FloatValue() - right.FloatValue()
	var retVal CompareResult
	switch {
	case result < 0:
		retVal = LESS_THAN
	case result == 0:
		retVal = EQUAL
	case result > 0:
		retVal = GREATER_THAN
	}
	return retVal
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (e *Error) Equals(other Object) bool {
	if other.Type() == ERROR_OBJ {
		return other.(*Error).Message == e.Message
	}
	return false
}

func UnsupportedOperation(left ObjectType, operator string, right ObjectType) *Error {
	return &Error{fmt.Sprintf("Unsupported operation: %s %s %s", left, operator, right)}
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
func (f *Function) Equals(other Object) bool {
	return f == other
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return HASH_OBJ }

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hashable interface {
	HashKey() HashKey
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}
	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (f *Float) HashKey() HashKey {
	h := fnv.New64a()
	strVal := fmt.Sprintf("%f", f.Value)
	h.Write([]byte(strVal))

	return HashKey{Type: f.Type(), Value: h.Sum64()}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
