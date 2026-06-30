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
	instructions code.Instructions
	constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
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

	case *parser.StarSlashExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.PlusMinusExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.RelationExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.LiteralExpressionContext:
		return c.compileLiteralExpression(node)
	}
	return nil
}

func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)
	return pos
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
			return common.NewError("could not parse %q as float", floatLit.GetSymbol().GetText())
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
