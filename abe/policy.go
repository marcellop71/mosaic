package abe

import (
  "encoding/json"
  "fmt"
  "strconv"
  "strings"

  antlr "github.com/antlr/antlr4/runtime/Go/antlr"

  log "github.com/unicredit/abe/log"
  parser "github.com/unicredit/abe/parser"
  z3 "github.com/unicredit/abe/z3"
)

// binary tree with string labels and
// vectors for the LSSS representation
type btree struct {
	label string
	child [2]*btree
  v []int
}

type policyListener struct {
	*parser.BasePolicyListener

	s []*btree // stack of trees
  leaves []string // list of leaves
}

func (ln *policyListener) addLeaf(label string, isvalue bool) {
	t := new(btree); t.label = label
  ln.s = append(ln.s, t)
  if (isvalue == false) {
    ln.leaves = append(ln.leaves, label)
  }
}

func (ln *policyListener) addBinaryNode(label string) {
  t := new(btree); t.label = label
	t.child[0] = ln.s[len(ln.s) - 2] // left subtree
	t.child[1] = ln.s[len(ln.s) - 1] // right subtree
	ln.s = ln.s[:len(ln.s) - 2] // pop 2 trees from the stack
	ln.s = append(ln.s, t) // push tree to the stack
}

func (ln *policyListener) addNode(label string, childcount int) {
  for i := 0; i < childcount; i += 2 {
    t := new(btree); t.label = label
		t.child[0] = ln.s[len(ln.s) - 2]
		t.child[1] = ln.s[len(ln.s) - 1]
		ln.s = ln.s[:len(ln.s) - 2]
  	ln.s = append(ln.s, t)
  }
}

func (ln *policyListener) ExitExpr_andor(ctx *parser.Expr_andorContext) {
  switch ctx.GetOp().GetTokenType() {
  case parser.PolicyParserTK_AND:
    ln.addBinaryNode("and")
  case parser.PolicyParserTK_OR:
    ln.addBinaryNode("or")
  }
}

func (ln *policyListener) ExitSexpr_andor(ctx *parser.Sexpr_andorContext) {
  cc := ctx.GetRuleContext().GetChildCount() - 3
  switch ctx.GetOp().GetTokenType() {
  case parser.PolicyParserTK_AND:
    ln.addNode("and", cc)
  case parser.PolicyParserTK_OR:
    ln.addNode("or", cc)
  }
}

func (ln *policyListener) ExitExpr_attrrange0(ctx *parser.Expr_attrrange0Context) {
  s := ctx.GetRuleContext().GetText()
  s_ := strings.Split(s, ctx.GetOp().GetText())

  nbit := 32
  value, _ := strconv.Atoi(s_[1])
  attr := s_[0]
  attrs := GetBagOfBits(attr, value)

  switch ctx.GetOp().GetTokenType() {
  case parser.PolicyParserTK_EQ:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(attrs[i], false)
    }

    for i := nbit - 1; i > 0; i-- {
      ln.addBinaryNode("and")
    }
  case parser.PolicyParserTK_GT:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 1), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("or")
      case '1':
        ln.addBinaryNode("and")
      }
    }
  case parser.PolicyParserTK_GTEQ:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 1), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("or")
      case '1':
        ln.addBinaryNode("and")
      }
    }

    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(attrs[i], false)
    }

    for i := nbit - 1; i > 0; i-- {
      ln.addBinaryNode("and")
    }

    ln.addBinaryNode("or")
  case parser.PolicyParserTK_LT:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 0), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("and")
      case '1':
        ln.addBinaryNode("or")
      }
    }
  case parser.PolicyParserTK_LTEQ:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 0), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("and")
      case '1':
        ln.addBinaryNode("or")
      }
    }

    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(attrs[i], false)
    }

    for i := nbit - 1; i > 0; i-- {
      ln.addBinaryNode("and")
    }

    ln.addBinaryNode("or")
  }
}

func (ln *policyListener) ExitSexpr_attrrange0(ctx *parser.Sexpr_attrrange0Context) {
  s := ctx.GetRuleContext().GetText()
  s_ := strings.Split(s, " ")

  nbit := 32
  value, _ := strconv.Atoi(s_[2])
  attr := s_[1]
  attrs := GetBagOfBits(attr, value)

  switch ctx.GetOp().GetTokenType() {
  case parser.PolicyParserTK_EQ:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(attrs[i], false)
    }

    for i := nbit - 1; i > 0; i-- {
      ln.addBinaryNode("and")
    }
  case parser.PolicyParserTK_GT:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 1), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("or")
      case '1':
        ln.addBinaryNode("and")
      }
    }
  case parser.PolicyParserTK_GTEQ:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 1), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("or")
      case '1':
        ln.addBinaryNode("and")
      }
    }

    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(attrs[i], false)
    }

    for i := nbit - 1; i > 0; i-- {
      ln.addBinaryNode("and")
    }

    ln.addBinaryNode("or")
  case parser.PolicyParserTK_LT:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 0), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("and")
      case '1':
        ln.addBinaryNode("or")
      }
    }
  case parser.PolicyParserTK_LTEQ:
    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(FormatAttr(attr, i, 0), false)
    }

    v := GetBitString(value)
    for i := nbit - 2; i >= 0; i-- {
      switch v[i] {
      case '0':
        ln.addBinaryNode("and")
      case '1':
        ln.addBinaryNode("or")
      }
    }

    for i := nbit - 1; i >= 0; i-- {
      ln.addLeaf(attrs[i], false)
    }

    for i := nbit - 1; i > 0; i-- {
      ln.addBinaryNode("and")
    }

    ln.addBinaryNode("or")
  }
}

