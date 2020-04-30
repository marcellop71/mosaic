grammar Policy;

policy:
  expr
| sexpr
;

expr:
    expr op=(TK_AND|TK_OR) expr # expr_andor
  | TK_LPAREN expr TK_RPAREN # expr_paren
  | attr op=(TK_EQ|TK_LT|TK_LTEQ|TK_GT|TK_GTEQ) value # expr_attrrange0
  | attr TK_LBRACKET value TK_COMMA value TK_LBRACKET # expr_attrrange1
  | attr # expr_attr
;

sexpr:
    TK_LPAREN op=(TK_AND|TK_OR) sexpr+ TK_RPAREN # sexpr_andor
  | op=(TK_EQ|TK_LT|TK_LTEQ|TK_GT|TK_GTEQ) attr value # sexpr_attrrange0
  | attr TK_LBRACKET value TK_COMMA value TK_LBRACKET # sexpr_attrrange1
  | attr # sexpr_attr
;

attr:
  Name TK_AT Name
;

value:
  Number
;

Name:
  (Letter | SpecialCharacter)+ (Letter | Digit | SpecialCharacter)*
;

Number:
  Digit+
;

WS: [ \t\n\r]+ -> skip;

TK_EQ: '==' ;
TK_NEQ: '!=' ;
TK_LT: '<' ;
TK_LTEQ: '<=' ;
TK_GT: '>' ;
TK_GTEQ: '>=' ;
TK_AT: '@' ;
TK_AND: '/\\' ;
TK_OR: '\\/' ;
TK_COMMA: ',' ;
TK_LBRACKET: '[' ;
TK_RBRACKET: ']' ;
TK_LPAREN: '(' ;
TK_RPAREN: ')' ;

fragment SpecialCharacter:
  '$' | '£' | '§' | '#' | '.' | '&' | '%' | '-'
;

fragment Letter: LetterUppercase | LetterLowercase;
fragment LetterUppercase: 'A'  ..  'Z';
fragment LetterLowercase: 'a' .. 'z';
fragment Digit: '0' .. '9';
fragment NonZeroDigit: '1' .. '9';
fragment ZeroDigit: '0';
