package main

import (
	"fmt"
	"log"
	"math/big"
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

		fmt.Printf("Intermediate polinome: %v\n", interPoli)

		polinome.Add(interPoli)
	}

	fmt.Printf("Polinome: %v", polinome)
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

func (p *Polinome) Add(other Polinome) {
	if p.Degree != other.Degree {
		log.Panic("no")
	}
	for i := range p.Coefs {
		p.Coefs[i].Add(p.Coefs[i], other.Coefs[i])
	}
}

func (p *Polinome) Parse() {
	first := big.NewRat(1, 1)
	first.Set(p.A)
	p.Coefs = append(p.Coefs, first)

	second := big.NewRat(0, 1)
	for _, root := range p.Roots {
		second.Add(second, root)
	}
	second.Mul(second, p.A)
	second.Neg(second)
	p.Coefs = append(p.Coefs, second)

	third := big.NewRat(0, 1)
	for i := range p.Roots {
		for j := i + 1; j < len(p.Roots); j++ {
			inter := big.NewRat(1, 1)
			inter.Mul(p.Roots[i], p.Roots[j])
			third.Add(third, inter)
		}
	}
	third.Mul(third, p.A)
	p.Coefs = append(p.Coefs, third)
}