func (ln *policyListener) ExitExpr_attrrange1(ctx *parser.Expr_attrrange1Context) {
}

func (ln *policyListener) ExitSexpr_attrrange1(ctx *parser.Sexpr_attrrange1Context) {
}

func (ln *policyListener) ExitExpr_attr(ctx *parser.Expr_attrContext) {
	ln.addLeaf(ctx.GetText(), false)
}

func (ln *policyListener) ExitSexpr_attr(ctx *parser.Sexpr_attrContext) {
	ln.addLeaf(ctx.GetText(), false)
}

func dumpSExpr(t *btree, s string) string {
  internal := (t.child[0] != nil) && (t.child[1] != nil)
  if (internal == true) {
    s += "(" + t.label + " "; s = dumpSExpr(t.child[0], s)
    s += " "; s = dumpSExpr(t.child[1], s); s += ")"
  } else {
    s += fmt.Sprint(t.label)
  }
  return s
}

func parsePolicy(s string) *policyListener {
  lexer := parser.NewPolicyLexer(antlr.NewInputStream(s))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPolicyParser(stream)
	ln := new(policyListener)
	antlr.ParseTreeWalkerDefault.Walk(ln, p.Policy())
  return ln
}

func FormatAttr(attr string, bitidx int, bit int) string {
  s := strings.SplitN(attr, "@", 2)
  return fmt.Sprintf("%s%%%d%%%d@%s", s[0], bitidx, bit, s[1])
}

func GetBitString(x int) string {
  bits := strconv.FormatInt(int64(x), 2)
  nbit := 32
  s := ""
	for bitidx := len(bits); bitidx < nbit; bitidx++ {
    s += "0"
	}
  s += bits
  return s
}

func GetBagOfBits(attr string, value int) []string {
	var attrs []string
  bits := GetBitString(value)
	for bitidx, bit_s := range bits {
		bit, _ := strconv.Atoi(string(bit_s))
		attrs = append(attrs, FormatAttr(attr, bitidx, bit))
	}
  nbit := 32
	for bitidx := len(bits); bitidx < nbit; bitidx++ {
    attrs = append(attrs, FormatAttr(attr, bitidx, 0))
	}
	return attrs
}

func GetBagOfBitsComplete(attr string) []string {
	var attrs []string
	nbit := 32
	for bitidx := 0; bitidx < nbit; bitidx++ {
    attrs = append(attrs, FormatAttr(attr, bitidx, 0))
    attrs = append(attrs, FormatAttr(attr, bitidx, 1))
	}
	return attrs
}

func z3eval(s string) string {
  ctx := z3.NewCtx(z3.NewConf())
  var tmp string
  for _, v := range strings.Split(s, "\n") {
    tmp = ctx.EvalSmt2(v)
  }
  return tmp
}

func simplifyPolicy(policy_sexpr string, leaves_u []string) string {
  smt2 := "(set-logic QF_UF)\n"
  for _, v := range leaves_u {
    smt2 += fmt.Sprintf("(declare-fun %s () Bool)\n", v)
  }
  smt2 += fmt.Sprintf("(simplify %s)", policy_sexpr)
  return z3eval(smt2)
}

func assertPolicy(policy_sexpr string, leaves_u []string) map[string]int {
  smt2 := "(set-option :print-success false)\n"
  smt2 += "(set-logic QF_UF)\n"
  leaves_expr := "("
  for _, v := range leaves_u {
    smt2 += fmt.Sprintf("(declare-fun %s () Bool)\n", v)
    leaves_expr += fmt.Sprintf("%s ", v)
  }
  leaves_expr = leaves_expr[:len(leaves_expr)-1] + ")"
  smt2 += fmt.Sprintf("(assert %s)\n", policy_sexpr)
  smt2 += fmt.Sprintf("(check-sat)\n")
  smt2 += fmt.Sprintf("(get-value %s)", leaves_expr)

  s := z3eval(smt2)
  s = strings.ReplaceAll(s, "\n", "")
  s = strings.ReplaceAll(s, "(", "")
  s = strings.ReplaceAll(s, ")", "")
  s_ := strings.Split(s, " ")
  q := make(map[string]int)
  for i := 0; i < len(s_); i += 2 {
    var tmp int
    if (s_[i+1] == "false") {
      tmp = 0
    }
    if (s_[i+1] == "true") {
      tmp = 1
    }
    q[s_[i]] = tmp
  }
  return q
}

