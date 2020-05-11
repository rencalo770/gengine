grammar gengine;

primary: ruleEntity+;

ruleEntity:  RULE ruleName ruleDescription? salience? BEGIN ruleContent END;
ruleName : stringLiteral;
ruleDescription : stringLiteral;
salience : SALIENCE integer;
ruleContent : statements;
statements: statement+;

statement : ifStmt | methodCall  | functionCall | assignment;

expression : mathExpression
            | expression comparisonOperator expression
            | expression logicalOperator expression
            | notOperator ? expressionAtom
            | notOperator ? '(' expression  ')'
            ;

mathExpression : mathExpression  mathMdOperator mathExpression
               | mathExpression  mathPmOperator mathExpression
               | expressionAtom
               | '(' mathExpression ')'
               ;

expressionAtom
    : methodCall
    | functionCall
    | constant
    | mapVar
    | variable
    ;

assignment : (mapVar | variable) (assignOperator | setOperator)  mathExpression;

ifStmt : 'if' expression '{' statements? '}' elseStmt? ;

elseStmt : 'else' '{' statements? '}';

constant
    : booleanLiteral
    | integer
    | realLiteral
    | stringLiteral
    | atName
    ;

functionArgs
    : (constant | variable  | functionCall | methodCall | mapVar)  (','(constant | variable | functionCall | methodCall | mapVar))*
    ;

integer : '-'? INT;

realLiteral : MINUS? REAL_LITERAL;

stringLiteral: DQUOTA_STRING ;

booleanLiteral : TRUE | FALSE;

functionCall : SIMPLENAME '(' functionArgs? ')';

methodCall : DOTTEDNAME '(' functionArgs? ')';

variable :  SIMPLENAME | DOTTEDNAME ;

mathPmOperator : PLUS | MINUS ;

mathMdOperator : MUL | DIV ;

comparisonOperator : GT | LT | GTE | LTE | EQUALS | NOTEQUALS ;

logicalOperator : AND | OR ;

assignOperator: ASSIGN ;

setOperator: SET;

notOperator: NOT;

mapVar: variable LSQARE (integer |stringLiteral | variable ) RSQARE;

atName : '@name';

fragment DEC_DIGIT          : [0-9];
fragment A                  : [aA] ;
fragment B                  : [bB] ;
fragment C                  : [cC] ;
fragment D                  : [dD] ;
fragment E                  : [eE] ;
fragment F                  : [fF] ;
fragment G                  : [gG] ;
fragment H                  : [hH] ;
fragment I                  : [iI] ;
fragment J                  : [jJ] ;
fragment K                  : [kK] ;
fragment L                  : [lL] ;
fragment M                  : [mM] ;
fragment N                  : [nN] ;
fragment O                  : [oO] ;
fragment P                  : [pP] ;
fragment Q                  : [qQ] ;
fragment R                  : [rR] ;
fragment S                  : [sS] ;
fragment T                  : [tT] ;
fragment U                  : [uU] ;
fragment V                  : [vV] ;
fragment W                  : [wW] ;
fragment X                  : [xX] ;
fragment Y                  : [yY] ;
fragment Z                  : [zZ] ;
fragment EXPONENT_NUM_PART  : 'E' '-'? DEC_DIGIT+;

NIL                         : N I L;
RULE                        : R U L E  ;
AND                         : '&&' ;
OR                          : '||' ;

TRUE                        : T R U E ;
FALSE                       : F A L S E ;
NULL_LITERAL                : N U L L ;
SALIENCE                    : S A L I E N C E ;
BEGIN                       : B E G I N;
END                         : E N D;

SIMPLENAME :  ('a'..'z' |'A'..'Z')+  ;

INT : '0'..'9' + ;
PLUS                        : '+' ;
MINUS                       : '-' ;
DIV                         : '/' ;
MUL                         : '*' ;

EQUALS                      : '==' ;
GT                          : '>' ;
LT                          : '<' ;
GTE                         : '>=' ;
LTE                         : '<=' ;
NOTEQUALS                   : '!=' ;
NOT                         : '!' ;

ASSIGN                      : ':=' ;
SET                         : '=';

LSQARE    : '[' ;
RSQARE    : ']' ;

SEMICOLON                   : ';' ;
LR_BRACE                    : '{';
RR_BRACE                    : '}';
LR_BRACKET                  : '(';
RR_BRACKET                  : ')';
DOT                         : '.' ;
DQUOTA_STRING               : '"' ( '\\'. | '""' | ~('"'| '\\') )* '"';
DOTTEDNAME                  : SIMPLENAME  DOT SIMPLENAME  ;
REAL_LITERAL                : (DEC_DIGIT+)? '.' DEC_DIGIT+
                            | DEC_DIGIT+ '.' EXPONENT_NUM_PART
                            | (DEC_DIGIT+)? '.' (DEC_DIGIT+ EXPONENT_NUM_PART)
                            | DEC_DIGIT+ EXPONENT_NUM_PART
                            ;

SL_COMMENT: '//' .*? '\n' -> skip ;

WS  :   [ \t\n\r]+ -> skip ;