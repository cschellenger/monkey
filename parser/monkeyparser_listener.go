// Code generated from MonkeyParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MonkeyParser
import "github.com/antlr4-go/antlr/v4"

// MonkeyParserListener is a complete listener for a parse tree produced by MonkeyParser.
type MonkeyParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterLetStatement is called when entering the letStatement production.
	EnterLetStatement(c *LetStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIdentifierExpression is called when entering the identifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterNegatedExpression is called when entering the negatedExpression production.
	EnterNegatedExpression(c *NegatedExpressionContext)

	// EnterRelationExpression is called when entering the relationExpression production.
	EnterRelationExpression(c *RelationExpressionContext)

	// EnterParenExpression is called when entering the parenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterStarSlashExpression is called when entering the starSlashExpression production.
	EnterStarSlashExpression(c *StarSlashExpressionContext)

	// EnterCallExpression is called when entering the callExpression production.
	EnterCallExpression(c *CallExpressionContext)

	// EnterArrayLiteralExpression is called when entering the arrayLiteralExpression production.
	EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// EnterPlusMinusExpression is called when entering the plusMinusExpression production.
	EnterPlusMinusExpression(c *PlusMinusExpressionContext)

	// EnterHashLiteralExpression is called when entering the hashLiteralExpression production.
	EnterHashLiteralExpression(c *HashLiteralExpressionContext)

	// EnterIndexExpression is called when entering the indexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterIfExpression is called when entering the ifExpression production.
	EnterIfExpression(c *IfExpressionContext)

	// EnterLiteralExpression is called when entering the literalExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterExpression_list is called when entering the expression_list production.
	EnterExpression_list(c *Expression_listContext)

	// EnterExpression_pair is called when entering the expression_pair production.
	EnterExpression_pair(c *Expression_pairContext)

	// EnterArray_literal is called when entering the array_literal production.
	EnterArray_literal(c *Array_literalContext)

	// EnterHash_literal is called when entering the hash_literal production.
	EnterHash_literal(c *Hash_literalContext)

	// EnterFunction_literal is called when entering the function_literal production.
	EnterFunction_literal(c *Function_literalContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterIf_expression is called when entering the if_expression production.
	EnterIf_expression(c *If_expressionContext)

	// EnterCall_expression is called when entering the call_expression production.
	EnterCall_expression(c *Call_expressionContext)

	// EnterLet_statement is called when entering the let_statement production.
	EnterLet_statement(c *Let_statementContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitLetStatement is called when exiting the letStatement production.
	ExitLetStatement(c *LetStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIdentifierExpression is called when exiting the identifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitNegatedExpression is called when exiting the negatedExpression production.
	ExitNegatedExpression(c *NegatedExpressionContext)

	// ExitRelationExpression is called when exiting the relationExpression production.
	ExitRelationExpression(c *RelationExpressionContext)

	// ExitParenExpression is called when exiting the parenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitStarSlashExpression is called when exiting the starSlashExpression production.
	ExitStarSlashExpression(c *StarSlashExpressionContext)

	// ExitCallExpression is called when exiting the callExpression production.
	ExitCallExpression(c *CallExpressionContext)

	// ExitArrayLiteralExpression is called when exiting the arrayLiteralExpression production.
	ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// ExitPlusMinusExpression is called when exiting the plusMinusExpression production.
	ExitPlusMinusExpression(c *PlusMinusExpressionContext)

	// ExitHashLiteralExpression is called when exiting the hashLiteralExpression production.
	ExitHashLiteralExpression(c *HashLiteralExpressionContext)

	// ExitIndexExpression is called when exiting the indexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitIfExpression is called when exiting the ifExpression production.
	ExitIfExpression(c *IfExpressionContext)

	// ExitLiteralExpression is called when exiting the literalExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitExpression_list is called when exiting the expression_list production.
	ExitExpression_list(c *Expression_listContext)

	// ExitExpression_pair is called when exiting the expression_pair production.
	ExitExpression_pair(c *Expression_pairContext)

	// ExitArray_literal is called when exiting the array_literal production.
	ExitArray_literal(c *Array_literalContext)

	// ExitHash_literal is called when exiting the hash_literal production.
	ExitHash_literal(c *Hash_literalContext)

	// ExitFunction_literal is called when exiting the function_literal production.
	ExitFunction_literal(c *Function_literalContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitIf_expression is called when exiting the if_expression production.
	ExitIf_expression(c *If_expressionContext)

	// ExitCall_expression is called when exiting the call_expression production.
	ExitCall_expression(c *Call_expressionContext)

	// ExitLet_statement is called when exiting the let_statement production.
	ExitLet_statement(c *Let_statementContext)
}
