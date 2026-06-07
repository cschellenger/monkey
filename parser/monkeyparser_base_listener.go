// Code generated from MonkeyParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MonkeyParser
import "github.com/antlr4-go/antlr/v4"

// BaseMonkeyParserListener is a complete listener for a parse tree produced by MonkeyParser.
type BaseMonkeyParserListener struct{}

var _ MonkeyParserListener = &BaseMonkeyParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMonkeyParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMonkeyParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMonkeyParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMonkeyParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseMonkeyParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseMonkeyParserListener) ExitProg(ctx *ProgContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseMonkeyParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseMonkeyParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterLetStatement is called when production letStatement is entered.
func (s *BaseMonkeyParserListener) EnterLetStatement(ctx *LetStatementContext) {}

// ExitLetStatement is called when production letStatement is exited.
func (s *BaseMonkeyParserListener) ExitLetStatement(ctx *LetStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseMonkeyParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseMonkeyParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseMonkeyParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseMonkeyParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIdentifierExpression is called when production identifierExpression is entered.
func (s *BaseMonkeyParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production identifierExpression is exited.
func (s *BaseMonkeyParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterNegatedExpression is called when production negatedExpression is entered.
func (s *BaseMonkeyParserListener) EnterNegatedExpression(ctx *NegatedExpressionContext) {}

// ExitNegatedExpression is called when production negatedExpression is exited.
func (s *BaseMonkeyParserListener) ExitNegatedExpression(ctx *NegatedExpressionContext) {}

// EnterRelationExpression is called when production relationExpression is entered.
func (s *BaseMonkeyParserListener) EnterRelationExpression(ctx *RelationExpressionContext) {}

// ExitRelationExpression is called when production relationExpression is exited.
func (s *BaseMonkeyParserListener) ExitRelationExpression(ctx *RelationExpressionContext) {}

// EnterParenExpression is called when production parenExpression is entered.
func (s *BaseMonkeyParserListener) EnterParenExpression(ctx *ParenExpressionContext) {}

// ExitParenExpression is called when production parenExpression is exited.
func (s *BaseMonkeyParserListener) ExitParenExpression(ctx *ParenExpressionContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseMonkeyParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseMonkeyParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterStarSlashExpression is called when production starSlashExpression is entered.
func (s *BaseMonkeyParserListener) EnterStarSlashExpression(ctx *StarSlashExpressionContext) {}

// ExitStarSlashExpression is called when production starSlashExpression is exited.
func (s *BaseMonkeyParserListener) ExitStarSlashExpression(ctx *StarSlashExpressionContext) {}

// EnterCallExpression is called when production callExpression is entered.
func (s *BaseMonkeyParserListener) EnterCallExpression(ctx *CallExpressionContext) {}

// ExitCallExpression is called when production callExpression is exited.
func (s *BaseMonkeyParserListener) ExitCallExpression(ctx *CallExpressionContext) {}

// EnterArrayLiteralExpression is called when production arrayLiteralExpression is entered.
func (s *BaseMonkeyParserListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// ExitArrayLiteralExpression is called when production arrayLiteralExpression is exited.
func (s *BaseMonkeyParserListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {}

// EnterPlusMinusExpression is called when production plusMinusExpression is entered.
func (s *BaseMonkeyParserListener) EnterPlusMinusExpression(ctx *PlusMinusExpressionContext) {}

// ExitPlusMinusExpression is called when production plusMinusExpression is exited.
func (s *BaseMonkeyParserListener) ExitPlusMinusExpression(ctx *PlusMinusExpressionContext) {}

// EnterHashLiteralExpression is called when production hashLiteralExpression is entered.
func (s *BaseMonkeyParserListener) EnterHashLiteralExpression(ctx *HashLiteralExpressionContext) {}

// ExitHashLiteralExpression is called when production hashLiteralExpression is exited.
func (s *BaseMonkeyParserListener) ExitHashLiteralExpression(ctx *HashLiteralExpressionContext) {}

// EnterIndexExpression is called when production indexExpression is entered.
func (s *BaseMonkeyParserListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production indexExpression is exited.
func (s *BaseMonkeyParserListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterIfExpression is called when production ifExpression is entered.
func (s *BaseMonkeyParserListener) EnterIfExpression(ctx *IfExpressionContext) {}

// ExitIfExpression is called when production ifExpression is exited.
func (s *BaseMonkeyParserListener) ExitIfExpression(ctx *IfExpressionContext) {}

// EnterLiteralExpression is called when production literalExpression is entered.
func (s *BaseMonkeyParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production literalExpression is exited.
func (s *BaseMonkeyParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMonkeyParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMonkeyParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMonkeyParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMonkeyParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseMonkeyParserListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseMonkeyParserListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseMonkeyParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseMonkeyParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterPair is called when production pair is entered.
func (s *BaseMonkeyParserListener) EnterPair(ctx *PairContext) {}

// ExitPair is called when production pair is exited.
func (s *BaseMonkeyParserListener) ExitPair(ctx *PairContext) {}

// EnterParams is called when production params is entered.
func (s *BaseMonkeyParserListener) EnterParams(ctx *ParamsContext) {}

// ExitParams is called when production params is exited.
func (s *BaseMonkeyParserListener) ExitParams(ctx *ParamsContext) {}

// EnterLet is called when production let is entered.
func (s *BaseMonkeyParserListener) EnterLet(ctx *LetContext) {}

// ExitLet is called when production let is exited.
func (s *BaseMonkeyParserListener) ExitLet(ctx *LetContext) {}
