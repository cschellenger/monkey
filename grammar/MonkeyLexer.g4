lexer grammar MonkeyLexer;

FN : 'fn';
IF : 'if';
ELSE : 'else';
WHILE : 'while';
LET : 'let';
RETURN : 'return';
ASSIGN : '=';
PLUS : '+';
MINUS : '-';
BANG : '!';
STAR : '*';
SLASH : '/';
PERCENT : '%';
END : ';';
LT : '<';
GT : '>';
GE : '>=';
LE : '<=';
EQ : '==';
NOT_EQ : '!=';
COMMA : ',';
COLON : ':';

fragment DIGIT: [0-9];
fragment ALPHA: [a-zA-Z_];
fragment TRUE : 'true';
fragment FALSE : 'false';

BooleanLiteral : TRUE | FALSE;
StringLiteral : '"' ~ ["\r\n]* '"' ;
Identifier : ALPHA ( ALPHA | DIGIT )*;
FloatLiteral : MINUS? DIGIT+ '.' DIGIT+;
IntegerLiteral : MINUS? DIGIT+;

LPAREN : '(';
RPAREN : ')';
LBRACE : '{';
RBRACE : '}';
LBRACKET : '[';
RBRACKET : ']';

WS
    : [ \t\n\r]+ -> skip
    ;

