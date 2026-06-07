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

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterPair is called when entering the pair production.
	EnterPair(c *PairContext)

	// EnterParams is called when entering the params production.
	EnterParams(c *ParamsContext)

	// EnterLet is called when entering the let production.
	EnterLet(c *LetContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitLetStatement is called when exiting the letStatement production.
	ExitLetStatement(c *LetStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

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

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitPair is called when exiting the pair production.
	ExitPair(c *PairContext)

	// ExitParams is called when exiting the params production.
	ExitParams(c *ParamsContext)

	// ExitLet is called when exiting the let production.
	ExitLet(c *LetContext)
}
