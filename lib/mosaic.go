package main

import "C"

import (
  "fmt"

  abe "github.com/unicredit/abe/abe"
)

//export newRandomCurve
func newRandomCurve(curveName *C.char) *C.char {
  return C.CString(abe.NewRandomCurve(C.GoString(curveName)))
}

//export newRandomOrg
func newRandomOrg(curve_json *C.char) *C.char {
  return C.CString(abe.NewRandomOrg(C.GoString(curve_json)))
}

//export newRandomAuth
func newRandomAuth(org_json *C.char) *C.char {
  pub, prv := abe.NewRandomAuth(C.GoString(org_json))
  return C.CString(fmt.Sprintf("%s|%s", pub, prv))
}

//export newRandomUserkey
func newRandomUserkey(user *C.char, attr *C.char, authprv_json *C.char) *C.char {
  return C.CString(abe.NewRandomUserkey(C.GoString(user), C.GoString(attr), C.GoString(authprv_json)))
}

//export newRandomMsg
func newRandomMsg(org_json *C.char) *C.char {
  return C.CString(abe.NewRandomMsg(C.GoString(org_json)))
}

//export checkPolicy
func checkPolicy(policy *C.char) *C.char {
  return C.CString(abe.CheckPolicy(C.GoString(policy)))
}

//export extractAuthsFromPolicy
func extractAuthsFromPolicy(policy *C.char) *C.char {
  return C.CString(abe.ExtractAuthsFromPolicy(C.GoString(policy)))
}

//export extractPolicyFromCiphertext
func extractPolicyFromCiphertext(policy *C.char) *C.char {
  return C.CString(abe.ExtractPolicyFromCiphertext(C.GoString(policy)))
}

//export selectUserAttrs
func selectUserAttrs(user *C.char, policy *C.char, userattrs_json *C.char) *C.char {
  return C.CString(abe.SelectUserAttrs(C.GoString(user), C.GoString(policy), C.GoString(userattrs_json)))
}

//export encrypt
func encrypt(msg *C.char, policy *C.char, policyattrs_json *C.char) *C.char {
  return C.CString(abe.Encrypt(C.GoString(msg), C.GoString(policy), C.GoString(policyattrs_json)))
}

//export decrypt
func decrypt(ct_json *C.char, userattrs_json *C.char, userkeys_json *C.char) *C.char {
  return C.CString(abe.Decrypt(C.GoString(ct_json), C.GoString(userattrs_json), C.GoString(userkeys_json)))
}

func main() {}
