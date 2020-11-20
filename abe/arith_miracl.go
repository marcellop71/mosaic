// +build miracl

package abe

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	core "github.com/marcellop71/mosaic/abe/miracl/core"
	crv "github.com/marcellop71/mosaic/abe/miracl/core/BN254"
	//crv "github.com/marcellop71/mosaic/abe/miracl/core/BLS12381"
	//crv "github.com/marcellop71/mosaic/abe/miracl/core/BN462"
	"github.com/marcellop71/mosaic/abe/log"
)

func NewPoint() Point {
	return new(MiraclPoint)
}

func NewCurve() Curve {
	return new(MiraclCurve)
}

type MiraclPoint struct {
	p1    *crv.ECP
	p2    *crv.ECP2
	pt    *crv.FP12
	P     string `json:"p"`
	Group string `json:"group"`
}

func (p *MiraclPoint) isPoint() {}

func (p *MiraclPoint) GetP() string {
	return p.P
}

func (p *MiraclPoint) GetGroup() string {
	return p.Group
}

// MiraclPoint type to its json representation
func (p *MiraclPoint) ToJsonObj() Point {
	buf := make([]byte, 12*crv.MODBYTES)
	switch p.Group {
	case "G1":
		p.p1.ToBytes(buf, false)
	case "G2":
		p.p2.ToBytes(buf, false)
	case "GT":
		p.pt.ToBytes(buf)
	}
	p.P = Encode(string(buf))
	return p
}

// PbcPoint type from its json representation
func (p *MiraclPoint) OfJsonObj(curve Curve) Point {
	switch p.Group {
	case "G1":
		p.p1 = crv.ECP_fromBytes([]byte(Decode(p.P)))
	case "G2":
		p.p2 = crv.ECP2_fromBytes([]byte(Decode(p.P)))
	case "GT":
		p.pt = crv.FP12_fromBytes([]byte(Decode(p.P)))
	}
	return p
}

// curve on miracl library
type MiraclCurve struct {
	rng *core.RAND
	Name string `json:"name"`
	Seed string `json:"seed"`
}

func (curve *MiraclCurve) isCurve() {}

func (curve *MiraclCurve) ToJsonObj() Curve {
	return curve
}

func (curve *MiraclCurve) OfJsonObj() Curve {
	curve.InitRng()
	return curve
}

func (curve *MiraclCurve) SetSeed(seed string) Curve {
	curve.Seed = seed
	return curve
}

func (curve *MiraclCurve) InitRng() Curve {
	curve.rng = core.NewRAND()
	hseed := sha256.Sum256([]byte(curve.Seed))
	curve.rng.Seed(32, hseed[:])
	return curve
}

// new point instance on a curve
func (curve *MiraclCurve) NewPointOn(group string) Point {
	p := new(MiraclPoint)
	switch group {
	case "G1":
		p.p1 = crv.NewECP()
	case "G2":
		p.p2 = crv.NewECP2()
	case "GT":
		p.pt = crv.NewFP12()
	}
	p.Group = group
	return p
}

// new random exponent
func (curve *MiraclCurve) NewRandomExp() *big.Int {
	order_ := crv.NewBIGints(crv.CURVE_Order)
	order := crv.BIGtoBig(order_)
	rng := rand.Reader
	x, _ := rand.Int(rng, order)

	//m := crv.NewBIGints(crv.Modulus)
	//xtmp := crv.Randomnum(m, curve.rng)
	//x := crv.BIGtoBig(xtmp)
	return x
}

// new random point on a group ("G1", "G2", "GT") in the bilinear pairing
func (curve *MiraclCurve) NewRandomPointOn(group string) Point {
	p := new(MiraclPoint)

	switch group {
	case "G1":
		x := curve.NewRandomExp()
		x_ := crv.BigToBIG(x)
		q1 := crv.ECP_generator()
		p.p1 = crv.G1mul(q1, x_)
	case "G2":
		x := curve.NewRandomExp()
		x_ := crv.BigToBIG(x)
		q2 := crv.ECP2_generator()
		p.p2 = crv.G2mul(q2, x_)
	case "GT":
		p.pt = crv.NewFP12rand(curve.rng)
		buf := make([]byte, 12*crv.MODBYTES)
		p.pt.ToBytes(buf)
		log.Info("%s", Encode(string(buf)))
	}
	p.Group = group
	return p
}

// unit on group ("G1", "G2", "GT") in the bilinear pairing
func (curve *MiraclCurve) UnitOnGroup(group string) Point {
	p := curve.NewPointOn(group).(*MiraclPoint)
	switch group {
	case "G1":
	case "G2":
	case "GT":
		p.pt.One()
	}
	return p
}

// hash to group ("G1", "G2", "GT") in the bilinear pairing
func (curve *MiraclCurve) HashToGroup(x string, group string) Point {
	hx := sha256.Sum256([]byte(x))
	p := curve.NewPointOn(group).(*MiraclPoint)
	switch group {
	case "G1":
		p.p1 = crv.Hash_to_point_G1(hx[:])
	case "G2":
		p.p2 = crv.Hash_to_point_G2(hx[:])
	case "GT":
	}
	return p
}

