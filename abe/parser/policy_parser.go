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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 60, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 5,
	2, 15, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 24, 10, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 32, 10, 3, 12, 3, 14, 3, 35,
	11, 3, 3, 4, 3, 4, 3, 4, 6, 4, 40, 10, 4, 13, 4, 14, 4, 41, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 52, 10, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 2, 3, 4, 7, 2, 4, 6, 8, 10, 2, 4, 3, 2, 12, 13,
	4, 2, 5, 5, 7, 10, 2, 63, 2, 14, 3, 2, 2, 2, 4, 23, 3, 2, 2, 2, 6, 51,
	3, 2, 2, 2, 8, 53, 3, 2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 15, 5, 4, 3, 2,
	13, 15, 5, 6, 4, 2, 14, 12, 3, 2, 2, 2, 14, 13, 3, 2, 2, 2, 15, 3, 3, 2,
	2, 2, 16, 17, 8, 3, 1, 2, 17, 18, 7, 17, 2, 2, 18, 19, 5, 4, 3, 2, 19,
	20, 7, 18, 2, 2, 20, 24, 3, 2, 2, 2, 21, 24, 5, 8, 5, 2, 22, 24, 5, 10,
	6, 2, 23, 16, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 22, 3, 2, 2, 2, 24, 33,
	3, 2, 2, 2, 25, 26, 12, 7, 2, 2, 26, 27, 9, 2, 2, 2, 27, 32, 5, 4, 3, 8,
	28, 29, 12, 5, 2, 2, 29, 30, 9, 3, 2, 2, 30, 32, 5, 4, 3, 6, 31, 25, 3,
	2, 2, 2, 31, 28, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33,
	34, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 36, 37, 7, 17, 2,
	2, 37, 39, 9, 2, 2, 2, 38, 40, 5, 6, 4, 2, 39, 38, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 44, 7, 18, 2, 2, 44, 52, 3, 2, 2, 2, 45, 46, 9, 3, 2, 2, 46, 47, 5,
	4, 3, 2, 47, 48, 5, 4, 3, 2, 48, 52, 3, 2, 2, 2, 49, 52, 5, 8, 5, 2, 50,
	52, 5, 10, 6, 2, 51, 36, 3, 2, 2, 2, 51, 45, 3, 2, 2, 2, 51, 49, 3, 2,
	2, 2, 51, 50, 3, 2, 2, 2, 52, 7, 3, 2, 2, 2, 53, 54, 7, 3, 2, 2, 54, 55,
	7, 11, 2, 2, 55, 56, 7, 3, 2, 2, 56, 9, 3, 2, 2, 2, 57, 58, 7, 4, 2, 2,
	58, 11, 3, 2, 2, 2, 8, 14, 23, 31, 33, 41, 51,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'@'", "'/\\'",
	"'\\/'", "','", "'['", "']'", "'('", "')'",
}
var symbolicNames = []string{
	"", "Name", "Number", "TK_EQ", "TK_NEQ", "TK_LT", "TK_LTEQ", "TK_GT", "TK_GTEQ",
	"TK_AT", "TK_AND", "TK_OR", "TK_COMMA", "TK_LBRACKET", "TK_RBRACKET", "TK_LPAREN",
	"TK_RPAREN", "TK_LFCR", "TK_WS",
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

type Policy struct {
	*antlr.BaseParser
}

func NewPolicy(input antlr.TokenStream) *Policy {
	this := new(Policy)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Policy.g4"

	return this
}

// Policy tokens.
const (
	PolicyEOF         = antlr.TokenEOF
	PolicyName        = 1
	PolicyNumber      = 2
	PolicyTK_EQ       = 3
	PolicyTK_NEQ      = 4
	PolicyTK_LT       = 5
	PolicyTK_LTEQ     = 6
	PolicyTK_GT       = 7
	PolicyTK_GTEQ     = 8
	PolicyTK_AT       = 9
	PolicyTK_AND      = 10
	PolicyTK_OR       = 11
	PolicyTK_COMMA    = 12
	PolicyTK_LBRACKET = 13
	PolicyTK_RBRACKET = 14
	PolicyTK_LPAREN   = 15
	PolicyTK_RPAREN   = 16
	PolicyTK_LFCR     = 17
	PolicyTK_WS       = 18
)

// Policy rules.
const (
	PolicyRULE_policy = 0
	PolicyRULE_expr   = 1
	PolicyRULE_sexpr  = 2
	PolicyRULE_attr   = 3
	PolicyRULE_value  = 4
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
	p.RuleIndex = PolicyRULE_policy
	return p
}

func (*PolicyContext) IsPolicyContext() {}

func NewPolicyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolicyContext {
	var p = new(PolicyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyRULE_policy

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

func (p *Policy) Policy() (localctx IPolicyContext) {
	localctx = NewPolicyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PolicyRULE_policy)

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
	p.RuleIndex = PolicyRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyRULE_expr

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

type Expr_valueContext struct {
	*ExprContext
}

func NewExpr_valueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_valueContext {
	var p = new(Expr_valueContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_valueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Expr_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_value(s)
	}
}

func (s *Expr_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_value(s)
	}
}

type Expr_linearContext struct {
	*ExprContext
	op antlr.Token
}

func NewExpr_linearContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_linearContext {
	var p = new(Expr_linearContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Expr_linearContext) GetOp() antlr.Token { return s.op }

func (s *Expr_linearContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_linearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_linearContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Expr_linearContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_linearContext) TK_EQ() antlr.TerminalNode {
	return s.GetToken(PolicyTK_EQ, 0)
}

func (s *Expr_linearContext) TK_LT() antlr.TerminalNode {
	return s.GetToken(PolicyTK_LT, 0)
}

func (s *Expr_linearContext) TK_LTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyTK_LTEQ, 0)
}

func (s *Expr_linearContext) TK_GT() antlr.TerminalNode {
	return s.GetToken(PolicyTK_GT, 0)
}

func (s *Expr_linearContext) TK_GTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyTK_GTEQ, 0)
}

