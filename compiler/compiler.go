package compiler

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/cschellenger/monkey/code"
	"github.com/cschellenger/monkey/common"
	"github.com/cschellenger/monkey/object"
	"github.com/cschellenger/monkey/parser"
)

type Compiler struct {
	constants   []object.Object
	symbolTable *SymbolTable
	scopes      []CompilationScope
	scopeIndex  int
}

type CompilationScope struct {
	instructions        code.Instructions
	lastInstruction     EmittedInstruction
	previousInstruction EmittedInstruction
	currentIdentifier   *Symbol
}

type EmittedInstruction struct {
	OpCode   code.Opcode
	Position int
}

func New() *Compiler {
	mainScope := CompilationScope{
		instructions:        code.Instructions{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}

	symbolTable := NewSymbolTable()

	for i, v := range object.Builtins {
		symbolTable.DefineBuiltin(i, v.Name)
	}

	return &Compiler{
		constants:   []object.Object{},
		symbolTable: symbolTable,
		scopes:      []CompilationScope{mainScope},
		scopeIndex:  0,
	}
}

func NewWithState(s *SymbolTable, constants []object.Object) *Compiler {
	compiler := New()
	compiler.symbolTable = s
	compiler.constants = constants
	return compiler
}

func (c *Compiler) currentInstructions() code.Instructions {
	return c.scopes[c.scopeIndex].instructions
}

func (c *Compiler) addInstruction(ins []byte) int {
	posNewInstruction := len(c.currentInstructions())
	updatedInstructions := append(c.currentInstructions(), ins...)

	c.scopes[c.scopeIndex].instructions = updatedInstructions

	return posNewInstruction
}

func (c *Compiler) setLastInstruction(op code.Opcode, pos int) {
	previous := c.scopes[c.scopeIndex].lastInstruction
	last := EmittedInstruction{OpCode: op, Position: pos}

	c.scopes[c.scopeIndex].previousInstruction = previous
	c.scopes[c.scopeIndex].lastInstruction = last
}

func (c *Compiler) lastInstructionIs(op code.Opcode) bool {
	if len(c.currentInstructions()) == 0 {
		return false
	}

	return c.scopes[c.scopeIndex].lastInstruction.OpCode == op
}

func (c *Compiler) removeLastPop() {
	last := c.scopes[c.scopeIndex].lastInstruction
	previous := c.scopes[c.scopeIndex].previousInstruction

	old := c.currentInstructions()
	new := old[:last.Position]

	c.scopes[c.scopeIndex].instructions = new
	c.scopes[c.scopeIndex].lastInstruction = previous
}

func (c *Compiler) replaceInstruction(pos int, newInstruction []byte) {
	ins := c.currentInstructions()

	for i := range newInstruction {
		ins[pos+i] = newInstruction[i]
	}
}

func (c *Compiler) changeOperand(opPos int, operand int) {
	op := code.Opcode(c.currentInstructions()[opPos])
	newInstruction := code.Make(op, operand)

	c.replaceInstruction(opPos, newInstruction)
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
		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		}

		// Emit an `OpJump` with a bogus value
		jumpPos := c.emit(code.OpJump, 9999)
		afterConsequencePos := len(c.currentInstructions())
		c.changeOperand(jumpNotTruthyPos, afterConsequencePos)

		if node.ELSE() == nil {
			c.emit(code.OpNull)
		} else {
			// else block
			if err := c.Compile(node.CompoundStatement(1)); err != nil {
				return err
			}
			if c.lastInstructionIs(code.OpPop) {
				c.removeLastPop()
			}
		}
		afterElsePos := len(c.currentInstructions())
		c.changeOperand(jumpPos, afterElsePos)

	case *parser.WhileStatementContext:
		startPos := len(c.currentInstructions())
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		// emit on `OpJumpNotTruthy` with a bogus value
		jumpNotTruthyPos := c.emit(code.OpJumpNotTruthy, 9999)

		// main conditional block
		if err := c.Compile(node.CompoundStatement()); err != nil {
			return err
		}
		if c.lastInstructionIs(code.OpPop) {
			c.removeLastPop()
		}

		// Emit an `OpJump` with the startPos
		c.emit(code.OpJump, startPos)
		afterLoopPos := len(c.currentInstructions())
		c.changeOperand(jumpNotTruthyPos, afterLoopPos)
		c.emit(code.OpNull)

	case *parser.CompoundStatementContext:
		for _, s := range node.AllStatement() {
			if err := c.Compile(s); err != nil {
				return err
			}
		}

	case *parser.FunctionExpressionContext:
		identifierSymbol := c.scopes[c.scopeIndex].currentIdentifier
		c.enterScope()

		if identifierSymbol != nil && identifierSymbol.Name != "" {
			c.symbolTable.DefineFunctionName(identifierSymbol.Name)
		}

		params := node.Params().AllIdentifier()
		for _, p := range params {
			c.symbolTable.Define(p.GetText())
		}

		if err := c.Compile(node.CompoundStatement()); err != nil {
			return err
		}
		if c.lastInstructionIs(code.OpPop) {
			c.replaceLastPopWithReturn()
		}
		if !c.lastInstructionIs(code.OpReturnValue) {
			c.emit(code.OpReturn)
		}

		freeSymbols := c.symbolTable.FreeSymbols
		numLocals := c.symbolTable.numDefinitions
		instructions := c.leaveScope()

		for _, s := range freeSymbols {
			c.loadSymbol(s)
		}

		compiledFn := &object.CompiledFunction{
			Instructions: instructions,
			NumLocals:    numLocals,
			NumParams:    len(params),
		}
		fnIndex := c.addConstant(compiledFn)
		c.emit(code.OpClosure, fnIndex, len(freeSymbols))

	case *parser.ReturnStatementContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		c.emit(code.OpReturnValue)

	case *parser.CallExpressionContext:
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		args := node.ExpressionList().AllExpression()
		for _, arg := range args {
			if err := c.Compile(arg); err != nil {
				return err
			}
		}
		c.emit(code.OpCall, len(args))

	case *parser.ArrayLiteralExpressionContext:
		exprs := node.ExpressionList().AllExpression()
		for _, elem := range exprs {
			if err := c.Compile(elem); err != nil {
				return err
			}
		}
		c.emit(code.OpArray, len(exprs))

	case *parser.HashLiteralExpressionContext:
		pairs := node.AllPair()
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].Expression(0).GetText() < pairs[j].Expression(0).GetText()
		})
		for _, elem := range pairs {
			if err := c.Compile(elem.Expression(0)); err != nil {
				return err
			}
			if err := c.Compile(elem.Expression(1)); err != nil {
				return err
			}
		}
		c.emit(code.OpHash, len(pairs)*2)

	case *parser.IndexExpressionContext:
		if err := c.Compile(node.Expression(0)); err != nil {
			return err
		}
		if err := c.Compile(node.Expression(1)); err != nil {
			return err
		}
		c.emit(code.OpIndex)

	case *parser.StarSlashExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.PlusMinusExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.RelationExpressionContext:
		return c.compileInfixExpression(node)

	case *parser.LiteralExpressionContext:
		return c.compileLiteralExpression(node)

	case *parser.LetStatementContext:
		identifier := node.Identifier().GetText()
		symbol := c.symbolTable.Define(identifier)
		c.scopes[c.scopeIndex].currentIdentifier = &symbol
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		c.scopes[c.scopeIndex].currentIdentifier = nil
		if symbol.Scope == GlobalScope {
			c.emit(code.OpSetGlobal, symbol.Index)
		} else {
			c.emit(code.OpSetLocal, symbol.Index)
		}

	case *parser.AssignStatementContext:
		identifier := node.Identifier().GetText()
		symbol, found := c.symbolTable.Resolve(identifier)
		if !found {
			symbol = c.symbolTable.Define(identifier)
			c.scopes[c.scopeIndex].currentIdentifier = &symbol
		}
		if err := c.Compile(node.Expression()); err != nil {
			return err
		}
		c.scopes[c.scopeIndex].currentIdentifier = nil
		if symbol.Scope == GlobalScope {
			c.emit(code.OpSetGlobal, symbol.Index)
		} else {
			c.emit(code.OpSetLocal, symbol.Index)
		}

	case *parser.IdentifierExpressionContext:
		symbol, ok := c.symbolTable.Resolve(node.Identifier().GetText())
		if !ok {
			return fmt.Errorf("undefined variable %s", node.Identifier().GetText())
		}
		c.loadSymbol(symbol)

	}

	return nil
}

