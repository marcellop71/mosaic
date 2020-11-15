package abe

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/marcellop71/mosaic/abe/log"
	"github.com/marcellop71/mosaic/abe/parser"
)

// binary tree with string labels and
// vectors for the LSSS representation
type btree struct {
	label string
	child [2]*btree
	v     []int
}

type policyListener struct {
	*parser.BasePolicyListener

	method int
	s      []*btree          // stack of trees
	vars   []string          // list of varieables
	vars_t map[string]string // map variables to their type
}

// add leaf (if is a value then it is not appended to the vars list)
func (ln *policyListener) addLeaf(label string, isvalue bool) {
	t := new(btree)
	t.label = label
	ln.s = append(ln.s, t) // push tree to the stack
	if isvalue == false {
		ln.vars = append(ln.vars, label)
		ln.vars_t[label] = "Bool" // init the guess on type with bool
	}
}

// add internal node
func (ln *policyListener) addNode(label string, childcount int) {
	for i := 1; i < childcount; i++ {
		t := new(btree)
		t.label = label
		t.child[0] = ln.s[len(ln.s)-2] // left subtree
		t.child[1] = ln.s[len(ln.s)-1] // right subtree
		ln.s = ln.s[:len(ln.s)-2]      // pop 2 trees from the stack
		ln.s = append(ln.s, t)         // push tree to the stack
	}
}

// rewrite eq as a bit expression
func (ln *policyListener) eq(attr string, attrs []string, v string) {
	nbit := len(v)

	for i := nbit - 1; i >= 0; i-- {
		ln.addLeaf(attrs[i], false)
	}

	for i := nbit - 1; i > 0; i-- {
		ln.addNode("and", 2)
	}
}

// rewrite gt as a bit expression
func (ln *policyListener) gt(attr string, attrs []string, v string) {
	nbit := len(v)

	trailing_ones := 0
	for i := nbit - 1; i >= 0; i-- {
		if v[i] == '1' {
			trailing_ones++
		} else {
			break
		}
	}

	for i := nbit - 1; i >= trailing_ones; i-- {
		ln.addLeaf(formatAttr(attr, i, 1), false)
	}

	for i := nbit - 2 - trailing_ones; i >= 0; i-- {
		switch v[i] {
		case '0':
			ln.addNode("or", 2)
		case '1':
			ln.addNode("and", 2)
		}
	}
}

// rewrite gt as a bit expression
func (ln *policyListener) gteq(attr string, attrs []string, v string) {
	nbit := len(v)

	trailing_zeros := 0
	for i := nbit - 1; i >= 0; i-- {
		if v[i] == '0' {
			trailing_zeros++
		} else {
			break
		}
	}

	for i := nbit - 1; i >= trailing_zeros; i-- {
		ln.addLeaf(formatAttr(attr, i, 1), false)
	}

	for i := nbit - 2 - trailing_zeros; i >= 0; i-- {
		switch v[i] {
		case '0':
			ln.addNode("or", 2)
		case '1':
			ln.addNode("and", 2)
		}
	}
}

// rewrite lt as a bit expression
func (ln *policyListener) lt(attr string, attrs []string, v string) {
	nbit := len(v)

	trailing_zeros := 0
	for i := nbit - 1; i >= 0; i-- {
		if v[i] == '0' {
			trailing_zeros++
		} else {
			break
		}
	}

	for i := nbit - 1; i >= trailing_zeros; i-- {
		ln.addLeaf(formatAttr(attr, i, 0), false)
	}

	for i := nbit - 2 - trailing_zeros; i >= 0; i-- {
		switch v[i] {
		case '0':
			ln.addNode("and", 2)
		case '1':
			ln.addNode("or", 2)
		}
	}
}

// rewrite lt as a bit expression
func (ln *policyListener) lteq(attr string, attrs []string, v string) {
	nbit := len(v)

	trailing_ones := 0
	for i := nbit - 1; i >= 0; i-- {
		if v[i] == '1' {
			trailing_ones++
		} else {
			break
		}
	}

	for i := nbit - 1; i >= trailing_ones; i-- {
		ln.addLeaf(formatAttr(attr, i, 0), false)
	}

	for i := nbit - 2 - trailing_ones; i >= 0; i-- {
		switch v[i] {
		case '0':
			ln.addNode("and", 2)
		case '1':
			ln.addNode("or", 2)
		}
	}
}

func (ln *policyListener) ExitExpr_andor(ctx *parser.Expr_andorContext) {
	switch ctx.GetOp().GetTokenType() {
	case parser.PolicyTK_AND:
		ln.addNode("and", 2)
	case parser.PolicyTK_OR:
		ln.addNode("or", 2)
	}
}

