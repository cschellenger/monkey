package compiler

import (
	"fmt"

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

	case *parser.StarSlashExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.PlusMinusExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.RelationExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.LiteralExpressionContext:
		lit := common.EvalLiteralExpression(node)
		c.emit(code.OpConstant, c.addConstant(lit))
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
