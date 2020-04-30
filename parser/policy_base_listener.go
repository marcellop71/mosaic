// Code generated from Policy.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Policy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePolicyListener is a complete listener for a parse tree produced by PolicyParser.
type BasePolicyListener struct{}

var _ PolicyListener = &BasePolicyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePolicyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePolicyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePolicyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePolicyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPolicy is called when production policy is entered.
func (s *BasePolicyListener) EnterPolicy(ctx *PolicyContext) {}

// ExitPolicy is called when production policy is exited.
func (s *BasePolicyListener) ExitPolicy(ctx *PolicyContext) {}

// EnterExpr_attrrange0 is called when production expr_attrrange0 is entered.
func (s *BasePolicyListener) EnterExpr_attrrange0(ctx *Expr_attrrange0Context) {}

// ExitExpr_attrrange0 is called when production expr_attrrange0 is exited.
func (s *BasePolicyListener) ExitExpr_attrrange0(ctx *Expr_attrrange0Context) {}

// EnterExpr_attrrange1 is called when production expr_attrrange1 is entered.
func (s *BasePolicyListener) EnterExpr_attrrange1(ctx *Expr_attrrange1Context) {}

// ExitExpr_attrrange1 is called when production expr_attrrange1 is exited.
func (s *BasePolicyListener) ExitExpr_attrrange1(ctx *Expr_attrrange1Context) {}

// EnterExpr_paren is called when production expr_paren is entered.
func (s *BasePolicyListener) EnterExpr_paren(ctx *Expr_parenContext) {}

// ExitExpr_paren is called when production expr_paren is exited.
func (s *BasePolicyListener) ExitExpr_paren(ctx *Expr_parenContext) {}

// EnterExpr_attr is called when production expr_attr is entered.
func (s *BasePolicyListener) EnterExpr_attr(ctx *Expr_attrContext) {}

// ExitExpr_attr is called when production expr_attr is exited.
func (s *BasePolicyListener) ExitExpr_attr(ctx *Expr_attrContext) {}

// EnterExpr_andor is called when production expr_andor is entered.
func (s *BasePolicyListener) EnterExpr_andor(ctx *Expr_andorContext) {}

// ExitExpr_andor is called when production expr_andor is exited.
func (s *BasePolicyListener) ExitExpr_andor(ctx *Expr_andorContext) {}

// EnterSexpr_andor is called when production sexpr_andor is entered.
func (s *BasePolicyListener) EnterSexpr_andor(ctx *Sexpr_andorContext) {}

// ExitSexpr_andor is called when production sexpr_andor is exited.
func (s *BasePolicyListener) ExitSexpr_andor(ctx *Sexpr_andorContext) {}

// EnterSexpr_attrrange0 is called when production sexpr_attrrange0 is entered.
func (s *BasePolicyListener) EnterSexpr_attrrange0(ctx *Sexpr_attrrange0Context) {}

// ExitSexpr_attrrange0 is called when production sexpr_attrrange0 is exited.
func (s *BasePolicyListener) ExitSexpr_attrrange0(ctx *Sexpr_attrrange0Context) {}

// EnterSexpr_attrrange1 is called when production sexpr_attrrange1 is entered.
func (s *BasePolicyListener) EnterSexpr_attrrange1(ctx *Sexpr_attrrange1Context) {}

// ExitSexpr_attrrange1 is called when production sexpr_attrrange1 is exited.
func (s *BasePolicyListener) ExitSexpr_attrrange1(ctx *Sexpr_attrrange1Context) {}

// EnterSexpr_attr is called when production sexpr_attr is entered.
func (s *BasePolicyListener) EnterSexpr_attr(ctx *Sexpr_attrContext) {}

// ExitSexpr_attr is called when production sexpr_attr is exited.
func (s *BasePolicyListener) ExitSexpr_attr(ctx *Sexpr_attrContext) {}

// EnterAttr is called when production attr is entered.
func (s *BasePolicyListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BasePolicyListener) ExitAttr(ctx *AttrContext) {}

// EnterValue is called when production value is entered.
func (s *BasePolicyListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePolicyListener) ExitValue(ctx *ValueContext) {}
