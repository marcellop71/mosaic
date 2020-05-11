// Code generated from Policy.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Policy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PolicyListener is a complete listener for a parse tree produced by Policy.
type PolicyListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterExpr_value is called when entering the expr_value production.
	EnterExpr_value(c *Expr_valueContext)

	// EnterExpr_linear is called when entering the expr_linear production.
	EnterExpr_linear(c *Expr_linearContext)

	// EnterExpr_paren is called when entering the expr_paren production.
	EnterExpr_paren(c *Expr_parenContext)

	// EnterExpr_attr is called when entering the expr_attr production.
	EnterExpr_attr(c *Expr_attrContext)

	// EnterExpr_andor is called when entering the expr_andor production.
	EnterExpr_andor(c *Expr_andorContext)

	// EnterSexpr_andor is called when entering the sexpr_andor production.
	EnterSexpr_andor(c *Sexpr_andorContext)

	// EnterSexpr_linear is called when entering the sexpr_linear production.
	EnterSexpr_linear(c *Sexpr_linearContext)

	// EnterSexpr_attr is called when entering the sexpr_attr production.
	EnterSexpr_attr(c *Sexpr_attrContext)

	// EnterSexpr_value is called when entering the sexpr_value production.
	EnterSexpr_value(c *Sexpr_valueContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitExpr_value is called when exiting the expr_value production.
	ExitExpr_value(c *Expr_valueContext)

	// ExitExpr_linear is called when exiting the expr_linear production.
	ExitExpr_linear(c *Expr_linearContext)

	// ExitExpr_paren is called when exiting the expr_paren production.
	ExitExpr_paren(c *Expr_parenContext)

	// ExitExpr_attr is called when exiting the expr_attr production.
	ExitExpr_attr(c *Expr_attrContext)

	// ExitExpr_andor is called when exiting the expr_andor production.
	ExitExpr_andor(c *Expr_andorContext)

	// ExitSexpr_andor is called when exiting the sexpr_andor production.
	ExitSexpr_andor(c *Sexpr_andorContext)

	// ExitSexpr_linear is called when exiting the sexpr_linear production.
	ExitSexpr_linear(c *Sexpr_linearContext)

	// ExitSexpr_attr is called when exiting the sexpr_attr production.
	ExitSexpr_attr(c *Sexpr_attrContext)

	// ExitSexpr_value is called when exiting the sexpr_value production.
	ExitSexpr_value(c *Sexpr_valueContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
