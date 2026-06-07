// Code generated from MonkeyParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MonkeyParser
import "github.com/antlr4-go/antlr/v4"

type BaseMonkeyParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMonkeyParserVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLetStatement(ctx *LetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitNegatedExpression(ctx *NegatedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitRelationExpression(ctx *RelationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitParenExpression(ctx *ParenExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitStarSlashExpression(ctx *StarSlashExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitPlusMinusExpression(ctx *PlusMinusExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitHashLiteralExpression(ctx *HashLiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIndexExpression(ctx *IndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIfExpression(ctx *IfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLiteralExpression(ctx *LiteralExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpression_list(ctx *Expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitExpression_pair(ctx *Expression_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitArray_literal(ctx *Array_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitHash_literal(ctx *Hash_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitFunction_literal(ctx *Function_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitParams(ctx *ParamsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitIf_expression(ctx *If_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitCall_expression(ctx *Call_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMonkeyParserVisitor) VisitLet_statement(ctx *Let_statementContext) interface{} {
	return v.VisitChildren(ctx)
}
