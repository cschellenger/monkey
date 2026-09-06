package vm

import (
	"fmt"

	"github.com/cschellenger/monkey/code"
	"github.com/cschellenger/monkey/compiler"
	"github.com/cschellenger/monkey/object"
)

const StackSize = 2048
const GlobalsSize = 65536
const MaxFrames = 1024

var True = &object.Boolean{Value: true}
var False = &object.Boolean{Value: false}
var Null = &object.Null{}

type VM struct {
	constants []object.Object

	stack   []object.Object
	sp      int // Always points to the next value. Top of stack is stack[sp-1]
	globals []object.Object

	frames      []*Frame
	framesIndex int
}

func New(bytecode *compiler.Bytecode) *VM {
	mainFn := object.CompiledFunction{Instructions: bytecode.Instructions}
	mainFrame := NewFrame(&mainFn, 0)

	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame

	return &VM{
		constants: bytecode.Constants,

		stack:   make([]object.Object, StackSize),
		sp:      0,
		globals: make([]object.Object, GlobalsSize),

		frames:      frames,
		framesIndex: 1,
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
	var ip int
	var ins code.Instructions
	var op code.Opcode

	for vm.currentFrame().ip < len(vm.currentFrame().Instructions())-1 {
		vm.currentFrame().ip++

		ip = vm.currentFrame().ip
		ins = vm.currentFrame().Instructions()
		op = code.Opcode(ins[ip])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

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
			pos := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip = pos - 1

		case code.OpJumpNotTruthy:
			pos := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				vm.currentFrame().ip = pos - 1
			}

		case code.OpSetGlobal:
			globalIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			vm.globals[globalIndex] = vm.pop()

		case code.OpGetGlobal:
			globalIndex := code.ReadUint16(ins[ip+1:])
			vm.currentFrame().ip += 2

			err := vm.push(vm.globals[globalIndex])
			if err != nil {
				return err
			}

		case code.OpArray:
			numElements := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			array := vm.buildArray(vm.sp-numElements, vm.sp)
			vm.sp = vm.sp - numElements

			err := vm.push(array)
			if err != nil {
				return err
			}

		case code.OpHash:
			numElements := int(code.ReadUint16(ins[ip+1:]))
			vm.currentFrame().ip += 2

			hash, err := vm.buildHash(vm.sp-numElements, vm.sp)
			if err != nil {
				return err
			}
			vm.sp = vm.sp - numElements

			err = vm.push(hash)
			if err != nil {
				return err
			}

		case code.OpIndex:
			index := vm.pop()
			left := vm.pop()

			err := vm.executeIndexExpression(left, index)
			if err != nil {
				return err
			}

		case code.OpCall:
			numArgs := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			err := vm.executeCall(int(numArgs))
			if err != nil {
				return err
			}

		case code.OpReturnValue:
			returnValue := vm.pop()

			frame := vm.popFrame()
			vm.sp = frame.basePointer - 1

			err := vm.push(returnValue)
			if err != nil {
				return err
			}

		case code.OpReturn:
			frame := vm.popFrame()
			vm.sp = frame.basePointer - 1

			err := vm.push(Null)
			if err != nil {
				return err
			}

		case code.OpSetLocal:
			localIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			frame := vm.currentFrame()

			vm.stack[frame.basePointer+int(localIndex)] = vm.pop()

		case code.OpGetLocal:
			localIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			frame := vm.currentFrame()

			err := vm.push(vm.stack[frame.basePointer+int(localIndex)])
			if err != nil {
				return err
			}
		case code.OpGetBuiltin:
			builtinIndex := code.ReadUint8(ins[ip+1:])
			vm.currentFrame().ip += 1

			definition := object.Builtins[builtinIndex]

			err := vm.push(definition.Builtin)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (vm *VM) executeIndexExpression(left, index object.Object) error {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return vm.executeArrayIndex(left, index)
	case left.Type() == object.HASH_OBJ:
		return vm.executeHashIndex(left, index)
	default:
		return fmt.Errorf("index operator not supported: %s", left.Type())
	}
}

func (vm *VM) executeArrayIndex(array, index object.Object) error {
	arrayObject := array.(*object.Array)
	i := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if i < 0 || i > max {
		return vm.push(Null)
	}

	return vm.push(arrayObject.Elements[i])
}

func (vm *VM) executeHashIndex(hash, index object.Object) error {
	hashObject := hash.(*object.Hash)

	key, ok := index.(object.Hashable)
	if !ok {
		return fmt.Errorf("unusable as hash key: %s", index.Type())
	}

	pair, ok := hashObject.Pairs[key.HashKey()]
	if !ok {
		return vm.push(Null)
	}

	return vm.push(pair.Value)
}

func (vm *VM) buildArray(startIndex, endIndex int) object.Object {
	elements := make([]object.Object, endIndex-startIndex)

	for i := startIndex; i < endIndex; i++ {
		elements[i-startIndex] = vm.stack[i]
	}

	return &object.Array{Elements: elements}
}

func (vm *VM) buildHash(startIndex, endIndex int) (object.Object, error) {
	hashedPairs := make(map[object.HashKey]object.HashPair)

	for i := startIndex; i < endIndex; i += 2 {
		key := vm.stack[i]
		value := vm.stack[i+1]

		pair := object.HashPair{Key: key, Value: value}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return nil, fmt.Errorf("unusable as hash key: %s", key.Type())
		}

		hashedPairs[hashKey.HashKey()] = pair
	}

	return &object.Hash{Pairs: hashedPairs}, nil
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
			result = NativeBoolToBooleanObject(leftOpEqual.Equals(right))
		} else {
			return object.UnsupportedOperation(left.Type(), "==", right.Type())
		}
	case code.OpNotEqual:
		if leftNotEqual, ok := left.(object.OperatorEquals); ok {
			result = NativeBoolToBooleanObject(!leftNotEqual.Equals(right))
		} else {
			return object.UnsupportedOperation(left.Type(), "!=", right.Type())
		}
	case code.OpGreaterThan:
		if leftGreaterThan, ok := left.(object.Comparable); ok {
			if cmp, err := leftGreaterThan.CompareTo(right); err == nil {
				result = NativeBoolToBooleanObject(cmp == object.GREATER_THAN)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), ">", right.Type())
		}
	case code.OpGreaterEqual:
		if leftGreaterEqual, ok := left.(object.Comparable); ok {
			if cmp, err := leftGreaterEqual.CompareTo(right); err == nil {
				result = NativeBoolToBooleanObject(cmp == object.GREATER_THAN || cmp == object.EQUAL)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), ">=", right.Type())
		}
	case code.OpLessThan:
		if leftLessThan, ok := left.(object.Comparable); ok {
			if cmp, err := leftLessThan.CompareTo(right); err == nil {
				result = NativeBoolToBooleanObject(cmp == object.LESS_THAN)
			} else {
				return err
			}
		} else {
			return object.UnsupportedOperation(left.Type(), "<", right.Type())
		}
	case code.OpLessEqual:
		if leftLessEqual, ok := left.(object.Comparable); ok {
			if cmp, err := leftLessEqual.CompareTo(right); err == nil {
				result = NativeBoolToBooleanObject(cmp == object.LESS_THAN || cmp == object.EQUAL)
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

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.framesIndex-1]
}

