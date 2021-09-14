grammar Calc;

NUMBER: DECIMALNUMBER | OCTONARYNUMBER | HEXADECIMALNUMBER;
DECIMALNUMBER: [0-9][0-9_]*;
OCTONARYNUMBER: '0' [oO] [0-9]+;
HEXADECIMALNUMBER: '0' [xX] [0-9a-fA-F]+;
WHITESPACE : [ \t\r\n\u000C]+ -> skip;

prog
    : expr
    | (expr WHITESPACE)* ;
expr:	expr ('*'|'/') expr
    |	expr ('+'|'-') expr
    |	NUMBER
    |	'(' expr ')'
    ;