func andorgNotationInSExpr(s string) string {
  s = strings.ReplaceAll(s, "(and", "(/\\ ")
  s = strings.ReplaceAll(s, "(or", "(\\/ ")
  return s
}

func buildAccessPolicyMatrix(policy string) ([][]int, int, *btree) {
  log.Debug("build access policy attributes")
  ln := parsePolicy(policy)

  var c int
  var ap [][]int
	encodeAccessPolicy(ln.s[0], &c, &ap)
	return ap, c, ln.s[0]
}

func CheckPolicy(policy string) string {
  log.Debug("check policy")
  ln := parsePolicy(policy)
  policy_sexpr := dumpSExpr(ln.s[0], "")

  //leaves_u := unique(ln.leaves)
  //policy_assert := assertPolicy(policy_sexpr, leaves_u)
  //policy_sexpr_s := simplifyPolicy(policy_sexpr, leaves_u)
  //policy_sexpr_s = andorgNotationInSExpr(policy_sexpr_s)
  //ln = parsePolicy(policy_sexpr_s)
  //policy_sexpr = dumpSExpr(ln.s[0], "")
  policy_sexpr = andorgNotationInSExpr(policy_sexpr)

  log.Info("policy: %s", policy)
  log.Info("policy sexpr: %s", policy_sexpr)
  //log.Debug("policy sexpr: %s -> %s", policy_sexpr, policy_sexpr_s)
  //log.Debug("policy sexpr (simplified, binary form): %s", policy_sexpr_bin)
  //log.Debug("policy assert: %s", policy_assert)
  return policy_sexpr
}

func ExtractPolicyFromCiphertext(ct_json string) string {
	log.Debug("extract policy from ciphertext")
	var ct_ AbeCipherText_
	json.Unmarshal([]byte(ct_json), &ct_)
	return string(decode(ct_.Policy))
}

func ExtractAttrsFromPolicy(policy string) string {
  log.Debug("extract attributes from policy")
  ln := parsePolicy(policy)

  row := make(map[string][]int)
	for i, v := range ln.leaves {
		row[v] = append(row[v], i)
	}
  log.Debug("policy leaves: %s", ln.leaves)
  log.Debug("map leaves to rows: %s", row)

  return JsonObj2Str(
    &AbePolicyAttrs_{Row: row})
}

func ExtractAuthsFromPolicy(policy string) string {
  log.Debug("extract authorities from policy")
  ln := parsePolicy(policy)

  authpub := make(map[string]string)
	for _, v := range ln.leaves {
    v_ := strings.SplitN(v, "@", 2)[1]
    authpub[v_] = ""
	}
  log.Debug("policy authorities: %s", authpub)

  return JsonObj2Str(
    &AbeAuthPubs_{AuthPub: authpub})
}

func SelectUserAttrs(user string, policy string, userattrs_json string) string {
	log.Debug("select attributes")
  var userattrs AbeUserAttrs_
  json.Unmarshal([]byte(userattrs_json), &userattrs)

  if (len(userattrs.Coeff) > 0) {
    policyattrs_json := ExtractAttrsFromPolicy(policy)
    var policyattrs AbePolicyAttrs_
    json.Unmarshal([]byte(policyattrs_json), &policyattrs)

  	ap, _, _ := buildAccessPolicyMatrix(policy)
  	log.Debug("access policy (unfiltered, %d x %d matrix) %d", len(ap[0]), len(ap), ap)

    // policy matrix intersection with user attributes
  	var ap_ [][]int
    tmp0 := make(map[string][]int)
    i_ := 0
    for attr, _ := range userattrs.Coeff {
      if rows, exist := policyattrs.Row[attr]; exist {
        for _, i := range rows {
          tmp := make([]int, len(ap[i])); copy(tmp, ap[i])
          ap_ = append(ap_, tmp)
          tmp0[attr] = append(tmp0[attr], i_)
          i_++
        }
      }
    }

    if (len(ap_) > 0) {
      log.Debug("user access policy %d (%d x %d matrix)", ap_, len(ap_[0]), len(ap_))

      c := computeCoefficients(ap_)

      for attr, rows := range tmp0 {
        for _, i_ := range rows {
          userattrs.Coeff[attr] = append(userattrs.Coeff[attr], c[i_])
        }
      }

  		log.Debug("coefficients for user attributes: %d", c)
    } else {
      log.Debug("user access policy empty")
    }

  	return JsonObj2Str(
      &AbeUserAttrs_{User: user, Coeff: userattrs.Coeff})
  }
  return JsonObj2Str(
    &AbeUserAttrs_{User: user, Coeff: make(map[string][]int)})
}
