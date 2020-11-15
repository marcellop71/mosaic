package abe

import (
	"math"
	"math/big"

	"github.com/marcellop71/mosaic/abe/log"
)

// ap linear access policy encoded into a full matrix by row
func btreeLabeling(t *btree, c *int) {
	if *c == 0 {
		*c = 1
		t.v = []int{1}
	}

	internal := (t.child[0] != nil) && (t.child[1] != nil)
	if internal == true {
		if t.label == "or" {
			padWithZeros(&t.v, *c)
			t.child[0].v = t.v
			t.child[1].v = t.v
			btreeLabeling(t.child[0], c)
			btreeLabeling(t.child[1], c)
		}
		if t.label == "and" {
			padWithZeros(&t.v, *c)
			zeros := make([]int, len(t.v))
			t.child[0].v = append(t.v, 1)
			t.child[1].v = append(zeros, -1)
			*c = *c + 1
			btreeLabeling(t.child[0], c)
			btreeLabeling(t.child[1], c)
		}
	}
}

func btreeExtractLabelsOnLeaves(t *btree, ap *[][]int) {
	internal := (t.child[0] != nil) && (t.child[1] != nil)
	if internal == true {
		log.Debug(">>> %s %d %d", t.label, t.child[0].v, t.child[1].v)
		btreeExtractLabelsOnLeaves(t.child[0], ap)
		btreeExtractLabelsOnLeaves(t.child[1], ap)
	} else {
		*ap = append(*ap, t.v)
	}
}

func encodeAccessPolicy(t *btree) [][]int {
	var c int
	var aptmp [][]int
	btreeLabeling(t, &c)
	btreeExtractLabelsOnLeaves(t, &aptmp)
	ap := make([][]int, len(aptmp))
	for i := 0; i < len(aptmp); i++ {
		ap[i] = make([]int, c)
		for j := 0; j < len(aptmp[i]); j++ {
			ap[i][j] = aptmp[i][j]
		}
	}
	return ap
}

func computeShares(s []*big.Int, ap [][]int) []*big.Int {
	sh := make([]*big.Int, len(ap))
	for i := 0; i < len(ap); i++ {
		sh[i] = big.NewInt(0)
		for j := 0; j < len(ap[i]); j++ {
			if ap[i][j] == 1 {
				sh[i].Add(sh[i], s[j])
			}
			if ap[i][j] == -1 {
				sh[i].Sub(sh[i], s[j])
			}
		}
	}
	return sh
}

func computeCoefficients(ap [][]int) []int {
	hap, u := hermiteNormalForm(ap)
	log.Debug("HNF %d", hap)
	log.Debug("U %d", u)

	m := len(ap[0])
	n := len(hap)
	c := make([]int, n)

	y := make([]int, n)
	for i := 0; i < n; i++ {
		y[i] = 0
	}
	if m >= n {
		y[0] = 1
	} else {
		y[n-m] = 1
	}

	for i := 0; i < n; i++ {
		c[i] = 0
		for j := 0; j < n; j++ {
			c[i] += u[j][i] * y[j]
		}
	}

	return c
}

func padWithZeros(x *[]int, tlen int) {
	for i := 0; i < tlen-len(*x); i++ {
		*x = append(*x, 0)
	}
}

func swapCol(a *[][]int, j0 int, j1 int) {
	for i := 0; i < len((*a)[0]); i++ {
		tmp0 := (*a)[j0][i]
		tmp1 := (*a)[j1][i]
		(*a)[j0][i] = tmp1
		(*a)[j1][i] = tmp0
	}
}

func negCol(a *[][]int, j0 int) {
	for i := 0; i < len((*a)[0]); i++ {
		(*a)[j0][i] = -(*a)[j0][i]
	}
}

func lcCol(a *[][]int, j0 int, j1 int, b int) {
	for i := 0; i < len((*a)[0]); i++ {
		(*a)[j0][i] = (*a)[j0][i] - (b * (*a)[j1][i])
	}
}

func clone(a [][]int) [][]int {
	a_ := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		a_[i] = make([]int, len(a[0]))
		copy(a_[i], a[i])
	}
	return a_
}

func identity(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			a[i][j] = 0
		}
		a[i][i] = 1
	}
	return a
}

func row(a [][]int, i int) []int {
	x := make([]int, len(a))
	for j := 0; j < len(a); j++ {
		x[j] = a[j][i]
	}
	return x
}

func isZeroUpTo(x []int, n int) (int, int) {
	for i := 0; i < n; i++ {
		if x[i] != 0 {
			return x[i], i
		}
	}
	return 0, 0
}

func nearestint(a int, b int) int {
	z := float64(a) / float64(b)
	return int(math.Floor(z + 0.50))
}

func floor(a int, b int) int {
	z := float64(a) / float64(b)
	return int(math.Floor(z))
}

// matrix a m x n by columns
func hermiteNormalForm(a [][]int) ([][]int, [][]int) {
	a_ := clone(a)

	m := len(a_[0])
	n := len(a_)
	q := 0
	if m > n {
		q = m - n
	}

	u := identity(n)

	for i, k := m-1, n-1; i >= q; i-- {
		w, j0 := isZeroUpTo(row(a_, i), k)
		if w != 0 {
			swapCol(&a_, j0, k)
			swapCol(&u, j0, k)
			if row(a_, i)[k] < 0 {
				negCol(&a_, k)
				negCol(&u, k)
			}
			for j := 0; j < k; j++ {
				tmp := nearestint(a_[j][i], a_[k][i])
				lcCol(&a_, j, k, tmp)
				lcCol(&u, j, k, tmp)
			}
		} else {
			if row(a_, i)[k] < 0 {
				negCol(&a_, k)
				negCol(&u, k)
			}
		}

		if a_[k][i] == 0 {
			k += 1
		} else {
			for j := k + 1; j < n; j++ {
				tmp := floor(a_[j][i], a_[k][i])
				lcCol(&a_, j, k, tmp)
				lcCol(&u, j, k, tmp)
			}
		}

		k--
	}

	return a_, u
}
