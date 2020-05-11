// Code generated from Policy.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Policy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePolicyListener is a complete listener for a parse tree produced by Policy.
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

// EnterExpr_value is called when production expr_value is entered.
func (s *BasePolicyListener) EnterExpr_value(ctx *Expr_valueContext) {}

// ExitExpr_value is called when production expr_value is exited.
func (s *BasePolicyListener) ExitExpr_value(ctx *Expr_valueContext) {}

// EnterExpr_linear is called when production expr_linear is entered.
func (s *BasePolicyListener) EnterExpr_linear(ctx *Expr_linearContext) {}

// ExitExpr_linear is called when production expr_linear is exited.
func (s *BasePolicyListener) ExitExpr_linear(ctx *Expr_linearContext) {}

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

// EnterSexpr_linear is called when production sexpr_linear is entered.
func (s *BasePolicyListener) EnterSexpr_linear(ctx *Sexpr_linearContext) {}

// ExitSexpr_linear is called when production sexpr_linear is exited.
func (s *BasePolicyListener) ExitSexpr_linear(ctx *Sexpr_linearContext) {}

// EnterSexpr_attr is called when production sexpr_attr is entered.
func (s *BasePolicyListener) EnterSexpr_attr(ctx *Sexpr_attrContext) {}

// ExitSexpr_attr is called when production sexpr_attr is exited.
func (s *BasePolicyListener) ExitSexpr_attr(ctx *Sexpr_attrContext) {}

// EnterSexpr_value is called when production sexpr_value is entered.
func (s *BasePolicyListener) EnterSexpr_value(ctx *Sexpr_valueContext) {}

// ExitSexpr_value is called when production sexpr_value is exited.
func (s *BasePolicyListener) ExitSexpr_value(ctx *Sexpr_valueContext) {}

// EnterAttr is called when production attr is entered.
func (s *BasePolicyListener) EnterAttr(ctx *AttrContext) {}

// ExitAttr is called when production attr is exited.
func (s *BasePolicyListener) ExitAttr(ctx *AttrContext) {}

// EnterValue is called when production value is entered.
func (s *BasePolicyListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BasePolicyListener) ExitValue(ctx *ValueContext) {}
