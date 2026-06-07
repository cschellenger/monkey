// Code generated from MonkeyParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MonkeyParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by MonkeyParser.
type MonkeyParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MonkeyParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by MonkeyParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#letStatement.
	VisitLetStatement(ctx *LetStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) interface{}

	// Visit a parse tree produced by MonkeyParser#identifierExpression.
	VisitIdentifierExpression(ctx *IdentifierExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#negatedExpression.
	VisitNegatedExpression(ctx *NegatedExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#relationExpression.
	VisitRelationExpression(ctx *RelationExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#parenExpression.
	VisitParenExpression(ctx *ParenExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#functionExpression.
	VisitFunctionExpression(ctx *FunctionExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#starSlashExpression.
	VisitStarSlashExpression(ctx *StarSlashExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#callExpression.
	VisitCallExpression(ctx *CallExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#arrayLiteralExpression.
	VisitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#plusMinusExpression.
	VisitPlusMinusExpression(ctx *PlusMinusExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hashLiteralExpression.
	VisitHashLiteralExpression(ctx *HashLiteralExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#indexExpression.
	VisitIndexExpression(ctx *IndexExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#ifExpression.
	VisitIfExpression(ctx *IfExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#literalExpression.
	VisitLiteralExpression(ctx *LiteralExpressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MonkeyParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expression_list.
	VisitExpression_list(ctx *Expression_listContext) interface{}

	// Visit a parse tree produced by MonkeyParser#expression_pair.
	VisitExpression_pair(ctx *Expression_pairContext) interface{}

	// Visit a parse tree produced by MonkeyParser#array_literal.
	VisitArray_literal(ctx *Array_literalContext) interface{}

	// Visit a parse tree produced by MonkeyParser#hash_literal.
	VisitHash_literal(ctx *Hash_literalContext) interface{}

	// Visit a parse tree produced by MonkeyParser#function_literal.
	VisitFunction_literal(ctx *Function_literalContext) interface{}

	// Visit a parse tree produced by MonkeyParser#params.
	VisitParams(ctx *ParamsContext) interface{}

	// Visit a parse tree produced by MonkeyParser#if_expression.
	VisitIf_expression(ctx *If_expressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#call_expression.
	VisitCall_expression(ctx *Call_expressionContext) interface{}

	// Visit a parse tree produced by MonkeyParser#let_statement.
	VisitLet_statement(ctx *Let_statementContext) interface{}
}
