grammar Gs;

program
    :   ( functionDefinition | structDefinition | statement )+ EOF
    ;

// 结构体定义
structDefinition
    :   'type' ID 'struct' '{' ID (',' ID)* '}'
    ;

// 函数定义
functionDefinition
    :   'func' ID '(' (ID (',' ID)* )? ')'  block
    ;

// 代码块（适配隐藏换行符，用分号或换行分隔语句）
block
    :   '{' (statement ( (SEMICOLON | NEWLINE)+ | EOF )? )* '}'
    ;

// 语句类型（新增自增/自减语句，标注为#incrDecrStmt）
statement
    :   ';'                                       #emptyStmt
    |   structDefinition                          #structStmt
    |   assign                                    #assignStmt
    |   selfAssign                                #selfOpAssignStmt
    |   incrDecr                                  #incrDecrStmt
    |   'return' (expr (',' expr)* )?             #returnStmt
    |   'print' expr (',' expr)*                  #printStmt
    |   'printf' expr (',' expr)*                 #printfStmt
    |   'if' (assign ';')? expr block ('else' block)?  #ifStmt
    |   'for' forInit? ';' expr? ';' forUpdate? block  #forCStyleStmt  // C风格for（无括号）
    |   'for' iterVar '=' 'range' expr block          #forRangeStmt     // 保留原有=，仅修复必须项
    |   'for' expr block                          #forCondStmt      // 条件循环（最后匹配，避免歧义）
    |   call                                      #callStmt
    |   'break'                                   #breakStmt
    |   'continue'                                #continueStmt
    ;

assign: qid (',' qid)* '=' expr (',' expr)*;
incrDecr: qid (INCR | DECR);



// 迭代变量
iterVar
    :   ID                          #singleIter
    |   ID ',' ID                   #doubleIter
    ;

// 循环初始化与更新
forInit : assign;

forUpdate : updateItem ;

selfAssign: qid selfAssignOp expr;

updateItem
    :   selfAssign                        #selfUpdate
    |   incrDecr                          #incrDecrUpdate  // 已有的循环内自增/自减
    |   assign                            #assignUpdate
    ;

// 赋值运算符
selfAssignOp
    :   '+=' | '-=' | '*=' | '/=' | '%='
    ;

// 函数调用（不支持限定标识符，如 a.b()）
call : ID '(' (expr (',' expr)* ','?)? ')' ;  // 支持末尾逗号

// 表达式层级
expr : logicalOrExpr ;

logicalOrExpr : logicalAndExpr (OR logicalAndExpr)* ;
logicalAndExpr : comparisonExpr (AND comparisonExpr)* ;
comparisonExpr : addExpr (compOp addExpr)? ;

addExpr : binExpr (addOp binExpr)* ;
binExpr : mulExpr (bitOp mulExpr)* ;
mulExpr : powExpr (mulOp powExpr)* ;
powExpr : atom (POW atom)* ;

// 原子表达式
atom
    :   SUB atom                #negAtom  // 负数字面量/表达式
    |   INT                     #intAtom
    |   FLOAT                   #floatAtom
    |   STRING                  #stringAtom
    |   TRUE                    #trueAtom
    |   FALSE                   #falseAtom
    |   NIL                     #nilAtom
    |   NOT expr                #notAtom
    |   qid                     #qidAtom
    |   call                    #callAtom
    |   instance                #instanceAtom
    |   arrayLiteral            #arrayAtom
    |   indexAccess             #indexAccessAtom  // 统一数组/字典访问
    |   dictLiteral             #dictAtom
    |   '(' expr ')'            #parenAtom
    ;

// 数组字面量（支持末尾逗号）
arrayLiteral : '[' (expr (',' expr)* ','?)? ']' ;

// 索引访问（合并数组/字典访问，支持切片）
indexAccess : qid '[' (expr | sliceExpr) ']' ;
sliceExpr : expr? COLON expr? ;

// 字典相关（支持末尾逗号）
dictLiteral : '{' (dictEntry (',' dictEntry)* ','?)? '}' ;
dictEntry
    :   (STRING|INT|FLOAT|TRUE|FALSE) ':' expr         #constKeyEntry
    |   qid ':' expr             #idKeyEntry  // 支持qid作为键
    ;

// 结构体实例化（支持末尾逗号）
instance : 'new' ID '{' (ID ':' expr (',' ID ':' expr)* ','?)? '}' ;

// 限定标识符（修正：引用可选链词法规则）
qid : primary ( (DOT | SAFE_DOT ) ID | (LBRACK | SAFE_LBRACK) expr ']' )* ;
primary : ID ;

