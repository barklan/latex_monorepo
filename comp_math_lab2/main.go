package main

import (
	"fmt"
	"log"
	"math/big"
	"math/bits"
)

func main() {
	table := map[float64]float64{
		0.0: 5.0,
		0.1: 2.5,
		0.2: 3.0,
		0.3: -2.5,
		0.4: -0.2,
	}

	bigTable := floatToBigTable(table)
	lagrangeCoefMap := make(map[*big.Rat]*big.Rat)

	for x, y := range bigTable {
		lagrangeCoef := big.NewRat(1, 1)
		multi := big.NewRat(1, 1)
		w := wPolinome(bigTable, x)
		lagrangeCoef.Mul(y, multi.Inv(w))
		fmt.Printf("x: %s\t coef: %s\n", x, lagrangeCoef)
		lagrangeCoefMap[x] = lagrangeCoef
	}

	degree := len(bigTable) - 1

	polinome := Polinome{
		Degree: degree,
		Coefs: []*big.Rat{
			big.NewRat(0, 1),
			big.NewRat(0, 1),
			big.NewRat(0, 1),
			big.NewRat(0, 1),
			big.NewRat(0, 1),
		},
	}

	for x, coef := range lagrangeCoefMap {
		roots := make([]*big.Rat, 0)
		for xOrig := range bigTable {
			if x.Cmp(xOrig) != 0 {
				xCopy := big.NewRat(1, 1)
				xCopy.Set(x)
				roots = append(roots, xCopy)
			}
		}
		interPoli := Polinome{
			Degree: degree,
			A:      coef,
			Roots:  roots,
		}
		interPoli.Parse()

		// fmt.Printf("Intermediate polinome: %v\n", interPoli)

		polinome.Add(interPoli)
	}

	fmt.Printf("Polinome:\nP = %s", polinome)

	derivative := polinome.Derivative()
	fmt.Printf("\nP' = %s", derivative)

	secondDerivative := derivative.Derivative()
	fmt.Printf("\nP'' = %s", secondDerivative)
}

func wPolinome(table map[*big.Rat]*big.Rat, xTarget *big.Rat) *big.Rat {
	w := big.NewRat(1, 1)
	for x := range table {
		if x.Cmp(xTarget) != 0 {
			xTargetCopy := big.NewRat(1, 1)
			w.Mul(w, xTargetCopy.Sub(xTarget, x))
		}
	}
	return w
}

func floatToBigTable(table map[float64]float64) map[*big.Rat]*big.Rat {
	bigTable := make(map[*big.Rat]*big.Rat)
	for k, v := range table {
		newK := big.NewRat(int64(k*10), 10)
		newV := big.NewRat(int64(v*10), 10)
		bigTable[newK] = newV
	}
	return bigTable
}

type Polinome struct {
	Degree int
	Coefs  []*big.Rat
	Roots  []*big.Rat
	A      *big.Rat
}

func (p Polinome) String() string {
	var s string
	for i := p.Degree; i >= 0; i-- {
		s += fmt.Sprintf("%v x^%d", p.Coefs[p.Degree-i], i)
		if i != 0 {
			s += fmt.Sprintf(" + ")
		}
	}
	return s
}

func (p *Polinome) Derivative() Polinome {
	derivative := Polinome{
		Degree: p.Degree - 1,
		Coefs:  make([]*big.Rat, len(p.Coefs)-1),
	}

	for i := 0; i < len(p.Coefs)-1; i++ {
		new := big.NewRat(1, 1)
		multiplier := big.NewRat(int64(p.Degree-i), 1)
		derivative.Coefs[i] = new.Mul(p.Coefs[i], multiplier)
	}

	return derivative
}

func (p *Polinome) Add(other Polinome) {
	if p.Degree != other.Degree {
		log.Panic("no")
	}
	for i := range p.Coefs {
		p.Coefs[i].Add(p.Coefs[i], other.Coefs[i])
	}
}

func (p *Polinome) Parse() {
	p.Coefs = make([]*big.Rat, p.Degree+1)
	p.Coefs[0] = big.NewRat(1, 1)

	for i := 1; i <= p.Degree; i++ {
		combinations := Combinations(p.Roots, i)
		coef := MultiSumCombinations(combinations)
		p.Coefs[i] = coef
	}

	for i := range p.Coefs {
		if i%2 != 0 {
			p.Coefs[i].Neg(p.Coefs[i])
		}
		p.Coefs[i].Mul(p.Coefs[i], p.A)
	}
}

func Combinations(set []*big.Rat, n int) (subsets [][]*big.Rat) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []*big.Rat

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

func MultiSumCombinations(set [][]*big.Rat) *big.Rat {
	result := big.NewRat(0, 1)
	for _, subset := range set {
		multi := big.NewRat(1, 1)
		for _, bigRat := range subset {
			multi.Mul(multi, bigRat)
		}
		result.Add(result, multi)
	}
	return result
}