func (ln *policyListener) ExitSexpr_andor(ctx *parser.Sexpr_andorContext) {
	cc := ctx.GetRuleContext().GetChildCount() - 3
	switch ctx.GetOp().GetTokenType() {
	case parser.PolicyTK_AND:
		ln.addNode("and", cc)
	case parser.PolicyTK_OR:
		ln.addNode("or", cc)
	}
}

func (ln *policyListener) ExitExpr_linear(ctx *parser.Expr_linearContext) {
	switch ln.method {
	case 0:
		attr := ln.s[len(ln.s)-2].label
		value := ln.s[len(ln.s)-1].label
		value_, _ := strconv.Atoi(value)
		nbit := 32
		attrs := getBagOfBits(attr, value_, nbit)
		v := getBitString(value_, nbit)
		ln.s = ln.s[0 : len(ln.s)-2]
		ln.vars = ln.vars[0 : len(ln.vars)-1]
		delete(ln.vars_t, attr)

		switch ctx.GetOp().GetTokenType() {
		case parser.PolicyTK_EQ:
			ln.eq(attr, attrs, v)
		case parser.PolicyTK_GT:
			ln.gt(attr, attrs, v)
		case parser.PolicyTK_GTEQ:
			ln.gteq(attr, attrs, v)
		case parser.PolicyTK_LT:
			ln.lt(attr, attrs, v)
		case parser.PolicyTK_LTEQ:
			ln.lteq(attr, attrs, v)
		}
	case 1:
		v := ln.vars[len(ln.vars)-1]
		ln.vars_t[v] = "Int"
		switch ctx.GetOp().GetTokenType() {
		case parser.PolicyTK_EQ:
			ln.addNode("=", 2)
		case parser.PolicyTK_GT:
			ln.addNode(">", 2)
		case parser.PolicyTK_GTEQ:
			ln.addNode(">=", 2)
		case parser.PolicyTK_LT:
			ln.addNode("<", 2)
		case parser.PolicyTK_LTEQ:
			ln.addNode("<=", 2)
		}
	}
}

func (ln *policyListener) ExitSexpr_linear(ctx *parser.Sexpr_linearContext) {
	switch ln.method {
	case 0:
		attr := ln.s[len(ln.s)-2].label
		value := ln.s[len(ln.s)-1].label
		value_, _ := strconv.Atoi(value)
		nbit := 32
		attrs := getBagOfBits(attr, value_, nbit)
		v := getBitString(value_, nbit)
		ln.vars = ln.vars[0 : len(ln.vars)-1]
		delete(ln.vars_t, attr)

		switch ctx.GetOp().GetTokenType() {
		case parser.PolicyTK_EQ:
			ln.eq(attr, attrs, v)
		case parser.PolicyTK_GT:
			ln.gt(attr, attrs, v)
		case parser.PolicyTK_GTEQ:
			ln.gteq(attr, attrs, v)
		case parser.PolicyTK_LT:
			ln.lt(attr, attrs, v)
		case parser.PolicyTK_LTEQ:
			ln.lteq(attr, attrs, v)
		}
	case 1:
		v := ln.vars[len(ln.vars)-1]
		ln.vars_t[v] = "Int"
		switch ctx.GetOp().GetTokenType() {
		case parser.PolicyTK_EQ:
			ln.addNode("=", 2)
		case parser.PolicyTK_GT:
			ln.addNode(">", 2)
		case parser.PolicyTK_GTEQ:
			ln.addNode(">=", 2)
		case parser.PolicyTK_LT:
			ln.addNode("<", 2)
		case parser.PolicyTK_LTEQ:
			ln.addNode("<=", 2)
		}
	}
}

func (ln *policyListener) ExitExpr_attr(ctx *parser.Expr_attrContext) {
	ln.addLeaf(ctx.GetText(), false)
}

func (ln *policyListener) ExitExpr_value(ctx *parser.Expr_valueContext) {
	ln.addLeaf(ctx.GetText(), true)
}

func (ln *policyListener) ExitSexpr_attr(ctx *parser.Sexpr_attrContext) {
	ln.addLeaf(ctx.GetText(), false)
}

func (ln *policyListener) ExitSexpr_value(ctx *parser.Sexpr_valueContext) {
	ln.addLeaf(ctx.GetText(), true)
}

func dumpSExpr(t *btree, s string) string {
	internal := (t.child[0] != nil) && (t.child[1] != nil)
	if internal == true {
		s += "(" + t.label + " "
		s = dumpSExpr(t.child[0], s)
		s += " "
		s = dumpSExpr(t.child[1], s)
		s += ")"
	} else {
		s += fmt.Sprint(t.label)
	}
	return s
}

