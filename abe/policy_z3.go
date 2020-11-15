// +build z3

package abe

import (
	"fmt"
	"strings"

	"github.com/unicredit/mosaic/abe/log"
	"github.com/unicredit/mosaic/abe/z3"
)

// evals on z3 library via smt2
func z3eval(s string) string {
	ctx := z3.NewCtx(z3.NewConf())
	var tmp string
	for _, v := range strings.Split(s, "\n") {
		tmp = ctx.EvalSmt2(v)
	}
	return tmp
}

// and, or notation as /\\, \\/
func andorgNotationInSExpr(s string) string {
	s = strings.ReplaceAll(s, "(and", "(/\\ ")
	s = strings.ReplaceAll(s, "(or", "(\\/ ")
	return s
}

// builds smt2 code to simplify a policy
func simplifyPolicyCode(policy_sexpr string, vars_t map[string]string) string {
	smt2 := "(set-option :print-success false)\n"
	smt2 += "(set-logic QF_UF)\n"
	smt2 += "(set-info :smt-lib-version 2.0)\n"
	for v, _ := range vars_t {
		smt2 += fmt.Sprintf("(declare-fun %s () Bool)\n", v)
	}
	smt2 += fmt.Sprintf("(simplify %s)", policy_sexpr)
	return smt2
}

// rewrites a policy in a simplified form
func RewritePolicy(policy string) string {
	method := 0
	ln := parsePolicy(policy, method)
	policy_sexpr := dumpSExpr(ln.s[0], "")
	smt2 := simplifyPolicyCode(policy_sexpr, ln.vars_t)
	log.Debug("%s", smt2)
	policy_sexpr_s := z3eval(smt2)
	policy_sexpr_s = andorgNotationInSExpr(policy_sexpr_s)
	ln = parsePolicy(policy_sexpr_s, method)
	policy_sexpr_ := dumpSExpr(ln.s[0], "")
	policy_sexpr_ = andorgNotationInSExpr(policy_sexpr_)
	log.Debug("policy: %s", policy)
	log.Debug("policy sexpr: %s", policy_sexpr)
	log.Debug("policy sexpr: %s -> %s", policy_sexpr, policy_sexpr_s)
	log.Debug("policy sexpr (simplified, binary form): %s", policy_sexpr_)
	return policy_sexpr_
}

// builds smt2 code to check satisfiability of a policy
func checkPolicyCode(policy_sexpr string, vars_t map[string]string, userattrs *UserAttrs) string {
	smt2 := "(set-option :print-success false)\n"
	smt2 += "(set-logic QF_LIA)\n"
	smt2 += "(set-info :smt-lib-version 2.0)\n"

	cons := ""
	for v, t := range vars_t {
		smt2 += fmt.Sprintf("(declare-fun %s () %s)\n", v, t)
		if userattrs != nil {
			if _, exist := userattrs.Coeff[v]; exist {
			} else {
				switch t {
				case "Bool":
					cons += fmt.Sprintf("(assert (not %s))\n", v)
				}
			}
		}
	}
	smt2 += fmt.Sprintf("(assert %s)\n", policy_sexpr)
	smt2 += cons
	smt2 += fmt.Sprintf("(check-sat)")
	log.Debug(">>> %s", smt2)
	return smt2
}

// checks satisfiability of a policy
func CheckPolicy(policy string, userattrs *UserAttrs) string {
	method := 1
	ln := parsePolicy(policy, method)
	policy_sexpr := dumpSExpr(ln.s[0], "")
	log.Debug("policy sexpr: %s", policy_sexpr)
	smt2 := checkPolicyCode(policy_sexpr, ln.vars_t, userattrs)
	s := z3eval(smt2)
	return s[0 : len(s)-1]
}

// Json API for CheckPolicy
func CheckPolicyJson(policy string, userattrsJson string) string {
	var userattrs *UserAttrs = nil
	if len(userattrsJson) > 0 {
		userattrs = NewUserAttrsOfJsonStr(userattrsJson).OfJsonObj()
	}
	return CheckPolicy(policy, userattrs)
}

// solves a policy
func SolvePolicy(policy string, userattrs *UserAttrs) map[string]int {
	method := 1
	ln := parsePolicy(policy, method)
	policy_sexpr := dumpSExpr(ln.s[0], "")
	smt2 := checkPolicyCode(policy_sexpr, ln.vars_t, userattrs)

	vars_expr := ""
	for attr, _ := range userattrs.Coeff {
		if _, exist := ln.vars_t[attr]; exist {
			vars_expr += fmt.Sprintf("%s ", attr)
		}
	}
	vars_expr = vars_expr[:len(vars_expr)-1]
	smt2 += fmt.Sprintf("\n(get-value (%s))", vars_expr)
	log.Debug(">>> %s", smt2)

	s := z3eval(smt2)
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s_ := strings.Split(s, " ")
	q := make(map[string]int)
	for i := 0; i < len(s_); i += 2 {
		var tmp int
		if s_[i+1] == "false" {
			tmp = 0
		}
		if s_[i+1] == "true" {
			tmp = 1
		}
		q[s_[i]] = tmp
	}
	log.Info("policy solve: %s", q)
	return q
}