// hash to group ("G1", "G2", "GT") in the bilinear pairing
func (curve *MiraclCurve) HashToPow(x string, g Point) Point {
	hx := sha256.Sum256([]byte(x))
	hx_ := new(big.Int).SetBytes(hx[:])
	hx__ := crv.BigToBIG(hx_)
	p := curve.NewPointOn(g.GetGroup()).(*MiraclPoint)
	switch g.GetGroup() {
	case "G1":
		p.p1 = crv.G1mul(g.(*MiraclPoint).p1, hx__)
	case "G2":
		p.p2 = crv.G2mul(g.(*MiraclPoint).p2, hx__)
	case "GT":
	}
	return p
}

// group operation
func (curve *MiraclCurve) Inv(g1 Point) Point {
	p := curve.NewPointOn(g1.GetGroup()).(*MiraclPoint)
	switch g1.GetGroup() {
	case "G1":
		p.p1.Copy(g1.(*MiraclPoint).p1)
		p.p1.Neg()
	case "G2":
	case "GT":
		p.pt.Copy(g1.(*MiraclPoint).pt)
		p.pt.Inverse()
	}
	return p
}

// group operation
func (curve *MiraclCurve) Mul(g1 Point, g2 Point) Point {
	p := curve.NewPointOn(g1.GetGroup()).(*MiraclPoint)
	switch g1.GetGroup() {
	case "G1":
		p.p1.Copy(g1.(*MiraclPoint).p1)
		p.p1.Add(g2.(*MiraclPoint).p1)
	case "G2":
		p.p2.Copy(g1.(*MiraclPoint).p2)
		p.p2.Add(g2.(*MiraclPoint).p2)
	case "GT":
		p.pt.Copy(g1.(*MiraclPoint).pt)
		p.pt.Mul(g2.(*MiraclPoint).pt)
	}
	return p
}

// group operation
func (curve *MiraclCurve) Div(g1 Point, g2 Point) Point {
	p := curve.NewPointOn(g1.GetGroup()).(*MiraclPoint)
	switch g1.GetGroup() {
	case "G1":
		p.p1.Copy(g1.(*MiraclPoint).p1)
		p.p1.Sub(g2.(*MiraclPoint).p1)
	case "G2":
		p.p2.Copy(g1.(*MiraclPoint).p2)
		p.p2.Sub(g2.(*MiraclPoint).p2)
	case "GT":
		p.pt.Copy(g2.(*MiraclPoint).pt)
		p.pt.Inverse()
		p.pt.Mul(g1.(*MiraclPoint).pt)
	}
	return p
}

// group operation
func (curve *MiraclCurve) Pow(g1 Point, e1 *big.Int) Point {
	p := curve.NewPointOn(g1.GetGroup()).(*MiraclPoint)
	e1_ := new(big.Int).Abs(e1)
	E1 := crv.BigToBIG(e1_)
	switch g1.GetGroup() {
	case "G1":
		p.p1 = crv.G1mul(g1.(*MiraclPoint).p1, E1)
	case "G2":
		p.p2 = crv.G2mul(g1.(*MiraclPoint).p2, E1)
	case "GT":
		p.pt = crv.GTpow(g1.(*MiraclPoint).pt, E1)
	}
	if e1.Sign() < 0 {
		switch g1.GetGroup() {
		case "G1":
			p.p1.Neg()
		case "G2":
		case "GT":
			p.pt.Inverse()
		}
	}
	return p
}

// pairing
func (curve *MiraclCurve) Pair(g1 Point, g2 Point) Point {
	p := curve.NewPointOn("GT").(*MiraclPoint)
	p.pt = crv.Ate(g2.(*MiraclPoint).p2, g1.(*MiraclPoint).p1)
	p.pt = crv.Fexp(p.pt)
	return p
}

// group operation
func (curve *MiraclCurve) Pow2(g1 Point, e1 *big.Int, g2 Point, e2 *big.Int) Point {
	return curve.Mul(curve.Pow(g1, e1), curve.Pow(g2, e2))
}

// group operation
func (curve *MiraclCurve) Pow3(g1 Point, e1 *big.Int, g2 Point, e2 *big.Int, g3 Point, e3 *big.Int) Point {
	return curve.Mul(curve.Pow(g1, e1), curve.Mul(curve.Pow(g2, e2), curve.Pow(g3, e3)))
}

// pairing
func (curve *MiraclCurve) ProdPair(g1 []Point, g2 []Point) Point {
	tmp := curve.UnitOnGroup("GT")
	for i := 0; i < len(g1); i++ {
		tmp = curve.Mul(tmp, curve.Pair(g1[i], g2[i]))
	}
	return tmp
}

// new random secret vector of length n
// if zerosecret is true, the first component of the vector is zero
func (curve *MiraclCurve) NewRandomSecret(n int, zerosecret bool) []*big.Int {
	s := make([]*big.Int, n)
	for i := 0; i < n; i++ {
		s[i] = curve.NewRandomExp()
	}
	if zerosecret {
		s[0] = new(big.Int).SetInt64(0)
	}
	return s
}
