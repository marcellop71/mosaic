// Code generated from Policy.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Policy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PolicyListener is a complete listener for a parse tree produced by PolicyParser.
type PolicyListener interface {
	antlr.ParseTreeListener

	// EnterPolicy is called when entering the policy production.
	EnterPolicy(c *PolicyContext)

	// EnterExpr_attrrange0 is called when entering the expr_attrrange0 production.
	EnterExpr_attrrange0(c *Expr_attrrange0Context)

	// EnterExpr_attrrange1 is called when entering the expr_attrrange1 production.
	EnterExpr_attrrange1(c *Expr_attrrange1Context)

	// EnterExpr_paren is called when entering the expr_paren production.
	EnterExpr_paren(c *Expr_parenContext)

	// EnterExpr_attr is called when entering the expr_attr production.
	EnterExpr_attr(c *Expr_attrContext)

	// EnterExpr_andor is called when entering the expr_andor production.
	EnterExpr_andor(c *Expr_andorContext)

	// EnterSexpr_andor is called when entering the sexpr_andor production.
	EnterSexpr_andor(c *Sexpr_andorContext)

	// EnterSexpr_attrrange0 is called when entering the sexpr_attrrange0 production.
	EnterSexpr_attrrange0(c *Sexpr_attrrange0Context)

	// EnterSexpr_attrrange1 is called when entering the sexpr_attrrange1 production.
	EnterSexpr_attrrange1(c *Sexpr_attrrange1Context)

	// EnterSexpr_attr is called when entering the sexpr_attr production.
	EnterSexpr_attr(c *Sexpr_attrContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// ExitPolicy is called when exiting the policy production.
	ExitPolicy(c *PolicyContext)

	// ExitExpr_attrrange0 is called when exiting the expr_attrrange0 production.
	ExitExpr_attrrange0(c *Expr_attrrange0Context)

	// ExitExpr_attrrange1 is called when exiting the expr_attrrange1 production.
	ExitExpr_attrrange1(c *Expr_attrrange1Context)

	// ExitExpr_paren is called when exiting the expr_paren production.
	ExitExpr_paren(c *Expr_parenContext)

	// ExitExpr_attr is called when exiting the expr_attr production.
	ExitExpr_attr(c *Expr_attrContext)

	// ExitExpr_andor is called when exiting the expr_andor production.
	ExitExpr_andor(c *Expr_andorContext)

	// ExitSexpr_andor is called when exiting the sexpr_andor production.
	ExitSexpr_andor(c *Sexpr_andorContext)

	// ExitSexpr_attrrange0 is called when exiting the sexpr_attrrange0 production.
	ExitSexpr_attrrange0(c *Sexpr_attrrange0Context)

	// ExitSexpr_attrrange1 is called when exiting the sexpr_attrrange1 production.
	ExitSexpr_attrrange1(c *Sexpr_attrrange1Context)

	// ExitSexpr_attr is called when exiting the sexpr_attr production.
	ExitSexpr_attr(c *Sexpr_attrContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)
}
