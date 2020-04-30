package abe

import (
	"encoding/json"
	"math/big"
	"strings"

	log "github.com/unicredit/abe/log"
)

// random organization
// (?): random init seeds for hash into group functions ?
// (?): 2 different hash functions into G2 for users and attributes ?
func NewRandomOrg(curve_json string) string {
	var curve_ AbeCurve_
  json.Unmarshal([]byte(curve_json), &curve_)
	curve := curve_.Decode()

	org := new(AbeOrg)
	org.curve = curve
	org.g1 = curve.NewRandomPointOn("G1")
	org.g2 = curve.NewRandomPointOn("G2")
	org.e = curve.Pair(org.g1, org.g2)
	return JsonObj2Str(org.Encode())
}

// random authority
func NewRandomAuth(org_json string) (string, string) {
	var org_ AbeOrg_
  json.Unmarshal([]byte(org_json), &org_)
	org := org_.Decode()
	curve := org.curve

	authprv := &AbeAuthPrv{
		org, curve.NewRandomExp(), curve.NewRandomExp()}

	authpub := &AbeAuthPub{
		org,
		curve.Pow(org.e, authprv.alpha),
		curve.Pow(org.g1, authprv.y)}

	return JsonObj2Str(authpub.Encode()), JsonObj2Str(authprv.Encode())
}

// random userkey
func NewRandomUserkey(user string, attr string, authprv_json string) string {
	var authprv_ AbeAuthPrv_
  json.Unmarshal([]byte(authprv_json), &authprv_)
	authprv := authprv_.Decode()
	org := authprv.org

	huser := org.curve.HashToGroup(user, "G2")
	hattr := org.curve.HashToGroup(attr, "G2")

	t := org.curve.NewRandomExp()

	userkey := new(AbeUserkey)
	userkey.org = org
	userkey.K = org.curve.Pow3(org.g2, authprv.alpha, huser, authprv.y, hattr, t)
	userkey.KP = org.curve.Pow(org.g1, t)

	return JsonObj2Str(userkey.Encode())
}

// random userkey
func NewRandomUserkeyWithValue(user string, attr string, authprv_json string, value int) string {
	GetBagOfBits(attr, value)
	GetBagOfBitsComplete(attr)
	return ""
	//NewRandomUserkey(user, attr_, authprv_json)
}

// random msg
func NewRandomMsg(org_json string) string {
	var org_ AbeOrg_
  json.Unmarshal([]byte(org_json), &org_)
	org := org_.Decode()
	return JsonObj2Str(org.curve.NewRandomPointOn("GT").Encode())
}

// encrypt a msg according to a given policy
// the function requires the public keys of the authorities
// mentioned in the policy
func Encrypt(msg string, policy string, authpubs_json string) string {
	log.Debug("encrypt msg according to policy %s", policy)

	var authpubs AbeAuthPubs_
	json.Unmarshal([]byte(authpubs_json), &authpubs)

	var policyattrs AbePolicyAttrs_
	policyattrs_json := ExtractAttrsFromPolicy(policy)
	json.Unmarshal([]byte(policyattrs_json), &policyattrs)

	ap, _, _ := buildAccessPolicyMatrix(policy)

	// all authorities in the policy from the same organization
	var authpub_ AbeAuthPub_
	for _, v := range authpubs.AuthPub {
		json.Unmarshal([]byte(v), &authpub_)
		break
	}
	authpub := authpub_.Decode()

	org := authpub.org
	curve := authpub.org.curve

	s := curve.NewRandomSecret(len(policyattrs.Row), false)
	w := curve.NewRandomSecret(len(policyattrs.Row), true)
	sh_s := computeShares(s, ap); sh_w := computeShares(w, ap)

	var M_ AbePoint_
	json.Unmarshal([]byte(msg), &M_)
	M := M_.Decode(curve)

	S0 := curve.Pow(org.e, s[0])
	C0 := curve.Mul(M, S0)

	log.Debug("s0 secret %s", s[0])
	for attr, rows := range policyattrs.Row {
		for _, i := range rows {
			log.Debug("share %d of secret s0 for %s: %s", i, attr, sh_s[i])
		}
	}

	C := make(map[string][][]*AbePoint)
	for attr, rows := range policyattrs.Row {
		hattr := curve.HashToGroup(attr, "G2")
		auth := strings.SplitN(attr, "@", 2)[1]
		json.Unmarshal([]byte(authpubs.AuthPub[auth]), &authpub_)
		authpub := authpub_.Decode()
		C[attr] = make([][]*AbePoint, len(rows))
		for k, i := range rows {
			tx := curve.NewRandomExp()

			C[attr][k] = make([]*AbePoint, 4)
			C[attr][k][0] = curve.Pow2(org.e, sh_s[i], authpub.ealpha, tx)
			C[attr][k][1] = curve.Inv(curve.Pow(org.g1, tx))
			C[attr][k][2] = curve.Pow2(org.g1, sh_w[i], authpub.g1y, tx)
			C[attr][k][3] = curve.Pow(hattr, tx)
		}
	}

	ct := &AbeCipherText{org, policy, C0, C}
	return JsonObj2Str(ct.Encode())
}

// decrypt a ciphertext
// the function requires the list of user attributes to use
// and the corresponding collection of userkey
func Decrypt(ct_json string, userattrs_json string, userkeys_json string) string {
	log.Debug("decrypt")

	var ct_ AbeCipherText_
	var userattrs AbeUserAttrs_
	var userkeys AbeUserKeys_

	json.Unmarshal([]byte(ct_json), &ct_)
	json.Unmarshal([]byte(userattrs_json), &userattrs)
  json.Unmarshal([]byte(userkeys_json), &userkeys)

	ct := ct_.Decode()
	org := ct.org
	curve := org.curve
	huser := curve.HashToGroup(userattrs.User, "G2")

	var userkey_ AbeUserkey_

	S0 := curve.UnitOnGroup("GT")
	for attr, cs := range userattrs.Coeff {
		for k, c := range cs {
			if (c != 0) {
				json.Unmarshal([]byte(userkeys.Userkey[attr]), &userkey_)
				userkey := userkey_.Decode()
				tmp := curve.ProdPair(
					[]*AbePoint{ct.C[attr][k][1], ct.C[attr][k][2], userkey.KP},
					[]*AbePoint{userkey.K, huser, ct.C[attr][k][3]})
				tmp = curve.Mul(tmp, ct.C[attr][k][0])
				tmp = curve.Pow(tmp, big.NewInt(int64(c)))
				S0 = curve.Mul(S0, tmp)
			}
		}
	}

	msg := curve.Div(ct.C0, S0)
	return JsonObj2Str(msg.Encode())
}
