package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      ObjectType = "INTEGER"
	FLOAT_OBJ        ObjectType = "FLOAT"
	BOOLEAN_OBJ      ObjectType = "BOOLEAN"
	STRING_OBJ       ObjectType = "STRING"
	NULL_OBJ         ObjectType = "NULL"
	RETURN_VALUE_OBJ ObjectType = "RETURN_VALUE"
	ERROR_OBJ        ObjectType = "ERROR"
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
	Equals(Object) bool
}

type Number interface {
	Object
	IntValue() int64
	FloatValue() float64
	Add(Number) Number
	Subtract(Number) Number
	Multiply(Number) Number
	Divide(Number) Number
	CompareTo(Number) CompareResult
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

func (i *Integer) Divide(other Number) Number {
	if other.Type() == FLOAT_OBJ {
		otherFloat := other.(*Float)
		if otherFloat.Value == 0.0 {
			return &Error{Message: "division by zero"}
		}
		return &Float{Value: i.FloatValue() / otherFloat.Value}
	}
	otherInt := other.(*Integer)
	if otherInt.Value == 0 {
		return &Error{Message: "division by zero"}
	}
	return &Integer{Value: i.Value / otherInt.Value}
}

func (i *Integer) CompareTo(other Number) CompareResult {
	return compare(i, other)
}

func (i *Integer) Equals(other Object) bool {
	if other.Type() == INTEGER_OBJ {
		return i.Value == other.(*Integer).Value
	}
	return false
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

func (f *Float) Add(other Number) Number {
	return &Float{Value: f.Value + other.FloatValue()}
}

func (f *Float) Subtract(other Number) Number {
	return &Float{Value: f.Value - other.FloatValue()}
}

func (f *Float) Multiply(other Number) Number {
	return &Float{Value: f.Value * other.FloatValue()}
}

func (f *Float) Divide(other Number) Number {
	if other.Type() == FLOAT_OBJ {
		otherFloat := other.(*Float)
		if otherFloat.Value == 0.0 {
			return &Error{Message: "division by zero"}
		}
		return &Float{Value: f.Value / otherFloat.Value}
	}
	// Convert integer divisor to float and perform division
	return &Float{Value: f.Value / float64(other.(*Integer).Value)}
}

func (f *Float) CompareTo(other Number) CompareResult {
	return compare(f, other)
}

func (f *Float) Equals(other Object) bool {
	if other.Type() == FLOAT_OBJ {
		return f.Value == other.(*Float).Value
	}
	return false
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

func (rv *ReturnValue) Equals(other Object) bool {
	if other.Type() == RETURN_VALUE_OBJ {
		orv := other.(*ReturnValue)
		if rv.Value == nil {
			return orv.Value == nil
		}
		return rv.Value.Equals(orv.Value)
	}
	return false
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