func (c *Compiler) loadSymbol(s Symbol) {
	switch s.Scope {
	case GlobalScope:
		c.emit(code.OpGetGlobal, s.Index)
	case LocalScope:
		c.emit(code.OpGetLocal, s.Index)
	case BuiltinScope:
		c.emit(code.OpGetBuiltin, s.Index)
	case FreeScope:
		c.emit(code.OpGetFree, s.Index)
	case FunctionScope:
		c.emit(code.OpCurrentClosure)
	}
}

func (c *Compiler) emit(op code.Opcode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)

	c.setLastInstruction(op, pos)
	return pos
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
		Instructions: c.currentInstructions(),
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
			return object.NewError("could not parse %s as integer", integerLit.GetSymbol().GetText())
		}
		c.emit(code.OpConstant, c.addConstant(&object.Integer{Value: value}))
		return nil
	}
	floatLit := lit.FloatLiteral()
	if floatLit != nil {
		value, err := strconv.ParseFloat(floatLit.GetSymbol().GetText(), 64)
		if err != nil {
			return object.NewError("could not parse %q as float", floatLit.GetSymbol().GetText())
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
	return object.NewError("unexpected literal: %s", lit.GetText())
}

func (c *Compiler) enterScope() {
	scope := CompilationScope{
		instructions:        code.Instructions{},
		lastInstruction:     EmittedInstruction{},
		previousInstruction: EmittedInstruction{},
	}
	c.scopes = append(c.scopes, scope)
	c.scopeIndex++
	c.symbolTable = NewEnclosedSymbolTable(c.symbolTable)
}

func (c *Compiler) leaveScope() code.Instructions {
	instructions := c.currentInstructions()

	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--

	c.symbolTable = c.symbolTable.Outer

	return instructions
}

func (c *Compiler) replaceLastPopWithReturn() {
	lastPos := c.scopes[c.scopeIndex].lastInstruction.Position
	c.replaceInstruction(lastPos, code.Make(code.OpReturnValue))

	c.scopes[c.scopeIndex].lastInstruction.OpCode = code.OpReturnValue
}
