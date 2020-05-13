package abe

import (
	"math/big"
	"strings"

	"github.com/unicredit/mosaic/abe/log"
)

// new random organization
func NewRandomOrg(curve Curve) (org *Org) {
	org = &Org{
		Crv: curve,
		G1: curve.NewRandomPointOn("G1"),
		G2: curve.NewRandomPointOn("G2"),
	}
	org.E = curve.Pair(org.G1, org.G2)
	return
}

// Json API for NewRandomOrg
func NewRandomOrgJson(curveJson string) string {
	curve := NewCurveOfJsonStr(curveJson).OfJsonObj()
  org := NewRandomOrg(curve)
	return Encode(JsonObjToStr(org.ToJsonObj()))
}

// new random authority
func NewRandomAuth(org *Org) (authkeys *AuthKeys) {
	authprv := &AuthPrv{
		Org: org,
		Alpha: org.Crv.NewRandomExp(),
		Y: org.Crv.NewRandomExp()}

	authpub := &AuthPub{
		Org: org,
		Ealpha: org.Crv.Pow(org.E, authprv.Alpha),
		G1y: org.Crv.Pow(org.G1, authprv.Y)}

	authkeys = &AuthKeys{
		AuthPub: authpub, AuthPrv: authprv}
	return
}

// Json API for NewRandomAuth
func NewRandomAuthJson(orgJson string) string {
	org := NewOrgOfJsonStr(orgJson).OfJsonObj()
  authkeys := NewRandomAuth(org)
	return Encode(JsonObjToStr(authkeys.ToJsonObj()))
}

// new random userkey
func NewRandomUserkey(user string, attr string, authprv *AuthPrv) (userattrs *UserAttrs) {
	org := authprv.Org

	userattrs = new(UserAttrs)
	userattrs.User = user
	userattrs.Userkey = make(map[string]*Userkey)
	userattrs.Coeff = make(map[string][]int)

	nbit := 32
	attrs := getBagOfBitsAttrs(attr, nbit)
	for _, attr := range attrs {
		huser := org.Crv.HashToGroup(user, "G2")
		hattr := org.Crv.HashToGroup(attr, "G2")
		t := org.Crv.NewRandomExp()

		userkey := &Userkey{
			Org: org,
			K: org.Crv.Pow3(org.G2, authprv.Alpha, huser, authprv.Y, hattr, t),
			KP: org.Crv.Pow(org.G1, t),
		}
		userattrs.Userkey[attr] = userkey
		userattrs.Coeff[attr] = []int{}
	}
	return
}

// Json API for NewRandomUserkey
func NewRandomUserkeyJson(user string, attr string, authprvJson string) string {
	authprv := NewAuthPrvOfJsonStr(authprvJson).OfJsonObj()
  userattrs := NewRandomUserkey(user, attr, authprv)
	return Encode(JsonObjToStr(userattrs.ToJsonObj()))
}

// new random secret
func NewRandomSecret(org *Org, seed string) (secret Point) {
	if (len(seed) == 0) {
		secret = org.Crv.NewRandomPointOn("GT")
	} else {
		secret = org.Crv.HashToGroup(seed, "GT")
	}
	return secret
}

// Json API for NewRandomSecret
func NewRandomSecretJson(orgJson string, seed string) string {
	org := NewOrgOfJsonStr(orgJson).OfJsonObj()
	p := NewRandomSecret(org, seed)
  return Encode(JsonObjToStr(p.ToJsonObj()))
}

func GetOrgFromAuthPubs(authpubs *AuthPubs) *Org {
	// all authorities in the policy from the same organization
	authpub := new(AuthPub)
	for _, authpubtmp := range authpubs.AuthPub {
		authpub = authpubtmp
		break
	}
	return authpub.Org
}