// 运算符集合
compOp  : EQ | LT | GT | NEQ | GEQ | LEQ ;
addOp   : ADD | SUB ;
bitOp   : BITAND | BITOR | XOR ;
mulOp   : MUL | DIV | MOD ;
powOp   : POW ;

// 关键字（全部放在ID前，利用优先级匹配）
TRUE    : 'true' ;
FALSE   : 'false' ;
NIL     : 'nil' ;
AND     : '&&' ;
OR      : '||' ;
NOT     : '!' ;
IF      : 'if' ;
ELSE    : 'else' ;
PRINT   : 'print' ;
FOR     : 'for' ;
RANGE   : 'range' ;
RETURN  : 'return' ;
FUNC    : 'func' ;
TYPE    : 'type' ;
STRUCT  : 'struct' ;
NEW     : 'new' ;
BREAK   : 'break' ;
CONTINUE: 'continue' ;

// 运算符（必须修复项：补充可选链复合token，优先于单个符号）
SAFE_DOT   : '?.' ;  // 可选链属性访问（新增）
SAFE_LBRACK: '?[' ;  // 可选链索引访问（新增）
ADD        : '+' ;
SUB        : '-' ;
MUL        : '*' ;
DIV        : '/' ;
MOD        : '%' ;
POW        : '**' ;
EQ         : '==' ;
LT         : '<' ;
GT         : '>' ;
GEQ        : '>=' ;
LEQ        : '<=' ;
NEQ        : '!=' ;
BITAND     : '&' ;
BITOR      : '|' ;
XOR        : '^' ;
INCR       : '++' ;  // 自增运算符
DECR       : '--' ;  // 自减运算符
DOT        : '.' ;  // 普通属性访问（在SAFE_DOT之后，避免被拆分）
LBRACK     : '[' ;  // 普通索引访问（在SAFE_LBRACK之后，避免被拆分）
LPAREN     : '(' ;
RPAREN     : ')' ;
LBRACE     : '{' ;
RBRACE     : '}' ;
RBRACK     : ']' ;
COLON      : ':' ;
SEMICOLON  : ';' ;
COMMA      : ',' ;

// 整数字面量（对齐Go风格：支持二进制、八进制、十进制、十六进制，下划线分隔）
INT
    :   '0' [bB] [01]+ ( '_' [01]+ )*            // 二进制（0b101_000, 0B111_111）
    |   '0' [oO] [0-7]+ ( '_' [0-7]+ )*          // 八进制（0o755_333, 0O123_456）
    |   '0' [xX] [0-9a-fA-F]+ ( '_' [0-9a-fA-F]+ )*  // 十六进制（0x1a_3f, 0XaB_Cd）
    |   [0-9]+ ( '_' [0-9]+ )*                  // 十进制（123_456, 987_654）
    ;

// 浮点数字面量（对齐Go风格：支持小数点、指数形式，下划线分隔）
FLOAT
    :   [0-9]+ ( '_' [0-9]+ )* '.' [0-9]* ( '_' [0-9]+ )* ([eE][+-]? [0-9]+ ( '_' [0-9]+ )*)?  // 123_456.789, 123.45_6e-7, 123.
    |   '.' [0-9]+ ( '_' [0-9]+ )* ([eE][+-]? [0-9]+ ( '_' [0-9]+ )*)?  // .45_6e+8, .789
    |   [0-9]+ ( '_' [0-9]+ )* [eE][+-]? [0-9]+ ( '_' [0-9]+ )*  // 123_456e-3, 789E+10
    ;

// 字符串字面量（对齐Go风格：双引号单行字符串+反引号原始字符串）
STRING
    :   '"' ( ESC | ~["\\] )* '"'                // 单行字符串（支持转义）
    |   '`' ( ~'`' )* '`'                        // 原始字符串（多行，不解析转义）
    ;

// 转义序列（对齐Go支持的转义，修复长度限制：u需要4位，U需要8位，x需要2位）
fragment ESC
    : '\\' [abfnrtv'\\"]                         // 基础转义（\a, \b, \n等）
    | '\\' 'u' [0-9a-fA-F]                    // Unicode 4位转义（\u0041）
    | '\\' 'U' [0-9a-fA-F]                    // Unicode 8位转义（\U00000041）
    | '\\' 'x' [0-9a-fA-F]                    // 十六进制转义（\x41）
    ;

// 空白与注释（换行符隐藏）
WS         : [ \t]+ -> channel(HIDDEN) ;
NEWLINE    : '\r'? '\n' -> channel(HIDDEN);
SL_COMMENT : '//' ~[\r\n]* -> channel(HIDDEN) ;
ML_COMMENT : '/*' .*? '*/' -> channel(HIDDEN) ;

// 标识符（在关键字后，自动排除关键字）
ID      : [a-zA-Z_] [a-zA-Z0-9_]* ;