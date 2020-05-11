lexer grammar PolicyLexer;

Name:
  (Letter | SpecialCharacter)+ (Letter | Digit | SpecialCharacter)*
;

Number:
  Digit+
;

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

TK_LFCR: [\n\r]+ -> skip ;
TK_WS: [ \t]+ -> skip;

fragment SpecialCharacter:
  '$' | '£' | '§' | '#' | '.' | '&' | '%' | '-'
;

fragment Letter: LetterUppercase | LetterLowercase;
fragment LetterUppercase: 'A'  ..  'Z';
fragment LetterLowercase: 'a' .. 'z';
fragment Digit: '0' .. '9';
fragment NonZeroDigit: '1' .. '9';
fragment ZeroDigit: '0';
