// Code generated from Policy.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Policy

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 73, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 34, 10, 3, 3, 3, 3,
	3, 3, 3, 7, 3, 39, 10, 3, 12, 3, 14, 3, 42, 11, 3, 3, 4, 3, 4, 3, 4, 6,
	4, 47, 10, 4, 13, 4, 14, 4, 48, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 3, 4, 7, 2, 4, 6, 8, 10, 2, 4, 4, 2,
	6, 6, 8, 11, 3, 2, 13, 14, 2, 76, 2, 14, 3, 2, 2, 2, 4, 33, 3, 2, 2, 2,
	6, 64, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 70, 3, 2, 2, 2, 12, 15, 5, 4,
	3, 2, 13, 15, 5, 6, 4, 2, 14, 12, 3, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15, 3,
	3, 2, 2, 2, 16, 17, 8, 3, 1, 2, 17, 18, 7, 18, 2, 2, 18, 19, 5, 4, 3, 2,
	19, 20, 7, 19, 2, 2, 20, 34, 3, 2, 2, 2, 21, 22, 5, 8, 5, 2, 22, 23, 9,
	2, 2, 2, 23, 24, 5, 10, 6, 2, 24, 34, 3, 2, 2, 2, 25, 26, 5, 8, 5, 2, 26,
	27, 7, 16, 2, 2, 27, 28, 5, 10, 6, 2, 28, 29, 7, 15, 2, 2, 29, 30, 5, 10,
	6, 2, 30, 31, 7, 16, 2, 2, 31, 34, 3, 2, 2, 2, 32, 34, 5, 8, 5, 2, 33,
	16, 3, 2, 2, 2, 33, 21, 3, 2, 2, 2, 33, 25, 3, 2, 2, 2, 33, 32, 3, 2, 2,
	2, 34, 40, 3, 2, 2, 2, 35, 36, 12, 7, 2, 2, 36, 37, 9, 3, 2, 2, 37, 39,
	5, 4, 3, 8, 38, 35, 3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2,
	40, 41, 3, 2, 2, 2, 41, 5, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 44, 7, 18,
	2, 2, 44, 46, 9, 3, 2, 2, 45, 47, 5, 6, 4, 2, 46, 45, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2,
	50, 51, 7, 19, 2, 2, 51, 65, 3, 2, 2, 2, 52, 53, 9, 2, 2, 2, 53, 54, 5,
	8, 5, 2, 54, 55, 5, 10, 6, 2, 55, 65, 3, 2, 2, 2, 56, 57, 5, 8, 5, 2, 57,
	58, 7, 16, 2, 2, 58, 59, 5, 10, 6, 2, 59, 60, 7, 15, 2, 2, 60, 61, 5, 10,
	6, 2, 61, 62, 7, 16, 2, 2, 62, 65, 3, 2, 2, 2, 63, 65, 5, 8, 5, 2, 64,
	43, 3, 2, 2, 2, 64, 52, 3, 2, 2, 2, 64, 56, 3, 2, 2, 2, 64, 63, 3, 2, 2,
	2, 65, 7, 3, 2, 2, 2, 66, 67, 7, 3, 2, 2, 67, 68, 7, 12, 2, 2, 68, 69,
	7, 3, 2, 2, 69, 9, 3, 2, 2, 2, 70, 71, 7, 4, 2, 2, 71, 11, 3, 2, 2, 2,
	7, 14, 33, 40, 48, 64,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'@'", "'/\\'",
	"'\\/'", "','", "'['", "']'", "'('", "')'",
}
var symbolicNames = []string{
	"", "Name", "Number", "WS", "TK_EQ", "TK_NEQ", "TK_LT", "TK_LTEQ", "TK_GT",
	"TK_GTEQ", "TK_AT", "TK_AND", "TK_OR", "TK_COMMA", "TK_LBRACKET", "TK_RBRACKET",
	"TK_LPAREN", "TK_RPAREN",
}

var ruleNames = []string{
	"policy", "expr", "sexpr", "attr", "value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PolicyParser struct {
	*antlr.BaseParser
}

