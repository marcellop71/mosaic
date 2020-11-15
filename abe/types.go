package abe

import (
	"encoding/base32"
	"encoding/json"
	"math/big"

	"github.com/marcellop71/mosaic/abe/log"
)

// point interface
type Point interface {
	isPoint()
	ToJsonObj() Point
	OfJsonObj(curve Curve) Point
	GetP() string
	GetGroup() string
}

// curve interface
type Curve interface {
	isCurve()
	ToJsonObj() Curve
	OfJsonObj() Curve
	InitRng(seed string) Curve
	NewPointOn(group string) Point
	NewRandomExp() *big.Int
	NewRandomSecret(n int, zerosecret bool) []*big.Int
	NewRandomPointOn(group string) Point
	UnitOnGroup(group string) Point
	HashToGroup(x string, group string) Point
	HashToPow(x string, g Point) Point
	Inv(g1 Point) Point
	Mul(g1 Point, g2 Point) Point
	Div(g1 Point, g2 Point) Point
	Pow(g1 Point, e1 *big.Int) Point
	Pow2(g1 Point, e1 *big.Int, g2 Point, e2 *big.Int) Point
	Pow3(g1 Point, e1 *big.Int, g2 Point, e2 *big.Int, g3 Point, e3 *big.Int) Point
	Pair(g1 Point, g2 Point) Point
	ProdPair(g1 []Point, g2 []Point) Point
}

// randomness source
type RandomSource struct {
	Seed string `json:"seed"`
}

// organization
type Org struct {
	Crv  Curve  `json:"-"`
	Crv_ string `json:"crv"`
	G1   Point  `json:"-"`
	G1_  string `json:"g1"`
	G2   Point  `json:"-"`
	G2_  string `json:"g2"`
	E    Point  `json:"-"`
	E_   string `json:"e"`
}

// authority public params
type AuthPub struct {
	Org     *Org   `json:"-"`
	Org_    string `json:"org"`
	Ealpha  Point  `json:"-"`
	Ealpha_ string `json:"ealpha"`
	G1y     Point  `json:"-"`
	G1y_    string `json:"g1y"`
}

// authority private params
type AuthPrv struct {
	Org    *Org     `json:"-"`
	Org_   string   `json:"org"`
	Alpha  *big.Int `json:"-"`
	Alpha_ string   `json:"alpha"`
	Y      *big.Int `json:"-"`
	Y_     string   `json:"y"`
}

// user key
type Userkey struct {
	Org  *Org   `json:"-"`
	Org_ string `json:"org"`
	K    Point  `json:"-"`
	K_   string `json:"k"`
	KP   Point  `json:"-"`
	KP_  string `json:"kp"`
}

// ciphertext encrypting the msg
type Ciphertext struct {
	Org     *Org                  `json:"-"`
	Org_    string                `json:"org"`
	Policy  string                `json:"-"`
	Policy_ string                `json:"policy"`
	C0      Point                 `json:"-"`
	C0_     string                `json:"c0"`
	C       map[string][][]Point  `json:"-"`
	C_      map[string][][]string `json:"c"`
}

// authority public and private keys
type AuthKeys struct {
	AuthPub  *AuthPub `json:"-"`
	AuthPub_ string   `json:"authpub"` // public key of the authority
	AuthPrv  *AuthPrv `json:"-"`
	AuthPrv_ string   `json:"authprv"` // private key of the authority
}

// map of authorities public keys
type AuthPubs struct {
	AuthPub  map[string]*AuthPub `json:"-"`
	AuthPub_ map[string]string   `json:"authpub"` // authority -> corresponding public key
}

// map of user attributes
type UserAttrs struct {
	User     string              `json:"user"`  // user
	Coeff    map[string][]int    `json:"coeff"` // attribute -> its coefficients
	Userkey  map[string]*Userkey `json:"-"`
	Userkey_ map[string]string   `json:"userkey"` // attribute -> corresponding userkey
}

// attributes in a given policy
type AccessPolicy struct {
	M    [][]int          // access policy matrix
	Vars []string         // attribute list
	Row  map[string][]int // attribute -> access policy matrix rows index
}

