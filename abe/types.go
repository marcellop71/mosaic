package abe

import (
	"math/big"
	"strconv"

	pbc "github.com/unicredit/abe/pbc"
)

// curve and pairing
type AbeCurve struct {
	pairing   *pbc.Pairing
	params    *pbc.Params
	maxrndexp *big.Int
}

type AbeCurve_ struct {
	Params    string `json:"params"`
	Maxrndexp string `json:"maxrndexp"`
}

// point on a group
type AbePoint struct {
	p     *pbc.Element
	group string
}

type AbePoint_ struct {
	P     string `json:"P"`
	Group string `json:"group"`
}

// organization
type AbeOrg struct {
	curve *AbeCurve
	g1    *AbePoint
	g2    *AbePoint
	e     *AbePoint
}

type AbeOrg_ struct {
	Curve string `json:"curve"`
	G1    string `json:"g1"`
	G2    string `json:"g2"`
	E     string `json:"e"`
}

// authority public params
type AbeAuthPub struct {
	org    *AbeOrg
	ealpha *AbePoint
	g1y    *AbePoint
}

type AbeAuthPub_ struct {
	Org    string `json:"org"`
	Ealpha string `json:"ealpha"`
	G1y    string `json:"g1y"`
}

// authority private params
type AbeAuthPrv struct {
	org   *AbeOrg
	alpha *big.Int
	y     *big.Int
}

type AbeAuthPrv_ struct {
	Org   string `json:"org"`
	Alpha string `json:"alpha"`
	Y     string `json:"y"`
}

// user key
// key is attached to user and attr
// key is private although it is not really a private key beacuse
// it is forgeable using authority private params
type AbeUserkey struct {
	org *AbeOrg
	K   *AbePoint
	KP  *AbePoint
}

type AbeUserkey_ struct {
	Org string `json:"org"`
	K   string `json:"K"`
	KP  string `json:"KP"`
}

// ciphertext encrypting the msg
type AbeCipherText struct {
	org    *AbeOrg
	policy string
	C0     *AbePoint
	C      map[string][][]*AbePoint
}

type AbeCipherText_ struct {
	Org    string                `json:"org"`
	Policy string                `json:"policy"`
	C0     string                `json:"C0"`
	C      map[string][][]string `json:"C"`
}

type AbeAuthPubs_ struct {
	AuthPub map[string]string `json:"authpub"` // authority -> public key of the corresponding authority
}

// attributes in a given policy
type AbePolicyAttrs_ struct {
	Row  map[string][]int `json:"row"`  // attribute -> access policy matrix rows index
}

// map of user attributes
type AbeUserAttrs_ struct {
	User string        	   `json:"user"` // user
	Coeff map[string][]int `json:"coeff"` // attribute -> coefficent for user attribute
}

// map of user keys
type AbeUserKeys_ struct {
	User string            		`json:"user"`    // user
	Userkey map[string]string `json:"userkey"` // attribute -> userkey for the attribute
}

type Point interface {
	Encode() *AbePoint_
	Decode(curve *AbeCurve) *AbePoint
}

type Curve interface {
	Encode() *AbeCurve_
	Decode() *AbeCurve
	NewPointOn(group string) *AbePoint
	NewRandomExp() *big.Int
	NewRandomSecret(n int, zerosecret bool) []*big.Int
	NewRandomPointOn(group string) *AbePoint
	UnitOnGroup(group string) *AbePoint
	HashToGroup(x string, group string) *AbePoint
	HashToPow(x string, g *AbePoint) *AbePoint
	Inv(g1 *AbePoint) *AbePoint
	Mul(g1 *AbePoint, g2 *AbePoint) *AbePoint
	Div(g1 *AbePoint, g2 *AbePoint) *AbePoint
	Pow(g1 *AbePoint, e1 *big.Int) *AbePoint
	Pow2(g1 *AbePoint, e1 *big.Int, g2 *AbePoint, e2 *big.Int) *AbePoint
	Pow3(g1 *AbePoint, e1 *big.Int, g2 *AbePoint, e2 *big.Int, g3 *AbePoint, e3 *big.Int) *AbePoint
	Pair(g1 *AbePoint, g2 *AbePoint) *AbePoint
	ProdPair(g1 []*AbePoint, g2 []*AbePoint) *AbePoint
}

func NewRandomCurve(curveName string) string {
	curve := new(AbeCurve)

	if curveName[0:2] == "BN" {
		bits, _ := strconv.ParseInt(string(curveName[2:]), 10, 32)
		curve.params = pbc.GenerateF(uint32(bits))
		curve.maxrndexp = new(big.Int).SetBit(new(big.Int).SetInt64(0), int(bits-1), 1)
	}

	if curveName == "SS512" {
		bits := 160
		curve.params = pbc.GenerateA(uint32(bits), 512)
		curve.maxrndexp = new(big.Int).SetBit(new(big.Int).SetInt64(0), int(bits-1), 1)
	}

	return JsonObj2Str(curve.Encode())
}
