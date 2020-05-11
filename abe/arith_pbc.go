// +build pbc

package abe

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"strconv"

	"github.com/unicredit/mosaic/abe/pbc"
)

func NewPoint() Point {
	return new(PbcPoint)
}

func NewCurve() Curve {
	return new(PbcCurve)
}

// point on pbc library
type PbcPoint struct {
	p *pbc.Element
	P string `json:"p"`
	Group string `json:"group"`
}

func (p *PbcPoint) isPoint() {}

func (p *PbcPoint) GetP() string {
	return p.P
}

func (p *PbcPoint) GetGroup() string {
	return p.Group
}

// PbcPoint type to its json representation
func (p *PbcPoint) ToJsonObj() Point {
	p.P = Encode(string(p.p.Bytes()))
	return p
}

// PbcPoint type from its json representation
func (p *PbcPoint) OfJsonObj(curve Curve) Point {
	tmp := curve.NewPointOn(p.Group).(*PbcPoint)
	p.p = tmp.p.SetBytes([]byte(Decode(p.P)))
	return p
}

// curve on pbc library
type PbcCurve struct {
	pairing *pbc.Pairing
	params *pbc.Params
	maxrndexp *big.Int
	Params string `json:"params"`
	Maxrndexp string `json:"maxrndexp"`
}

func (curve *PbcCurve) isCurve() {}

// PbcCurve type to its json representation
func (curve *PbcCurve) ToJsonObj() Curve {
	curve.Params = Encode(curve.params.String())
	curve.Maxrndexp = Encode(curve.maxrndexp.Bytes())
	return curve
}

// PbcCurve type from its json representation
func (curve *PbcCurve) OfJsonObj() Curve {
	curve.params = pbc.NewParamsFromString(Decode(curve.Params))
  curve.pairing = pbc.NewPairing(Decode(curve.params))
  curve.maxrndexp = new(big.Int).SetBytes(Decode(curve.Maxrndexp))
	return curve
}

// new point instance on a curve
func (curve *PbcCurve) NewPointOn(group string) Point {
	return &PbcPoint{
		p: curve.pairing.NewUncheckedElement(group),
	}
}

// new random exponent
func (curve *PbcCurve) NewRandomExp() *big.Int {
	rng := rand.Reader
	x, _ := rand.Int(rng, curve.maxrndexp)
	return x
}

// new random point on a group ("G1", "G2", "GT") in the bilinear pairing
func (curve *PbcCurve) NewRandomPointOn(group string) Point {
	tmp := curve.NewPointOn(group).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.Rand(),
		Group: group,
	}
}

// unit on group ("G1", "G2", "GT") in the bilinear pairing
func (curve *PbcCurve) UnitOnGroup(group string) Point {
	tmp := curve.NewPointOn(group).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.Set1(),
		Group: group,
	}
}

// hash to group ("G1", "G2", "GT") in the bilinear pairing
func (curve *PbcCurve) HashToGroup(x string, group string) Point {
	hx := sha256.Sum256([]byte(x))
	tmp := curve.NewPointOn(group).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.SetFromHash(hx[:]),
		Group: group,
	}
}

// hash to group ("G1", "G2", "GT") in the bilinear pairing
func (curve *PbcCurve) HashToPow(x string, g Point) Point {
	hx := sha256.Sum256([]byte(x))
	return curve.Pow(g, new(big.Int).SetBytes(hx[:]))
}

// group operation
func (curve *PbcCurve) Inv(g1 Point) Point {
	tmp := curve.NewPointOn(g1.GetGroup()).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.Invert(g1.(*PbcPoint).p),
		Group: g1.GetGroup(),
	}
}

// group operation
func (curve *PbcCurve) Mul(g1 Point, g2 Point) Point {
	tmp := curve.NewPointOn(g1.GetGroup()).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.Mul(g1.(*PbcPoint).p, g2.(*PbcPoint).p),
		Group: g1.GetGroup(),
	}
}

// group operation
func (curve *PbcCurve) Div(g1 Point, g2 Point) Point {
	tmp := curve.NewPointOn(g1.GetGroup()).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.Div(g1.(*PbcPoint).p, g2.(*PbcPoint).p),
		Group: g1.GetGroup(),
	}
}

// group operation
func (curve *PbcCurve) Pow(g1 Point, e1 *big.Int) Point {
	tmp := curve.NewPointOn(g1.GetGroup()).(*PbcPoint)
  tmp.p = tmp.p.PowBig(g1.(*PbcPoint).p, new(big.Int).Abs(e1))
  if (e1.Sign() < 0) {
    tmp.p.ThenInvert()
  }
	return &PbcPoint{
		p: tmp.p,
		Group: g1.GetGroup(),
	}
}

// pairing
func (curve *PbcCurve) Pair(g1 Point, g2 Point) Point {
	group := "GT"
	tmp := curve.NewPointOn(group).(*PbcPoint)
	return &PbcPoint{
		p: tmp.p.Pair(g1.(*PbcPoint).p, g2.(*PbcPoint).p),
		Group: group,
	}
}

// group operation
func (curve *PbcCurve) Pow2(g1 Point, e1 *big.Int, g2 Point, e2 *big.Int) Point {
  return curve.Mul(curve.Pow(g1, e1), curve.Pow(g2, e2))
}

// group operation
func (curve *PbcCurve) Pow3(g1 Point, e1 *big.Int, g2 Point, e2 *big.Int, g3 Point, e3 *big.Int) Point {
  return curve.Mul(curve.Pow(g1, e1), curve.Mul(curve.Pow(g2, e2), curve.Pow(g3, e3)))
}

// pairing
func (curve *PbcCurve) ProdPair(g1 []Point, g2 []Point) Point {
  tmp := curve.UnitOnGroup("GT")
  for i := 0; i < len(g1); i++ {
    tmp = curve.Mul(tmp, curve.Pair(g1[i], g2[i]))
  }
  return tmp
}

// new random secret vector of length n
// if zerosecret is true, the first component of the vector is zero
func (curve *PbcCurve) NewRandomSecret(n int, zerosecret bool) []*big.Int {
	s := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		s[i] = curve.NewRandomExp()
	}
	if (zerosecret) {
		s[0] = new(big.Int).SetInt64(0)
	}
	return s
}

func GetPbcCurve(curveStr string) string {
	curve := new(PbcCurve)

	if curveStr == "SS512" {
		bits := 160
		curve.params = pbc.GenerateA(uint32(bits), 512)
		curve.maxrndexp = new(big.Int).SetBit(new(big.Int).SetInt64(0), int(bits-1), 1)
	}

	if curveStr[0:2] == "BN" {
		bits, _ := strconv.ParseInt(string(curveStr[2:]), 10, 32)
		curve.params = pbc.GenerateF(uint32(bits))
		curve.maxrndexp = new(big.Int).SetBit(new(big.Int).SetInt64(0), int(bits-1), 1)
	}

	curve.Encode()
	return JsonObjToStr(curve)
}
