package main

import "C"

import (
  "github.com/unicredit/mosaic/abe"
)

//export newRandomOrg
func newRandomOrg(curveJson *C.char) *C.char {
  return C.CString(abe.NewRandomOrgJson(C.GoString(curveJson)))
}

//export newRandomAuth
func newRandomAuth(orgJson *C.char) *C.char {
  return C.CString(abe.NewRandomAuthJson(C.GoString(orgJson)))
}

//export newRandomUserkey
func newRandomUserkey(user *C.char, attr *C.char, authprvJson *C.char) *C.char {
  return C.CString(abe.NewRandomUserkeyJson(C.GoString(user), C.GoString(attr), C.GoString(authprvJson)))
}

//export newRandomSecret
func newRandomSecret(orgJson *C.char, seed *C.char) *C.char {
  return C.CString(abe.NewRandomSecretJson(C.GoString(orgJson), C.GoString(seed)))
}

//export rewritePolicy
func rewritePolicy(policy *C.char) *C.char {
  return C.CString(abe.RewritePolicy(C.GoString(policy)))
}

//export checkPolicy
func checkPolicy(policy *C.char, userattrsJson *C.char) *C.char {
  return C.CString(abe.CheckPolicyJson(C.GoString(policy), C.GoString(userattrsJson)))
}

//export authpubsOfPolicy
func authpubsOfPolicy(policy *C.char) *C.char {
  return C.CString(abe.AuthPubsOfPolicyJson(C.GoString(policy)))
}

//export policyOfCiphertextJson
func policyOfCiphertextJson(ctJson *C.char) *C.char {
  return C.CString(abe.PolicyOfCiphertextJson(C.GoString(ctJson)))
}

//export selectUserAttrs
func selectUserAttrs(user *C.char, policy *C.char, userattrsJson *C.char) *C.char {
  return C.CString(abe.SelectUserAttrsJson(C.GoString(user), C.GoString(policy), C.GoString(userattrsJson)))
}

//export encrypt
func encrypt(secret *C.char, policy *C.char, authpubsJson *C.char) *C.char {
  return C.CString(abe.EncryptJson(C.GoString(secret), C.GoString(policy), C.GoString(authpubsJson)))
}

//export decrypt
func decrypt(ctJson *C.char, userattrsJson *C.char) *C.char {
  return C.CString(abe.DecryptJson(C.GoString(ctJson), C.GoString(userattrsJson)))
}

//export getPbcCurve
func getPbcCurve(curveStr *C.char) *C.char {
  return C.CString(abe.GetPbcCurve(C.GoString(curveStr)))
}

func main() {}
