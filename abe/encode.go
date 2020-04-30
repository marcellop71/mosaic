package abe

import (
  "encoding/json"
  "math/big"

  log "github.com/unicredit/abe/log"
)

func (org *AbeOrg) Encode() (*AbeOrg_) {
  log.Debug("encoding organization")
	return &AbeOrg_{
		encode([]byte(JsonObj2Str(org.curve.Encode()))),
		encode([]byte(JsonObj2Str(org.g1.Encode()))),
    encode([]byte(JsonObj2Str(org.g2.Encode()))),
    encode([]byte(JsonObj2Str(org.e.Encode())))}
}

func (org_ *AbeOrg_) Decode() (*AbeOrg) {
  log.Debug("decoding organization")
  var curve_ AbeCurve_
  var g1_, g2_, e_ AbePoint_

  json.Unmarshal(decode(org_.Curve), &curve_)
	json.Unmarshal(decode(org_.G1), &g1_)
  json.Unmarshal(decode(org_.G2), &g2_)
  json.Unmarshal(decode(org_.E), &e_)

  curve := curve_.Decode()
	return &AbeOrg{
    curve, g1_.Decode(curve), g2_.Decode(curve), e_.Decode(curve)}
}

func (authpub *AbeAuthPub) Encode() (*AbeAuthPub_) {
  log.Debug("encoding authority (public params)")
	return &AbeAuthPub_{
		encode([]byte(JsonObj2Str(authpub.org.Encode()))),
    encode([]byte(JsonObj2Str(authpub.ealpha.Encode()))),
    encode([]byte(JsonObj2Str(authpub.g1y.Encode())))}
}

func (authpub_ *AbeAuthPub_) Decode() (*AbeAuthPub) {
  log.Debug("decoding authority (public params)")
  var org_ AbeOrg_
  var ealpha_, g1y_ AbePoint_
  json.Unmarshal(decode(authpub_.Org), &org_)
	json.Unmarshal(decode(authpub_.Ealpha), &ealpha_)
  json.Unmarshal(decode(authpub_.G1y), &g1y_)

  org := org_.Decode()
	return &AbeAuthPub{
		org, ealpha_.Decode(org.curve), g1y_.Decode(org.curve)}
}

func (authprv *AbeAuthPrv) Encode() (*AbeAuthPrv_) {
  log.Debug("encoding authority (private params)")
	return &AbeAuthPrv_{
    encode([]byte(JsonObj2Str(authprv.org.Encode()))),
		encode(authprv.alpha.Bytes()),
    encode(authprv.y.Bytes())}
}

func (authprv_ *AbeAuthPrv_) Decode() (*AbeAuthPrv) {
  log.Debug("decoding authority (private params)")
  var org_ AbeOrg_
  json.Unmarshal(decode(authprv_.Org), &org_)
	return &AbeAuthPrv{
		org_.Decode(),
		new(big.Int).SetBytes(decode(authprv_.Alpha)),
		new(big.Int).SetBytes(decode(authprv_.Y))}
}

func (userkey *AbeUserkey) Encode() (*AbeUserkey_) {
  log.Debug("encoding user keys")
	return &AbeUserkey_{
    encode([]byte(JsonObj2Str(userkey.org.Encode()))),
    encode([]byte(JsonObj2Str(userkey.K.Encode()))),
    encode([]byte(JsonObj2Str(userkey.KP.Encode())))}
}

func (userkey_ *AbeUserkey_) Decode() (*AbeUserkey) {
  log.Debug("decoding user keys")
  var org_ AbeOrg_
  var K_, KP_ AbePoint_
  json.Unmarshal(decode(userkey_.Org), &org_)
	json.Unmarshal(decode(userkey_.K), &K_)
  json.Unmarshal(decode(userkey_.KP), &KP_)

  org := org_.Decode()
	return &AbeUserkey{
		org, K_.Decode(org.curve), KP_.Decode(org.curve)}
}

func (ct *AbeCipherText) Encode() (*AbeCipherText_) {
  log.Debug("encoding ciphertext")
	ct_ := new(AbeCipherText_)
	ct_.Org = encode([]byte(JsonObj2Str(ct.org.Encode())))
	ct_.Policy = encode([]byte(ct.policy))
	ct_.C0 = encode([]byte(JsonObj2Str(ct.C0.Encode())))
	ct_.C = make(map[string][][]string)
	for attr, v := range ct.C {
    ct_.C[attr] = make([][]string, len(v))
    for k, w := range v {
      ct_.C[attr][k] = make([]string, 4)
  		for i := 0; i < 4; i++ {
  			ct_.C[attr][k][i] = encode([]byte(JsonObj2Str(w[i].Encode())))
  		}
    }
	}
	return ct_
}

func (ct_ *AbeCipherText_) Decode() (*AbeCipherText) {
  log.Debug("decoding ciphertext")
  var org_ AbeOrg_
  json.Unmarshal(decode(ct_.Org), &org_)
  org := org_.Decode()

  var tmp_ AbePoint_
	ct := new(AbeCipherText)
	ct.org = org_.Decode()
	ct.policy = string(decode(ct_.Policy))
  json.Unmarshal(decode(ct_.C0), &tmp_)
	ct.C0 = tmp_.Decode(org.curve)
	ct.C = make(map[string][][]*AbePoint)
	for attr, v := range ct_.C {
    ct.C[attr] = make([][]*AbePoint, len(v))
    for k, w := range v {
  		ct.C[attr][k] = make([]*AbePoint, 4)
      for i := 0; i < 4; i++ {
        json.Unmarshal(decode(w[i]), &tmp_)
        ct.C[attr][k][i] = tmp_.Decode(org.curve)
      }
    }
	}
	return ct
}