func NewPolicyParser(input antlr.TokenStream) *PolicyParser {
	this := new(PolicyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Policy.g4"

	return this
}

// PolicyParser tokens.
const (
	PolicyParserEOF         = antlr.TokenEOF
	PolicyParserName        = 1
	PolicyParserNumber      = 2
	PolicyParserWS          = 3
	PolicyParserTK_EQ       = 4
	PolicyParserTK_NEQ      = 5
	PolicyParserTK_LT       = 6
	PolicyParserTK_LTEQ     = 7
	PolicyParserTK_GT       = 8
	PolicyParserTK_GTEQ     = 9
	PolicyParserTK_AT       = 10
	PolicyParserTK_AND      = 11
	PolicyParserTK_OR       = 12
	PolicyParserTK_COMMA    = 13
	PolicyParserTK_LBRACKET = 14
	PolicyParserTK_RBRACKET = 15
	PolicyParserTK_LPAREN   = 16
	PolicyParserTK_RPAREN   = 17
)

// PolicyParser rules.
const (
	PolicyParserRULE_policy = 0
	PolicyParserRULE_expr   = 1
	PolicyParserRULE_sexpr  = 2
	PolicyParserRULE_attr   = 3
	PolicyParserRULE_value  = 4
)

// IPolicyContext is an interface to support dynamic dispatch.
type IPolicyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolicyContext differentiates from other interfaces.
	IsPolicyContext()
}

type PolicyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolicyContext() *PolicyContext {
	var p = new(PolicyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PolicyParserRULE_policy
	return p
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyParserRULE_policy

	return p
}

func (s *PolicyContext) GetParser() antlr.Parser { return s.parser }

func (s *PolicyContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PolicyContext) Sexpr() ISexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISexprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *PolicyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolicyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolicyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterPolicy(s)
	}
}

func (s *PolicyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitPolicy(s)
	}
}

func (p *PolicyParser) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PolicyParserRULE_policy)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(12)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(10)
			p.expr(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(11)
			p.Sexpr()
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PolicyParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expr_attrrange0Context struct {
	*ExprContext
	op antlr.Token
}

func NewExpr_attrrange0Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_attrrange0Context {
	var p = new(Expr_attrrange0Context)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_attrrange0Context) GetOp() antlr.Token { return s.op }

func (s *Expr_attrrange0Context) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_attrrange0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_attrrange0Context) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Expr_attrrange0Context) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Expr_attrrange0Context) TK_EQ() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_EQ, 0)
}

func (s *Expr_attrrange0Context) TK_LT() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LT, 0)
}

func (s *Expr_attrrange0Context) TK_LTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LTEQ, 0)
}

func (s *Expr_attrrange0Context) TK_GT() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_GT, 0)
}

func (s *Expr_attrrange0Context) TK_GTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_GTEQ, 0)
}

func (s *Expr_attrrange0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_attrrange0(s)
	}
}

func (s *Expr_attrrange0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_attrrange0(s)
	}
}

type Expr_attrrange1Context struct {
	*ExprContext
}

func NewExpr_attrrange1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_attrrange1Context {
	var p = new(Expr_attrrange1Context)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_attrrange1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_attrrange1Context) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Expr_attrrange1Context) AllTK_LBRACKET() []antlr.TerminalNode {
	return s.GetTokens(PolicyParserTK_LBRACKET)
}

func (s *Expr_attrrange1Context) TK_LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LBRACKET, i)
}

func (s *Expr_attrrange1Context) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Expr_attrrange1Context) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Expr_attrrange1Context) TK_COMMA() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_COMMA, 0)
}

func (s *Expr_attrrange1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_attrrange1(s)
	}
}

func (s *Expr_attrrange1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_attrrange1(s)
	}
}

type Expr_parenContext struct {
	*ExprContext
}

func NewExpr_parenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_parenContext {
	var p = new(Expr_parenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_parenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_parenContext) TK_LPAREN() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LPAREN, 0)
}

func (s *Expr_parenContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_parenContext) TK_RPAREN() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_RPAREN, 0)
}

func (s *Expr_parenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_paren(s)
	}
}

func (s *Expr_parenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_paren(s)
	}
}

type Expr_attrContext struct {
	*ExprContext
}

func NewExpr_attrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_attrContext {
	var p = new(Expr_attrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_attrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_attrContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Expr_attrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_attr(s)
	}
}

func (s *Expr_attrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_attr(s)
	}
}

type Expr_andorContext struct {
	*ExprContext
	op antlr.Token
}

func NewExpr_andorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_andorContext {
	var p = new(Expr_andorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_andorContext) GetOp() antlr.Token { return s.op }

func (s *Expr_andorContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_andorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_andorContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Expr_andorContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_andorContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_AND, 0)
}