func (vm *VM) pushFrame(f *Frame) {
	vm.frames[vm.framesIndex] = f
	vm.framesIndex++
}

func (vm *VM) popFrame() *Frame {
	vm.framesIndex--
	return vm.frames[vm.framesIndex]
}

func NativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return True
	}
	return False
}

func (vm *VM) executeCall(numArgs int) error {
	callee := vm.stack[vm.sp-1-numArgs]
	switch callee := callee.(type) {
	case *object.CompiledFunction:
		return vm.callFunction(callee, numArgs)
	case *object.Builtin:
		return vm.callBuiltin(callee, numArgs)
	default:
		return fmt.Errorf("calling non-function and non-built-in")
	}
}

func (vm *VM) callFunction(fn *object.CompiledFunction, numArgs int) error {
	if numArgs != fn.NumParams {
		return fmt.Errorf("wrong number of arguments: want=%d, got=%d",
			fn.NumParams, numArgs)
	}

	frame := NewFrame(fn, vm.sp-numArgs)
	vm.pushFrame(frame)

	vm.sp = frame.basePointer + fn.NumLocals

	return nil
}

func (vm *VM) callBuiltin(builtin *object.Builtin, numArgs int) error {
	args := vm.stack[vm.sp-numArgs : vm.sp]

	result := builtin.Fn(args...)
	vm.sp = vm.sp - numArgs - 1

	if result != nil {
		vm.push(result)
	} else {
		vm.push(Null)
	}

	return nil
}