func NewPointOfJsonStr(xJson string) (x Point) {
	x = NewPoint()
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewCurveOfJsonStr(xJson string) (x Curve) {
	x = NewCurve()
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewOrgOfJsonStr(xJson string) (x *Org) {
	x = new(Org)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewAuthPubOfJsonStr(xJson string) (x *AuthPub) {
	x = new(AuthPub)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewAuthPrvOfJsonStr(xJson string) (x *AuthPrv) {
	x = new(AuthPrv)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewUserkeyOfJsonStr(xJson string) (x *Userkey) {
	x = new(Userkey)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewCiphertextOfJsonStr(xJson string) (x *Ciphertext) {
	x = new(Ciphertext)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewAuthKeysOfJsonStr(xJson string) (x *AuthKeys) {
	x = new(AuthKeys)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewAuthPubsOfJsonStr(xJson string) (x *AuthPubs) {
	x = new(AuthPubs)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func NewUserAttrsOfJsonStr(xJson string) (x *UserAttrs) {
	x = new(UserAttrs)
	if len(xJson) > 0 {
		err := json.Unmarshal([]byte(Decode(xJson)), x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
	}
	return
}

func Encode(x string) string {
	return base32.StdEncoding.EncodeToString([]byte(x))
}

func Decode(x string) string {
	y, err := base32.StdEncoding.DecodeString(x)
	if err != nil {
		log.Error("%s", err)
	}
	return string(y)
}

func JsonObjToStr(x interface{}) string {
	x_ := ""
	if x != nil {
		x__, err := json.Marshal(x)
		if err != nil {
			log.Error("json unmarshal error: %s", err)
		}
		x_ = string(x__)
	}
	return x_
}

func JsonObjToEncStr(x interface{}) string {
	return Encode(JsonObjToStr(x))
}

// Org type to its json representation
func (org *Org) ToJsonObj() *Org {
	org.Crv_ = JsonObjToEncStr(org.Crv.ToJsonObj())
	org.G1_ = JsonObjToEncStr(org.G1.ToJsonObj())
	org.G2_ = JsonObjToEncStr(org.G2.ToJsonObj())
	org.E_ = JsonObjToEncStr(org.E.ToJsonObj())
	return org
}

// Org type from its json representation
func (org *Org) OfJsonObj() *Org {
	org.Crv = NewCurveOfJsonStr(org.Crv_).OfJsonObj()
	org.G1 = NewPointOfJsonStr(org.G1_).OfJsonObj(org.Crv)
	org.G2 = NewPointOfJsonStr(org.G2_).OfJsonObj(org.Crv)
	org.E = NewPointOfJsonStr(org.E_).OfJsonObj(org.Crv)
	return org
}

// AuthPub type to its json representation
func (authpub *AuthPub) ToJsonObj() *AuthPub {
	authpub.Org_ = JsonObjToEncStr(authpub.Org.ToJsonObj())
	authpub.Ealpha_ = JsonObjToEncStr(authpub.Ealpha.ToJsonObj())
	authpub.G1y_ = JsonObjToEncStr(authpub.G1y.ToJsonObj())
	return authpub
}

// AuthPub type from its json representation
func (authpub *AuthPub) OfJsonObj() *AuthPub {
	authpub.Org = NewOrgOfJsonStr(authpub.Org_).OfJsonObj()
	authpub.Ealpha = NewPointOfJsonStr(authpub.Ealpha_).OfJsonObj(authpub.Org.Crv)
	authpub.G1y = NewPointOfJsonStr(authpub.G1y_).OfJsonObj(authpub.Org.Crv)
	return authpub
}

// AuthPrv type to its json representation
func (authprv *AuthPrv) ToJsonObj() *AuthPrv {
	authprv.Org_ = JsonObjToEncStr(authprv.Org.ToJsonObj())
	authprv.Alpha_ = Encode(string(authprv.Alpha.Bytes()))
	authprv.Y_ = Encode(string(authprv.Y.Bytes()))
	return authprv
}

// AuthPrv type from its json representation
func (authprv *AuthPrv) OfJsonObj() *AuthPrv {
	authprv.Org = NewOrgOfJsonStr(authprv.Org_).OfJsonObj()
	authprv.Alpha = new(big.Int).SetBytes([]byte(Decode(authprv.Alpha_)))
	authprv.Y = new(big.Int).SetBytes([]byte(Decode(authprv.Y_)))
	return authprv
}

// Userkey type to its json representation
func (userkey *Userkey) ToJsonObj() *Userkey {
	userkey.Org_ = JsonObjToEncStr(userkey.Org.ToJsonObj())
	userkey.K_ = JsonObjToEncStr(userkey.K.ToJsonObj())
	userkey.KP_ = JsonObjToEncStr(userkey.KP.ToJsonObj())
	return userkey
}

// Userkey type from its json representation
func (userkey *Userkey) OfJsonObj() *Userkey {
	userkey.Org = NewOrgOfJsonStr(userkey.Org_).OfJsonObj()
	userkey.K = NewPointOfJsonStr(userkey.K_).OfJsonObj(userkey.Org.Crv)
	userkey.KP = NewPointOfJsonStr(userkey.KP_).OfJsonObj(userkey.Org.Crv)
	return userkey
}

// Ciphertext type to its json representation
func (ct *Ciphertext) ToJsonObj() *Ciphertext {
	ct.Org_ = JsonObjToEncStr(ct.Org.ToJsonObj())
	ct.Policy_ = Encode(ct.Policy)
	ct.C0_ = JsonObjToEncStr(ct.C0.ToJsonObj())
	ct.C_ = make(map[string][][]string)
	for attr, v := range ct.C {
		ct.C_[attr] = make([][]string, len(v))
		for k, w := range v {
			ct.C_[attr][k] = make([]string, 4)
			for i := 0; i < 4; i++ {
				ct.C_[attr][k][i] = JsonObjToEncStr(w[i].ToJsonObj())
			}
		}
	}
	return ct
}

// Ciphertext type from its json representation
func (ct *Ciphertext) OfJsonObj() *Ciphertext {
	ct.Org = NewOrgOfJsonStr(ct.Org_).OfJsonObj()
	ct.Policy = Decode(ct.Policy_)
	ct.C0 = NewPointOfJsonStr(ct.C0_).OfJsonObj(ct.Org.Crv)

	ct.C = make(map[string][][]Point)
	for attr, v := range ct.C_ {
		ct.C[attr] = make([][]Point, len(v))
		for k, w := range v {
			ct.C[attr][k] = make([]Point, 4)
			for i := 0; i < 4; i++ {
				ctmp := NewPointOfJsonStr(w[i]).OfJsonObj(ct.Org.Crv)
				ct.C[attr][k][i] = ctmp
			}
		}
	}
	return ct
}

// AuthKeys type to its json representation
func (authkeys *AuthKeys) ToJsonObj() *AuthKeys {
	authkeys.AuthPub_ = JsonObjToEncStr(authkeys.AuthPub.ToJsonObj())
	authkeys.AuthPrv_ = JsonObjToEncStr(authkeys.AuthPrv.ToJsonObj())
	return authkeys
}

// AuthKeys type from its json representation
func (authkeys *AuthKeys) OfJsonObj() *AuthKeys {
	authkeys.AuthPub = NewAuthPubOfJsonStr(Decode(authkeys.AuthPub_)).OfJsonObj()
	authkeys.AuthPrv = NewAuthPrvOfJsonStr(Decode(authkeys.AuthPrv_)).OfJsonObj()
	return authkeys
}

// AuthPubs type to its json representation
func (authpubs *AuthPubs) ToJsonObj() *AuthPubs {
	authpubs.AuthPub_ = make(map[string]string)
	for attr, authpub := range authpubs.AuthPub {
		tmp := Encode("{}")
		if authpub != nil {
			tmp = JsonObjToEncStr(authpub.ToJsonObj())
		}
		authpubs.AuthPub_[attr] = tmp
	}
	return authpubs
}

// AuthPubs type from its json representation
func (authpubs *AuthPubs) OfJsonObj() *AuthPubs {
	authpubs.AuthPub = make(map[string]*AuthPub)
	for attr, authpubJson := range authpubs.AuthPub_ {
		authpubs.AuthPub[attr] = NewAuthPubOfJsonStr(authpubJson).OfJsonObj()
	}
	return authpubs
}

// UserAttrs type to its json representation
func (userattrs *UserAttrs) ToJsonObj() *UserAttrs {
	userattrs.Userkey_ = make(map[string]string)
	for attr, userkey := range userattrs.Userkey {
		tmp := Encode("{}")
		if userkey != nil {
			tmp = JsonObjToEncStr(userkey.ToJsonObj())
		}
		userattrs.Userkey_[attr] = tmp
	}
	return userattrs
}

// UserAttrs type from its json representation
func (userattrs *UserAttrs) OfJsonObj() *UserAttrs {
	userattrs.Userkey = make(map[string]*Userkey)
	for attr, userkeyJson := range userattrs.Userkey_ {
		userattrs.Userkey[attr] = NewUserkeyOfJsonStr(userkeyJson).OfJsonObj()
	}
	return userattrs
}

// merge 2 UserAttrs
func (userattrs0 *UserAttrs) Add(userattrs1 *UserAttrs) *UserAttrs {
	for attr, userkey := range userattrs1.Userkey {
		userattrs0.Coeff[attr] = userattrs1.Coeff[attr]
		userattrs0.Userkey[attr] = userkey
	}
	return userattrs0
}