func (s *Expr_andorContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_OR, 0)
}

func (s *Expr_andorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_andor(s)
	}
}

func (s *Expr_andorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_andor(s)
	}
}

func (p *PolicyParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *PolicyParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PolicyParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpr_parenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(15)
			p.Match(PolicyParserTK_LPAREN)
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		{
			p.SetState(17)
			p.Match(PolicyParserTK_RPAREN)
		}

	case 2:
		localctx = NewExpr_attrrange0Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Attr()
		}
		{
			p.SetState(20)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_attrrange0Context).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PolicyParserTK_EQ)|(1<<PolicyParserTK_LT)|(1<<PolicyParserTK_LTEQ)|(1<<PolicyParserTK_GT)|(1<<PolicyParserTK_GTEQ))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_attrrange0Context).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(21)
			p.Value()
		}

	case 3:
		localctx = NewExpr_attrrange1Context(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Attr()
		}
		{
			p.SetState(24)
			p.Match(PolicyParserTK_LBRACKET)
		}
		{
			p.SetState(25)
			p.Value()
		}
		{
			p.SetState(26)
			p.Match(PolicyParserTK_COMMA)
		}
		{
			p.SetState(27)
			p.Value()
		}
		{
			p.SetState(28)
			p.Match(PolicyParserTK_LBRACKET)
		}

	case 4:
		localctx = NewExpr_attrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Attr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpr_andorContext(p, NewExprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, PolicyParserRULE_expr)
			p.SetState(33)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(34)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Expr_andorContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == PolicyParserTK_AND || _la == PolicyParserTK_OR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Expr_andorContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(35)
				p.expr(6)
			}

		}
		p.SetState(40)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// ISexprContext is an interface to support dynamic dispatch.
type ISexprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSexprContext differentiates from other interfaces.
	IsSexprContext()
}

type SexprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySexprContext() *SexprContext {
	var p = new(SexprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PolicyParserRULE_sexpr
	return p
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyParserRULE_sexpr

	return p
}

func (s *SexprContext) GetParser() antlr.Parser { return s.parser }

func (s *SexprContext) CopyFrom(ctx *SexprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SexprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SexprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Sexpr_andorContext struct {
	*SexprContext
	op antlr.Token
}

func NewSexpr_andorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Sexpr_andorContext {
	var p = new(Sexpr_andorContext)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *Sexpr_andorContext) GetOp() antlr.Token { return s.op }

func (s *Sexpr_andorContext) SetOp(v antlr.Token) { s.op = v }

func (s *Sexpr_andorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sexpr_andorContext) TK_LPAREN() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LPAREN, 0)
}

func (s *Sexpr_andorContext) TK_RPAREN() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_RPAREN, 0)
}

func (s *Sexpr_andorContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_AND, 0)
}

func (s *Sexpr_andorContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_OR, 0)
}

func (s *Sexpr_andorContext) AllSexpr() []ISexprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISexprContext)(nil)).Elem())
	var tst = make([]ISexprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISexprContext)
		}
	}

	return tst
}

func (s *Sexpr_andorContext) Sexpr(i int) ISexprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISexprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISexprContext)
}

func (s *Sexpr_andorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterSexpr_andor(s)
	}
}

func (s *Sexpr_andorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitSexpr_andor(s)
	}
}

type Sexpr_attrrange0Context struct {
	*SexprContext
	op antlr.Token
}

func NewSexpr_attrrange0Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Sexpr_attrrange0Context {
	var p = new(Sexpr_attrrange0Context)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *Sexpr_attrrange0Context) GetOp() antlr.Token { return s.op }

func (s *Sexpr_attrrange0Context) SetOp(v antlr.Token) { s.op = v }

func (s *Sexpr_attrrange0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sexpr_attrrange0Context) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Sexpr_attrrange0Context) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Sexpr_attrrange0Context) TK_EQ() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_EQ, 0)
}

func (s *Sexpr_attrrange0Context) TK_LT() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LT, 0)
}

func (s *Sexpr_attrrange0Context) TK_LTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LTEQ, 0)
}

func (s *Sexpr_attrrange0Context) TK_GT() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_GT, 0)
}

func (s *Sexpr_attrrange0Context) TK_GTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_GTEQ, 0)
}

func (s *Sexpr_attrrange0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterSexpr_attrrange0(s)
	}
}

