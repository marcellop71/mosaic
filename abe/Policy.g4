parser grammar Policy;
options { tokenVocab=PolicyLexer; }

@members {}

policy:
  expr
| sexpr
;

expr:
    expr op=(TK_AND|TK_OR) expr # expr_andor
  | TK_LPAREN expr TK_RPAREN # expr_paren
  | expr op=(TK_EQ|TK_LT|TK_LTEQ|TK_GT|TK_GTEQ) expr # expr_linear
  | attr # expr_attr
  | value # expr_value
;

sexpr:
    TK_LPAREN op=(TK_AND|TK_OR) sexpr+ TK_RPAREN # sexpr_andor
  | op=(TK_EQ|TK_LT|TK_LTEQ|TK_GT|TK_GTEQ) expr expr # sexpr_linear
  | attr # sexpr_attr
  | value # sexpr_value
;

attr:
  Name TK_AT Name
;

value:
  Number
;