func (s *Expr_linearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterExpr_linear(s)
	}
}

func (s *Expr_linearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitExpr_linear(s)
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
	return s.GetToken(PolicyTK_LPAREN, 0)
}

func (s *Expr_parenContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Expr_parenContext) TK_RPAREN() antlr.TerminalNode {
	return s.GetToken(PolicyTK_RPAREN, 0)
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
	return s.GetToken(PolicyTK_AND, 0)
}

func (s *Expr_andorContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(PolicyTK_OR, 0)
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

func (p *Policy) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *Policy) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, PolicyRULE_expr, _p)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PolicyTK_LPAREN:
		localctx = NewExpr_parenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(15)
			p.Match(PolicyTK_LPAREN)
		}
		{
			p.SetState(16)
			p.expr(0)
		}
		{
			p.SetState(17)
			p.Match(PolicyTK_RPAREN)
		}

	case PolicyName:
		localctx = NewExpr_attrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(19)
			p.Attr()
		}

	case PolicyNumber:
		localctx = NewExpr_valueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(20)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_andorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PolicyRULE_expr)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(24)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_andorContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == PolicyTK_AND || _la == PolicyTK_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_andorContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(25)
					p.expr(6)
				}

			case 2:
				localctx = NewExpr_linearContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, PolicyRULE_expr)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(27)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_linearContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PolicyTK_EQ)|(1<<PolicyTK_LT)|(1<<PolicyTK_LTEQ)|(1<<PolicyTK_GT)|(1<<PolicyTK_GTEQ))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_linearContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expr(4)
				}

			}

		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
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
	p.RuleIndex = PolicyRULE_sexpr
	return p
}

func (*SexprContext) IsSexprContext() {}

func NewSexprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SexprContext {
	var p = new(SexprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyRULE_sexpr

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

type Sexpr_valueContext struct {
	*SexprContext
}

func NewSexpr_valueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Sexpr_valueContext {
	var p = new(Sexpr_valueContext)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *Sexpr_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sexpr_valueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Sexpr_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterSexpr_value(s)
	}
}

func (s *Sexpr_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitSexpr_value(s)
	}
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
	return s.GetToken(PolicyTK_LPAREN, 0)
}

func (s *Sexpr_andorContext) TK_RPAREN() antlr.TerminalNode {
	return s.GetToken(PolicyTK_RPAREN, 0)
}

func (s *Sexpr_andorContext) TK_AND() antlr.TerminalNode {
	return s.GetToken(PolicyTK_AND, 0)
}

