package abe

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	//pbc "github.com/Nik-U/pbc"
	//core "github.com/miracl/core/go/core"
  //BN254 "github.com/miracl/core/go/core/BN254"

	log "github.com/unicredit/abe/log"
	pbc "github.com/unicredit/abe/pbc"
)

func (p *AbePoint) Encode() (*AbePoint_) {
  log.Debug("encoding point")
	return &AbePoint_{
		encode(p.p.Bytes()), p.group}
}

func (p_ *AbePoint_) Decode(curve *AbeCurve) (*AbePoint) {
  log.Debug("decoding point")
	return &AbePoint{
    curve.NewPointOn(p_.Group).p.SetBytes(decode(p_.P)),
    p_.Group}
}

func (curve *AbeCurve) Encode() (*AbeCurve_) {
  log.Debug("encoding curve")
  return &AbeCurve_{
    curve.params.String(),
    encode(curve.maxrndexp.Bytes())}
}

func (curve_ *AbeCurve_) Decode() (*AbeCurve) {
  log.Debug("decoding curve")
  curve := new(AbeCurve)
	curve.params = pbc.NewParamsFromString(curve_.Params)
  curve.pairing = pbc.NewPairing(curve.params)
  curve.maxrndexp = new(big.Int).SetBytes(decode(curve_.Maxrndexp))
  return curve
}

func (curve *AbeCurve) NewPointOn(group string) (*AbePoint) {
	return &AbePoint{
		curve.pairing.NewUncheckedElement(group), group}
}

func (curve *AbeCurve) NewRandomExp() *big.Int {
	rng := rand.Reader
	x, _ := rand.Int(rng, curve.maxrndexp)
	return x
}

func (curve *AbeCurve) NewRandomSecret(n int, zerosecret bool) []*big.Int {
	s := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		s[i] = curve.NewRandomExp()
	}
	if (zerosecret) {
		s[0] = new(big.Int).SetInt64(0)
	}
	return s
}

func (curve *AbeCurve) NewRandomPointOn(group string) *AbePoint {
	return &AbePoint{curve.NewPointOn(group).p.Rand(), group}
}

func (curve *AbeCurve) UnitOnGroup(group string) *AbePoint {
	return &AbePoint{curve.NewPointOn(group).p.Set1(), group}
}

func (curve *AbeCurve) HashToGroup(x string, group string) *AbePoint {
	hx := sha256.Sum256([]byte(x))
	return &AbePoint{
		curve.NewPointOn(group).p.SetFromHash(hx[:]), group}
}

func (curve *AbeCurve) HashToPow(x string, g *AbePoint) *AbePoint {
	hx := sha256.Sum256([]byte(x))
	return curve.Pow(g, new(big.Int).SetBytes(hx[:]))
}

func (curve *AbeCurve) Inv(g1 *AbePoint) *AbePoint {
	return &AbePoint{curve.NewPointOn(g1.group).p.Invert(g1.p), g1.group}
}

func (curve *AbeCurve) Mul(g1 *AbePoint, g2 *AbePoint) *AbePoint {
	return &AbePoint{curve.NewPointOn(g1.group).p.Mul(g1.p, g2.p), g1.group}
}

func (curve *AbeCurve) Div(g1 *AbePoint, g2 *AbePoint) *AbePoint {
	return &AbePoint{curve.NewPointOn(g1.group).p.Div(g1.p, g2.p), g1.group}
}

func (curve *AbeCurve) Pow(g1 *AbePoint, e1 *big.Int) *AbePoint {
  tmp := curve.NewPointOn(g1.group).p.PowBig(g1.p, new(big.Int).Abs(e1))
  if (e1.Sign() < 0) {
    tmp.ThenInvert()
  }
	return &AbePoint{tmp, g1.group}
}

func (curve *AbeCurve) Pow2(g1 *AbePoint, e1 *big.Int, g2 *AbePoint, e2 *big.Int) *AbePoint {
  return curve.Mul(curve.Pow(g1, e1), curve.Pow(g2, e2))
}

func (curve *AbeCurve) Pow3(g1 *AbePoint, e1 *big.Int, g2 *AbePoint, e2 *big.Int, g3 *AbePoint, e3 *big.Int) *AbePoint {
  return curve.Mul(curve.Pow(g1, e1), curve.Mul(curve.Pow(g2, e2), curve.Pow(g3, e3)))
}

func (curve *AbeCurve) Pair(g1 *AbePoint, g2 *AbePoint) *AbePoint {
  group := "GT"
	return &AbePoint{curve.NewPointOn(group).p.Pair(g1.p, g2.p), group}
}

func (curve *AbeCurve) ProdPair(g1 []*AbePoint, g2 []*AbePoint) *AbePoint {
  tmp := curve.UnitOnGroup("GT")
  for i := 0; i < len(g1); i++ {
    tmp = curve.Mul(tmp, curve.Pair(g1[i], g2[i]))
  }
  return tmp
}
