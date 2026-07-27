package vm

import (
	"fmt"

	"github.com/cschellenger/monkey/code"
	"github.com/cschellenger/monkey/common"
	"github.com/cschellenger/monkey/compiler"
	"github.com/cschellenger/monkey/object"
)

const StackSize = 2048
const GlobalsSize = 65536

var True = &object.Boolean{Value: true}
var False = &object.Boolean{Value: false}
var Null = &object.Null{}

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack   []object.Object
	sp      int // Always points to the next value. Top of stack is stack[sp-1]
	globals []object.Object
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack:   make([]object.Object, StackSize),
		sp:      0,
		globals: make([]object.Object, GlobalsSize),
	}
}

func NewWithGlobalsStore(bytecode *compiler.Bytecode, s []object.Object) *VM {
	vm := New(bytecode)
	vm.globals = s
	return vm
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			if err := vm.push(vm.constants[constIndex]); err != nil {
				return err
			}

		case code.OpBang:
			if err := vm.executeBangOperator(); err != nil {
				return err
			}

		case code.OpNull:
			if err := vm.push(Null); err != nil {
				return err
			}

		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv, code.OpEqual, code.OpNotEqual, code.OpGreaterThan, code.OpGreaterEqual, code.OpLessThan, code.OpLessEqual:
			if err := vm.executeBinaryOperation(op); err != nil {
				return err
			}
		case code.OpTrue:
			if err := vm.push(True); err != nil {
				return err
			}

		case code.OpFalse:
			if err := vm.push(False); err != nil {
				return err
			}

		case code.OpPop:
			vm.pop()

		case code.OpJump:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip = pos - 1

		case code.OpJumpNotTruthy:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				ip = pos - 1
			}

		case code.OpSetGlobal:
			globalIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			vm.globals[globalIndex] = vm.pop()

		case code.OpGetGlobal:
			globalIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.globals[globalIndex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {

	case *object.Boolean:
		return obj.Value

	case *object.Null:
		return false

	default:
		return true
	}
}

func (vm *VM) executeBangOperator() error {
	operand := vm.pop()

	switch operand {
	case True:
		return vm.push(False)
	case False:
		return vm.push(True)
	case Null:
		return vm.push(True)
	default:
		return vm.push(False)
	}
}

func (vm *VM) executeBinaryOperation(op code.Opcode) error {
	right := vm.pop()
	left := vm.pop()
	var result object.Object
	var err *object.Error
	switch op {
	case code.OpAdd:
		if leftOpAdd, ok := left.(object.OperatorPlus); ok {
			if result, err = leftOpAdd.Plus(right); err != nil {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "+", right.Type())
		}

	case code.OpSub:
		if leftOpSub, ok := left.(object.OperatorMinus); ok {
			if result, err = leftOpSub.Minus(right); err != nil {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "-", right.Type())
		}

	case code.OpMul:
		if leftOpMul, ok := left.(object.OperatorStar); ok {
			if result, err = leftOpMul.Star(right); err != nil {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "*", right.Type())
		}

	case code.OpDiv:
		if leftOpDiv, ok := left.(object.OperatorSlash); ok {
			if result, err = leftOpDiv.Slash(right); err != nil {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "/", right.Type())
		}
	case code.OpEqual:
		if leftOpEqual, ok := left.(object.OperatorEquals); ok {
			result = common.NativeBoolToBooleanObject(leftOpEqual.Equals(right))
		} else {
			return object.UnsupportedOperation(left.Type(), "==", right.Type())
		}
	case code.OpNotEqual:
		if leftNotEqual, ok := left.(object.OperatorEquals); ok {
			result = common.NativeBoolToBooleanObject(!leftNotEqual.Equals(right))
		} else {
			return object.UnsupportedOperation(left.Type(), "!=", right.Type())
		}
	case code.OpGreaterThan:
		if leftGreaterThan, ok := left.(object.Comparable); ok {
			if cmp, err := leftGreaterThan.CompareTo(right); err == nil {
				result = common.NativeBoolToBooleanObject(cmp == object.GREATER_THAN)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), ">", right.Type())
		}
	case code.OpGreaterEqual:
		if leftGreaterEqual, ok := left.(object.Comparable); ok {
			if cmp, err := leftGreaterEqual.CompareTo(right); err == nil {
				result = common.NativeBoolToBooleanObject(cmp == object.GREATER_THAN || cmp == object.EQUAL)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), ">=", right.Type())
		}
	case code.OpLessThan:
		if leftLessThan, ok := left.(object.Comparable); ok {
			if cmp, err := leftLessThan.CompareTo(right); err == nil {
				result = common.NativeBoolToBooleanObject(cmp == object.LESS_THAN)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "<", right.Type())
		}
	case code.OpLessEqual:
		if leftLessEqual, ok := left.(object.Comparable); ok {
			if cmp, err := leftLessEqual.CompareTo(right); err == nil {
				result = common.NativeBoolToBooleanObject(cmp == object.LESS_THAN || cmp == object.EQUAL)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "<=", right.Type())
		}
	}

	vm.push(result)
	return nil
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}
	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}
