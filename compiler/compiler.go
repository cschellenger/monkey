package compiler

import (
	"fmt"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/code"
	"github.com/cschellenger/monkey/common"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"
)

type Compiler struct {
	instructions        code.Instructions
	constants           []object.Object
	lastInstruction     EmittedInstruction
	previousInstruction EmittedInstruction
	symbolTable         *SymbolTable
}

type EmittedInstruction struct {
	OpCode   code.Opcode
	Position int
}

func New() *Compiler {
	return &Compiler{
		instructions:        code.Instructions{},
		constants:           []object.Object{},
		previousInstruction: EmittedInstruction{},
		lastInstruction:     EmittedInstruction{},
		symbolTable:         NewSymbolTable(),
	}
}

func NewWithState(s *SymbolTable, constants []object.Object) *Compiler {
	compiler := New()
	compiler.symbolTable = s
	compiler.constants = constants
	return compiler
}

func (c *Compiler) Compile(node antlr.ParserRuleContext) error {
	switch node := node.(type) {
	case parser.IProgContext:
		for _, s := range node.AllStatement() {
			if err := c.Compile(s); err != nil {
				return err
			}
		}
	case *parser.ExpressionStatementContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		c.emit(code.OpPop)

	case *parser.ParenExpressionContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
	case *parser.NegatedExpressionContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		c.emit(code.OpBang)

	case *parser.IfExpressionContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		// emit on `OpJumpNotTruthy` with a bogus value
		jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)

		// main conditional block
		if err := c.Compile(node.CompoundStatement(0)); err != nil {
			return err
		}
		if c.lastInstructionIsPop() {
			c.removeLastPop()
		}

		// Emit an `OpJump` with a bogus value
		jumpPos := c.emit(code.OpJump, 9999)
		afterConsequencePos := len(c.instructions)
		c.changeOperand(jumpNotTruthyPos, afterConsequencePos)

		if node.ELSE() == nil {
			c.emit(code.OpNull)
		} else {
			// else block
			if err := c.Compile(node.CompoundStatement(1)); err != nil {
				return err
			}
			if c.lastInstructionIsPop() {
				c.removeLastPop()
			}
		}
		afterElsePos := len(c.instructions)
		c.changeOperand(jumpPos, afterElsePos)

	case *parser.CompoundStatementContext:
		for _, s := range node.AllStatement() {
			if err := c.Compile(s); err != nil {
				return err
			}
		}

	case *parser.StarSlashExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.PlusMinusExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.RelationExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.LiteralExpressionContext:
		return c.compileLiteralExpression(node)

	case *parser.LetStatementContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		symbol := c.symbolTable.Define(node.Identifier().GetText())
		c.emit(code.OpSetGlobal, symbol.Index)

	case *parser.IdentifierExpressionContext:
		symbol, ok := c.symbolTable.Resolve(node.Identifier().GetText())
		if !ok {
			return fmt.Errorf("undefined variable %s", node.Identifier().GetText())
		}
		c.emit(code.OpGetGlobal, symbol.Index)
	}

	return nil
}

func (c *Compiler) replaceInstruction(pos int, newInstruction []byte) {
	for i := range newInstruction {
		c.instructions[pos+i] = newInstruction[i]
	}
}

func (c *Compiler) changeOperand(opPos int, operand int) {
	op := code.Opcode(c.instructions[opPos])
	newInstruction := code.Make(op, operand)

	c.replaceInstruction(opPos, newInstruction)
}

func (c *Compiler) setLastInstruction(op code.Opcode, pos int) {
	previous := c.lastInstruction
	last := EmittedInstruction{OpCode: op, Position: pos}

	c.previousInstruction = previous
	c.lastInstruction = last
}

func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)

	c.setLastInstruction(op, pos)
	return pos
}

func (c *Compiler) lastInstructionIsPop() bool {
	return c.lastInstruction.OpCode == code.OpPop
}

func (c *Compiler) removeLastPop() {
	c.instructions = c.instructions[:c.lastInstruction.Position]
	c.lastInstruction = c.previousInstruction
}

func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return posNewInstruction
}

func (c *Compiler) addConstant(object object.Object) int {
	c.constants = append(c.constants, object)
	return len(c.constants) - 1
}

func (c *Compiler) compileInfixExpression(node common.InfixExpression) error {
	if err := c.Compile(node.Expression(0)); err != nil {
		return err
	}
	if err := c.Compile(node.Expression(1)); err != nil {
		return err
	}
	op := node.GetOp().GetText()
	switch op {
	case "+":
		c.emit(code.OpAdd)
	case "-":
		c.emit(code.OpSub)
	case "*":
		c.emit(code.OpMul)
	case "/":
		c.emit(code.OpDiv)
	case ">":
		c.emit(code.OpGreaterThan)
	case "<":
		c.emit(code.OpLessThan)
	case ">=":
		c.emit(code.OpGreaterEqual)
	case "<=":
		c.emit(code.OpLessEqual)
	case "==":
		c.emit(code.OpEqual)
	case "!=":
		c.emit(code.OpNotEqual)
	default:
		return fmt.Errorf("unknown operator: %s", op)
	}
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

func (c *Compiler) compileLiteralExpression(le *parser.LiteralExpressionContext) error {
	lit := le.Literal()
	booleanLit := lit.BooleanLiteral()
	if booleanLit != nil {
		if booleanLit.GetSymbol().GetText() == "true" {
			c.emit(code.OpTrue)
		} else {
			c.emit(code.OpFalse)
		}
		return nil
	}
	integerLit := lit.IntegerLiteral()
	if integerLit != nil {
		value, err := strconv.ParseInt(integerLit.GetSymbol().GetText(), 0, 64)
		if err != nil {
			return common.NewError("could not parse %s as integer", integerLit.GetSymbol().GetText())
		}
		c.emit(code.OpConstant, c.addConstant(&object.Integer{Value: value}))
		return nil
	}
	floatLit := lit.FloatLiteral()
	if floatLit != nil {
		value, err := strconv.ParseFloat(floatLit.GetSymbol().GetText(), 64)
		if err != nil {
			return common.
				NewError("could not parse %q as float", floatLit.GetSymbol().GetText())
		}
		c.emit(code.OpConstant, c.addConstant(&object.Float{Value: value}))
		return nil
	}
	stringLit := lit.StringLiteral()
	if stringLit != nil {
		rawText := stringLit.GetText()
		c.emit(code.OpConstant, c.addConstant(&object.String{Value: rawText[1 : len(rawText)-1]}))
		return nil
	}
	return common.NewError("unexpected literal: %s", lit.GetText())
}
