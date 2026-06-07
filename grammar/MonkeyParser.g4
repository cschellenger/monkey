parser grammar MonkeyParser;


options { 
    tokenVocab = MonkeyLexer; 
}

prog : statement+ ;

statement
    : RETURN expression ';'                                 # returnStatement
    | LET Identifier ASSIGN expression ';'                  # letStatement
    | WHILE '(' expression ')' '{' compoundStatement '}'    # whileStatement
    | expression ';'?                                       # expressionStatement
    ;


expression
    : 'if' '(' expression ')' '{' compoundStatement '}' ('else' '{' compoundStatement '}')? # ifExpression
    | 'fn' '(' params ')' '{' compoundStatement '}'                                         # functionExpression
    | identifier '(' expressionList ')'                                                     # callExpression
    | expression '[' expression ']'                                                         # indexExpression
    | identifier                                                                            # identifierExpression
    | literal                                                                               # literalExpression
    | '[' expressionList ']'                                                                # arrayLiteralExpression
    | '{' (pair ',')* (pair ','?)? '}'                                                      # hashLiteralExpression
    | expression op=(STAR|SLASH|PERCENT) expression                                         # starSlashExpression
    | expression op=(PLUS|MINUS) expression                                                 # plusMinusExpression
    | expression op=(GT|LT|GE|LE|EQ|NOT_EQ) expression                                      # relationExpression
    | '(' expression ')'                                                                    # parenExpression
    | '!' expression                                                                        # negatedExpression
    ;

identifier : Identifier;

literal
    : IntegerLiteral
    | FloatLiteral
    | StringLiteral
    | BooleanLiteral
    ;

compoundStatement: (statement ';'?)*;

expressionList: (expression ',')* (expression ','?)?;

pair: expression ':' expression;

params
    : (Identifier ',')* (Identifier ','?)?;

let
    : LET identifier op=ASSIGN expression END
    ;