func (s *Sexpr_andorContext) TK_OR() antlr.TerminalNode {
	return s.GetToken(PolicyTK_OR, 0)
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

type Sexpr_linearContext struct {
	*SexprContext
	op antlr.Token
}

func NewSexpr_linearContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Sexpr_linearContext {
	var p = new(Sexpr_linearContext)

	p.SexprContext = NewEmptySexprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SexprContext))

	return p
}

func (s *Sexpr_linearContext) GetOp() antlr.Token { return s.op }

func (s *Sexpr_linearContext) SetOp(v antlr.Token) { s.op = v }

func (s *Sexpr_linearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sexpr_linearContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Sexpr_linearContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Sexpr_linearContext) TK_EQ() antlr.TerminalNode {
	return s.GetToken(PolicyTK_EQ, 0)
}

func (s *Sexpr_linearContext) TK_LT() antlr.TerminalNode {
	return s.GetToken(PolicyTK_LT, 0)
}

func (s *Sexpr_linearContext) TK_LTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyTK_LTEQ, 0)
}

func (s *Sexpr_linearContext) TK_GT() antlr.TerminalNode {
	return s.GetToken(PolicyTK_GT, 0)
}

func (s *Sexpr_linearContext) TK_GTEQ() antlr.TerminalNode {
	return s.GetToken(PolicyTK_GTEQ, 0)
}

func (s *Sexpr_linearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.EnterSexpr_linear(s)
	}
}

func (s *Sexpr_linearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PolicyListener); ok {
		listenerT.ExitSexpr_linear(s)
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

func (p *Policy) Sexpr() (localctx ISexprContext) {
	localctx = NewSexprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PolicyRULE_sexpr)
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

	p.SetState(49)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PolicyTK_LPAREN:
		localctx = NewSexpr_andorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Match(PolicyTK_LPAREN)
		}
		{
			p.SetState(35)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Sexpr_andorContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == PolicyTK_AND || _la == PolicyTK_OR) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Sexpr_andorContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PolicyName)|(1<<PolicyNumber)|(1<<PolicyTK_EQ)|(1<<PolicyTK_LT)|(1<<PolicyTK_LTEQ)|(1<<PolicyTK_GT)|(1<<PolicyTK_GTEQ)|(1<<PolicyTK_LPAREN))) != 0) {
			{
				p.SetState(36)
				p.Sexpr()
			}

			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(41)
			p.Match(PolicyTK_RPAREN)
		}

	case PolicyTK_EQ, PolicyTK_LT, PolicyTK_LTEQ, PolicyTK_GT, PolicyTK_GTEQ:
		localctx = NewSexpr_linearContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Sexpr_linearContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PolicyTK_EQ)|(1<<PolicyTK_LT)|(1<<PolicyTK_LTEQ)|(1<<PolicyTK_GT)|(1<<PolicyTK_GTEQ))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Sexpr_linearContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(44)
			p.expr(0)
		}
		{
			p.SetState(45)
			p.expr(0)
		}

	case PolicyName:
		localctx = NewSexpr_attrContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)
			p.Attr()
		}

	case PolicyNumber:
		localctx = NewSexpr_valueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)
			p.Value()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = PolicyRULE_attr
	return p
}

func (*AttrContext) IsAttrContext() {}

func NewAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrContext {
	var p = new(AttrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyRULE_attr

	return p
}

func (s *AttrContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(PolicyName)
}

func (s *AttrContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(PolicyName, i)
}

func (s *AttrContext) TK_AT() antlr.TerminalNode {
	return s.GetToken(PolicyTK_AT, 0)
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

func (p *Policy) Attr() (localctx IAttrContext) {
	localctx = NewAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PolicyRULE_attr)

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
		p.SetState(51)
		p.Match(PolicyName)
	}
	{
		p.SetState(52)
		p.Match(PolicyTK_AT)
	}
	{
		p.SetState(53)
		p.Match(PolicyName)
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
	p.RuleIndex = PolicyRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PolicyRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Number() antlr.TerminalNode {
	return s.GetToken(PolicyNumber, 0)
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

func (p *Policy) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PolicyRULE_value)

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
		p.SetState(55)
		p.Match(PolicyNumber)
	}

	return localctx
}

func (p *Policy) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *Policy) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
