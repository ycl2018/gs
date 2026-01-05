grammar Gs;

program
    :   ( functionDefinition | structDefinition | statement )+ EOF
    ;

structDefinition
    :   TYPE ID STRUCT '{' ID (',' ID)* '}'
    ;

functionDefinition
    :   FUNC ID '(' (ID (',' ID)* )? ')'  block
    ;

block
    :   '{' (statement ( (SEMICOLON | NEWLINE)+ | EOF )? )* '}'
    ;

statement
    :   ';'                                             #emptyStmt
    |   structDefinition                                #structStmt
    |   assign                                          #assignStmt
    |   selfAssign                                      #selfOpAssignStmt
    |   incrDecr                                        #incrDecrStmt
    |   RETURN (expr (',' expr)* )?                     #returnStmt
    |   ifStatement                                     #ifStmt
    |   FOR forInit? ';' expr? ';' forUpdate? block     #forCStyleStmt
    |   FOR (iterVar '=')? RANGE expr block             #forRangeStmt
    |   FOR expr block                                  #forCondStmt
    |   SWITCH (expr? | simpleStmt? ';' expr?)
        '{' exprCaseClause* '}'                         #switchStmt
    |   builtinCall                                     #builtinStmt
    |   call                                            #callStmt
    |   BREAK                                           #breakStmt
    |   CONTINUE                                        #continueStmt
    |   GLOBAL ID (',' ID)*                             #globalStmt
    ;

ifStatement:
    IF ( simpleStmt ';')? expr block (ELSE (block | ifStatement))?;

simpleStmt:
    assign | selfAssign | incrDecr | builtinCall | call;

assign:
    lvalue (',' lvalue)* '=' expr (',' expr)*;

incrDecr:
    lvalue (INCR | DECR);

exprCaseClause
    : exprSwitchCase COLON statement*
    ;

exprSwitchCase
    : CASE expr (',' expr)*
    | DEFAULT
    ;

builtinCall
    :   LEN '(' expr ')'                                                #lenCall
    |   INITREF '(' qid (',' expr)? ')'                                 #initRefCall
    |   NEWFROMTYPE '(' expr ')'                                        #newFromTypeCall
    |   APPEND '(' expr  ((',' expr)* |',' expr EXPAND) ')'             #appendCall
    |   DELETE '(' expr ',' expr ')'                                    #deleteCall
    |   (PRINT|PRINTLN|PRINTF|SPRINTF) '(' expr (',' expr)* ')'         #printXCall
    |   (UINT|UINT8|UINT16|UINT32|UINT64|INTS|INT8|
        INT16|INT32|INT64|FLOAT32|FLOAT64|STRINGS|BOOL) '(' expr ')'    #convertCall
    |   GO call                                                         #goCall
    ;

iterVar
    :   ID                          #singleIter
    |   ID ',' ID                   #doubleIter
    ;

forInit :
    assign;

forUpdate :
    updateItem ;

selfAssign:
    lvalue selfAssignOp expr;

updateItem
    :   selfAssign                        #selfUpdate
    |   incrDecr                          #incrDecrUpdate
    |   assign                            #assignUpdate
    ;

selfAssignOp
    :   '+=' | '-=' | '*=' | '/=' | '%='
    ;

call
    : ID '(' (expr (',' expr)* ','?)? ')'                   #innerCall
    | primary accessor+ '(' (expr (',' expr)* ','?)? ')'    #outerCall
    ;

expr
    : atom                          #atomExpr
    | expr mulOp expr               #mulExpr
    | expr addOp expr               #addExpr
    | expr compOp expr              #comparisonExpr
    |<assoc=right> expr AND expr    #logicalAndExpr
    |<assoc=right> expr OR expr     #logicalOrExpr
    ;

atom
    :   SUB atom                #negAtom
    |   NOT atom                #notAtom
    |   '*' lvalue              #derefAtom
    |   INT                     #intAtom
    |   FLOAT                   #floatAtom
    |   STRING                  #stringAtom
    |   TRUE                    #trueAtom
    |   FALSE                   #falseAtom
    |   NIL                     #nilAtom
    |   builtinCall             #builtinAtom
    |   call                    #callAtom
    |   instance                #instanceAtom
    |   arrayLiteral            #arrayAtom
    |   dictLiteral             #dictAtom
    |   '(' expr ')'            #parenAtom
    |   qid                     #qidAtom
    ;

lvalue
    : qid
    | '*' lvalue
    ;

arrayLiteral :
    '[' (expr (',' expr)* ','?)? ']' ;

sliceExpr :
    expr? COLON expr? ;

dictLiteral :
    '{' (dictEntry (',' dictEntry)* ','?)? '}' ;

dictEntry
    :   (STRING|INT|FLOAT|TRUE|FALSE) ':' expr         #constKeyEntry
    |   lvalue ':' expr                                #idKeyEntry
    ;