func (s *Sexpr_attrrange0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitSexpr_attrrange0(s)
	}
}

type Sexpr_attrContext struct {
	*SexprContext
}

func NewSexpr_attrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Sexpr_attrContext {
	var p = new(Sexpr_attrContext)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *Sexpr_attrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sexpr_attrContext) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Sexpr_attrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterSexpr_attr(s)
	}
}

func (s *Sexpr_attrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitSexpr_attr(s)
	}
}

type Sexpr_attrrange1Context struct {
	*SexprContext
}

func NewSexpr_attrrange1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Sexpr_attrrange1Context {
	var p = new(Sexpr_attrrange1Context)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *Sexpr_attrrange1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sexpr_attrrange1Context) Attr() IAttrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttrContext)
}

func (s *Sexpr_attrrange1Context) AllTK_LBRACKET() []antlr.TerminalNode {
	return s.GetTokens(PolicyParserTK_LBRACKET)
}

func (s *Sexpr_attrrange1Context) TK_LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_LBRACKET, i)
}

func (s *Sexpr_attrrange1Context) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *Sexpr_attrrange1Context) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Sexpr_attrrange1Context) TK_COMMA() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_COMMA, 0)
}

func (s *Sexpr_attrrange1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterSexpr_attrrange1(s)
	}
}

func (s *Sexpr_attrrange1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitSexpr_attrrange1(s)
	}
}

func (p *PolicyParser) Sexpr() (localctx ISexprContext) {
	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PolicyParserRULE_sexpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSexpr_andorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(PolicyParserTK_LPAREN)
		}
		{
			p.SetState(42)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Sexpr_andorContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PolicyParserTK_AND || _la == PolicyParserTK_OR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Sexpr_andorContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PolicyParserName)|(1<<PolicyParserTK_EQ)|(1<<PolicyParserTK_LT)|(1<<PolicyParserTK_LTEQ)|(1<<PolicyParserTK_GT)|(1<<PolicyParserTK_GTEQ)|(1<<PolicyParserTK_LPAREN))) != 0) {
			{
				p.SetState(43)
				p.Sexpr()
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(48)
			p.Match(PolicyParserTK_RPAREN)
		}

	case 2:
		localctx = NewSexpr_attrrange0Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(50)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Sexpr_attrrange0Context).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PolicyParserTK_EQ)|(1<<PolicyParserTK_LT)|(1<<PolicyParserTK_LTEQ)|(1<<PolicyParserTK_GT)|(1<<PolicyParserTK_GTEQ))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Sexpr_attrrange0Context).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(51)
			p.Attr()
		}
		{
			p.SetState(52)
			p.Value()
		}

	case 3:
		localctx = NewSexpr_attrrange1Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Attr()
		}
		{
			p.SetState(55)
			p.Match(PolicyParserTK_LBRACKET)
		}
		{
			p.SetState(56)
			p.Value()
		}
		{
			p.SetState(57)
			p.Match(PolicyParserTK_COMMA)
		}
		{
			p.SetState(58)
			p.Value()
		}
		{
			p.SetState(59)
			p.Match(PolicyParserTK_LBRACKET)
		}

	case 4:
		localctx = NewSexpr_attrContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.Attr()
		}

	}

	return localctx
}

// IAttrContext is an interface to support dynamic dispatch.
type IAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttrContext differentiates from other interfaces.
	IsAttrContext()
}

type AttrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrContext() *AttrContext {
	var p = new(AttrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PolicyParserRULE_attr
	return p
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyParserRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(PolicyParserName)
}

func (s *AttrContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(PolicyParserName, i)
}

func (s *AttrContext) TK_AT() antlr.TerminalNode {
	return s.GetToken(PolicyParserTK_AT, 0)
}

func (s *AttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterAttr(s)
	}
}

func (s *AttrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitAttr(s)
	}
}

func (p *PolicyParser) Attr() (localctx IAttrContext) {
	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PolicyParserRULE_attr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(PolicyParserName)
	}
	{
		p.SetState(65)
		p.Match(PolicyParserTK_AT)
	}
	{
		p.SetState(66)
		p.Match(PolicyParserName)
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PolicyParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Number() antlr.TerminalNode {
	return s.GetToken(PolicyParserNumber, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *PolicyParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PolicyParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(PolicyParserNumber)
	}

	return localctx
}

func (p *PolicyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PolicyParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
