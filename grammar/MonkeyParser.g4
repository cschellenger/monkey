parser grammar MonkeyParser;


options { 
    tokenVocab = MonkeyLexer; 
}

prog : statement+ ;

statement
    : RETURN expression ';'                 # returnStatement
    | LET Identifier ASSIGN expression ';'  # letStatement
    | WHILE '(' expression ')' statement    # whileStatement
    | '{' (statement ';'?)* '}'             # compoundStatement
    | expression ';'?                       # expressionStatement
    ;


expression
    : if_expression                                     # ifExpression
    | function_literal                                  # functionExpression
    | call_expression                                   # callExpression
    | expression '[' expression ']'                     # indexExpression
    | identifier                                        # identifierExpression
    | literal                                           # literalExpression
    | array_literal                                     # arrayLiteralExpression
    | hash_literal                                      # hashLiteralExpression
    | expression op=(STAR|SLASH|PERCENT) expression     # starSlashExpression
    | expression op=(PLUS|MINUS) expression             # plusMinusExpression
    | expression op=(GT|LT|GE|LE|EQ|NOT_EQ) expression  # relationExpression
    | '(' expression ')'                                # parenExpression
    | '!' expression                                    # negatedExpression
    ;

identifier : Identifier;

literal
    : IntegerLiteral
    | FloatLiteral
    | StringLiteral
    | BooleanLiteral
    ;

expression_list: (expression ',')* (expression ','?)?;

expression_pair: expression ':' expression;

array_literal: '[' expression_list ']';

hash_literal
    : '{' (expression_pair ',')* (expression_pair ','?)? '}';

function_literal
    : 'fn' '(' params ')' statement;

params
    : (Identifier ',')* (Identifier ','?)?;

if_expression
    : 'if' '(' expression ')' statement ('else' statement)?;

call_expression
    : identifier '(' expression_list ')';

let_statement
    : LET identifier op=ASSIGN expression END
    ;