// encrypt a secret according to a given policy
// the function requires the public keys of the authorities mentioned in the policy
func Encrypt(secret Point, policy string, authpubs *AuthPubs) (ct *Ciphertext) {
	org := GetOrgFromAuthPubs(authpubs)
	curve := org.Crv

	ap := buildAccessPolicy(policy)
	s := curve.NewRandomSecret(len(ap.Vars), false)
	w := curve.NewRandomSecret(len(ap.Vars), true)
	sh_s := computeShares(s, ap.M)
	sh_w := computeShares(w, ap.M)
	log.Debug("s0 secret %s", s[0])
	for attr, rows := range ap.Row {
		log.Debug("shares of secret s0 for %s", attr)
		for _, i := range rows {
			log.Debug("share on row %d: %s", i, sh_s[i])
		}
	}

	S0 := curve.Pow(org.E, s[0])
	C0 := curve.Mul(secret, S0)

	C := make(map[string][][]Point)
	for attr, rows := range ap.Row {
		auth := strings.SplitN(attr, "@", 2)[1]
		authpub := authpubs.AuthPub[auth]
		hattr := curve.HashToGroup(attr, "G2")

		C[attr] = make([][]Point, len(rows))
		for k, i := range rows {
			tx := curve.NewRandomExp()

			C[attr][k] = make([]Point, 4)
			C[attr][k][0] = curve.Pow2(org.E, sh_s[i], authpub.Ealpha, tx)
			C[attr][k][1] = curve.Inv(curve.Pow(org.G1, tx))
			C[attr][k][2] = curve.Pow2(org.G1, sh_w[i], authpub.G1y, tx)
			C[attr][k][3] = curve.Pow(hattr, tx)
		}
	}

	ct = &Ciphertext{
		Org: org, Policy: policy, C0: C0, C: C,}
	return
}

// Json API for Encrypt
func EncryptJson(secretJson string, policy string, authpubsJson string) string {
	secret := NewPointOfJsonStr(secretJson)
  authpubs := NewAuthPubsOfJsonStr(authpubsJson).OfJsonObj()
	org := GetOrgFromAuthPubs(authpubs)
	secret.OfJsonObj(org.Crv)
  ct := Encrypt(secret, policy, authpubs)
	return Encode(JsonObjToStr(ct.ToJsonObj()))
}

// decrypt a ciphertext
// the function requires the list of user attributes to use
// and the corresponding collection of userkey
func Decrypt(ct *Ciphertext, userattrs *UserAttrs) (secret Point) {
	org := ct.Org
	curve := org.Crv
	huser := curve.HashToGroup(userattrs.User, "G2")

	S0 := curve.UnitOnGroup("GT")
	for attr, cs := range userattrs.Coeff {
		for k, c := range cs {
			if (c != 0) {
				userkey := userattrs.Userkey[attr]
				tmp := curve.ProdPair(
					[]Point{ct.C[attr][k][1], ct.C[attr][k][2], userkey.KP},
					[]Point{userkey.K, huser, ct.C[attr][k][3]})
				tmp = curve.Mul(tmp, ct.C[attr][k][0])
				tmp = curve.Pow(tmp, big.NewInt(int64(c)))
				S0 = curve.Mul(S0, tmp)
			}
		}
	}

	secret = curve.Div(ct.C0, S0)
	return
}

// Json API for Decrypt
func DecryptJson(ctJson string, userattrsJson string) string {
  ct := NewCiphertextOfJsonStr(ctJson).OfJsonObj()
	userattrs := NewUserAttrsOfJsonStr(userattrsJson).OfJsonObj()
  secret := Decrypt(ct, userattrs)
	return secret.ToJsonObj().GetP()
}

func PolicyOfCiphertext(ct *Ciphertext) string {
	return ct.Policy
}

// Json API extracting policy from ciphertexts
func PolicyOfCiphertextJson(ctJson string) string {
	ct := NewCiphertextOfJsonStr(ctJson).OfJsonObj()
	return ct.Policy
}