// parse a policy string
func parsePolicy(policy string, method int) *policyListener {
	lexer := parser.NewPolicyLexer(antlr.NewInputStream(policy))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPolicy(stream)
	ln := new(policyListener)
	ln.method = method
	ln.vars_t = make(map[string]string)
	antlr.ParseTreeWalkerDefault.Walk(ln, p.Policy())
	return ln
}

// format for the bit-attribute-strings of a given attribute-string
func formatAttr(attr string, bitidx int, bit int) string {
	s := strings.SplitN(attr, "@", 2)
	return fmt.Sprintf("%s-%d-%d@%s", s[0], bitidx, bit, s[1])
}

// bit string representation of an integer
func getBitString(x int, nbit int) string {
	bits := strconv.FormatInt(int64(x), 2)
	s := ""
	for bitidx := len(bits); bitidx < nbit; bitidx++ {
		s += "0"
	}
	s += bits
	return s
}

// bag-of-bits representation of an attribute
func getBagOfBits(attr string, value int, nbit int) []string {
	var attrs []string
	bits := getBitString(value, nbit)
	for bitidx, bit_s := range bits {
		bit, _ := strconv.Atoi(string(bit_s))
		attrs = append(attrs, formatAttr(attr, len(bits)-1-bitidx, bit))
	}
	return attrs
}

// expands an attribute in its bag-of-bits representation
func getBagOfBitsAttrs(attr string, nbit int) []string {
	var attrs []string
	s0 := strings.Split(attr, "@")
	s1 := strings.Split(s0[0], "=")
	attr_tmp := fmt.Sprintf("%s@%s", s1[0], s0[1])
	if len(s1) == 1 {
		attrs = []string{attr_tmp}
	} else {
		value, _ := strconv.Atoi(s1[1])
		attrs = getBagOfBits(attr_tmp, value, nbit)
	}
	return attrs
}

// builds access policy matrix from policy string
func buildAccessPolicy(policy string) *AccessPolicy {
	method := 0
	ln := parsePolicy(policy, method)
	ap := encodeAccessPolicy(ln.s[0])
	row := make(map[string][]int)
	for i, v := range ln.vars {
		row[v] = append(row[v], i)
	}
	log.Debug("policy vars: %s", ln.vars)
	log.Debug("map vars to rows: %v", row)
	return &AccessPolicy{M: ap, Vars: ln.vars, Row: row}
}

// extracts authorities from a policy string
func AuthPubsOfPolicy(policy string) *AuthPubs {
	method := 0
	ln := parsePolicy(policy, method)
	authpub := make(map[string]*AuthPub)
	for _, v := range ln.vars {
		v_ := strings.SplitN(v, "@", 2)[1]
		authpub[v_] = nil
	}
	log.Debug("policy authorities: %v", authpub)
	return &AuthPubs{AuthPub: authpub}
}

// Json API for AuthPubsOfPolicy
func AuthPubsOfPolicyJson(policy string) string {
	authpubs := AuthPubsOfPolicy(policy)
	return Encode(JsonObjToStr(authpubs.ToJsonObj()))
}

// select user attrs relevant for a given policy string
func (userattrs *UserAttrs) SelectUserAttrs(user string, policy string) *UserAttrs {
	userattrs.User = user

	if len(userattrs.Coeff) > 0 {
		ap := buildAccessPolicy(policy)
		log.Debug(">>> policy: %s", policy)
		log.Debug("access policy (unfiltered, %d x %d matrix) %d",
			len(ap.M[0]), len(ap.M), ap.M)

		// policy matrix intersection with user attributes
		var ap_ [][]int
		row_idx := make(map[string][]int)
		i_ := 0
		for attr, _ := range userattrs.Coeff {
			if rows, exist := ap.Row[attr]; exist {
				for _, i := range rows {
					ap_row := make([]int, len(ap.M[i]))
					copy(ap_row, ap.M[i])
					ap_ = append(ap_, ap_row)
					row_idx[attr] = append(row_idx[attr], i_)
					i_++
				}
			}
		}

		if len(ap_) > 0 {
			c := computeCoefficients(ap_)
			for attr, rows := range row_idx {
				for _, i_ := range rows {
					userattrs.Coeff[attr] = append(userattrs.Coeff[attr], c[i_])
				}
			}
			log.Debug(">>> userattrs: %v", userattrs.Coeff)
			log.Debug("user access policy %d (%d x %d matrix)", ap_, len(ap_[0]), len(ap_))
			log.Debug("coefficients for user attributes: %d", c)
		} else {
			log.Debug("user access policy empty")
		}
	}
	return userattrs
}

// Json API for SelectUserAttrs
func SelectUserAttrsJson(user string, policy string, userattrsJson string) string {
	userattrs := NewUserAttrsOfJsonStr(userattrsJson).OfJsonObj()
	userattrs.SelectUserAttrs(user, policy)
	return Encode(JsonObjToStr(userattrs.ToJsonObj()))
}