instance :
    NEW ID '{' (ID ':' expr (',' ID ':' expr)* ','?)? '}' ;

qid :
    primary accessor* ;

accessor
    :   DOT ID                            #propertyAccess
    |   LBRACK (expr | sliceExpr) ']'     #indexAccess
    |   '(' (expr (',' expr)* ','?)? ')'  #methodCallAccess
    ;

primary : ID | ENV ;

compOp  : EQ | LT | GT | NEQ | GEQ | LEQ ;
addOp   : ADD | SUB | BITOR | XOR ;
mulOp   : MUL | DIV | MOD | LSHIFT | RSHIFT | BITAND;

ENV     : '$' ;
TRUE    : 'true' ;
FALSE   : 'false' ;
NIL     : 'nil' ;
AND     : '&&' ;
OR      : '||' ;
NOT     : '!' ;
IF      : 'if' ;
ELSE    : 'else' ;
FOR     : 'for' ;
RANGE   : 'range' ;
RETURN  : 'return' ;
FUNC    : 'func' ;
TYPE    : 'type' ;
STRUCT  : 'struct' ;
NEW     : 'new' ;
BREAK   : 'break' ;
CONTINUE: 'continue' ;
GLOBAL  : 'global' ;
CASE    : 'case' ;
DEFAULT : 'default' ;
SWITCH  : 'switch' ;

GO      : 'go';
LEN     : 'len' ;
APPEND  : 'append' ;
DELETE  : 'delete' ;
COPY    : 'copy' ;
TOSTRING: 'toString' ;
PRINT   : 'print' ;
PRINTF  : 'printf';
SPRINTF : 'sprintf';
PRINTLN : 'println';
INITREF : 'initRef';
NEWFROMTYPE : 'newFromType';

// ==== convert
UINT8   : 'uint8';
UINT16  : 'uint16';
UINT32  : 'uint32';
UINT64  : 'uint64';
UINT    : 'uint';
INT8    : 'int8';
INT16   : 'int16';
INT32   : 'int32';
INT64   : 'int64';
INTS    : 'int';
FLOAT32 : 'float32';
FLOAT64 : 'float64';
STRINGS : 'string';
BOOL    : 'bool';
EXPAND  : '...';



// 运算符
LSHIFT     : '<<' ;
RSHIFT     : '>>' ;
INCR       : '++' ;
DECR       : '--' ;
GEQ        : '>=' ;
LEQ        : '<=' ;
NEQ        : '!=' ;
EQ         : '==' ;
ADD        : '+' ;
SUB        : '-' ;
MUL        : '*' ;
DIV        : '/' ;
MOD        : '%' ;
LT         : '<' ;
GT         : '>' ;
BITAND     : '&' ;
BITOR      : '|' ;
XOR        : '^' ;
DOT        : '.' ;
LBRACK     : '[' ;
LPAREN     : '(' ;
RPAREN     : ')' ;
LBRACE     : '{' ;
RBRACE     : '}' ;
RBRACK     : ']' ;
COLON      : ':' ;
SEMICOLON  : ';' ;
COMMA      : ',' ;

INT
    :   '0' [bB] [01]+ ( '_' [01]+ )*
    |   '0' [oO] [0-7]+ ( '_' [0-7]+ )*
    |   '0' [xX] [0-9a-fA-F]+ ( '_' [0-9a-fA-F]+ )*
    |   [0-9]+ ( '_' [0-9]+ )*
    ;

FLOAT
    :   [0-9]+ ( '_' [0-9]+ )* '.' [0-9]* ( '_' [0-9]+ )* ([eE][+-]? [0-9]+ ( '_' [0-9]+ )*)?  // 123_456.789, 123.45_6e-7, 123.
    |   '.' [0-9]+ ( '_' [0-9]+ )* ([eE][+-]? [0-9]+ ( '_' [0-9]+ )*)?  // .45_6e+8, .789
    |   [0-9]+ ( '_' [0-9]+ )* [eE][+-]? [0-9]+ ( '_' [0-9]+ )*  // 123_456e-3, 789E+10
    ;

STRING
    :   '"' ( ESC | ~["\\] )* '"'
    |   '`' ( ~'`' )* '`'
    ;

fragment ESC
    : '\\' [abfnrtv'\\"]
    | '\\' 'u' [0-9a-fA-F]
    | '\\' 'U' [0-9a-fA-F]
    | '\\' 'x' [0-9a-fA-F]
    ;

WS         : [ \t]+ -> channel(HIDDEN) ;
NEWLINE    : '\r'? '\n' -> channel(HIDDEN);
SL_COMMENT : '//' ~[\r\n]* -> channel(HIDDEN) ;
ML_COMMENT : '/*' .*? '*/' -> channel(HIDDEN) ;

ID      : [a-zA-Z_] [a-zA-Z0-9_]* ;
