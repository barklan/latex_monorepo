package main

import (
	"fmt"
	"log"
	"math/big"
	"math/bits"
	"os"
	"strconv"
	"strings"
)

func data(variant int) (map[float64]float64, [][]float64) {
	var tableNtn [][]float64
	switch variant {
	case 1:
		tableNtn = [][]float64{
			{1910.0, 92228496.0},
			{1920.0, 106021537.0},
			{1930.0, 123202624.0},
			{1940.0, 132164569.0},
			{1950.0, 151325798.0},
			{1960.0, 179323175.0},
			{1970.0, 203211926.0},
			{1980.0, 226545805.0},
			{1990.0, 248709873.0},
			{2000.0, 281421906.0},
		}
	case 2:
		tableNtn = [][]float64{
			{5, 296},
			{7, 520},
			{9, 744},
			{11, 982},
			{13, 1248},
			{15, 1570},
			{17, 2256},
			{19, 2256},
			{21, 2256},
		}
	case 5:
		tableNtn = [][]float64{
			{7, 83.7},
			{12, 72.9},
			{17, 63.2},
			{22, 54.7},
			{27, 47.5},
			{32, 41.4},
			{37, 36.3},
		}
	case 6:
		tableNtn = [][]float64{
			{0.0, 0.5},
			{1.0, 0.3},
			{3.0, 0.3},
			{4.0, 0.2},
			{5.0, 0.1},
		}
	case 7:
		tableNtn = [][]float64{
			{0.0, 5.0},
			{0.1, 2.5},
			{0.2, 3.0},
			{0.3, -2.5},
			{0.4, -0.2},
		}
	case 666:
		var content string
		f, err := os.ReadFile("data.txt")
		if err == nil {
			content = string(f)
			content = strings.TrimSpace(content)
		} else {
			content = os.Args[1]
		}

		tableNtn = make([][]float64, 0)
		strNums := strings.Split(content, ",")
		temp := make([]float64, 2)
		for i, str := range strNums {
			parsed, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Println("failed to parse float64")
				log.Fatal(err)
			}
			if i%2 == 0 {
				temp[0] = parsed
			} else {
				temp[1] = parsed
				tempCopy := make([]float64, len(temp))
				copy(tempCopy, temp)
				tableNtn = append(tableNtn, tempCopy)
			}
		}
		fmt.Println(tableNtn)

		// tableNtn = [][]float64{
		// 	{1.0, 5.0},
		// 	{2.0, 7.0},
		// 	{3.0, 8.0},
		// 	{4.0, 10.0},
		// 	{5.0, 11.0},
		// }
	}

	table := map[float64]float64{}
	for _, xy := range tableNtn {
		table[xy[0]] = xy[1]
	}

	return table, tableNtn
}

func main() {
	variants := []int{666}
	for _, variant := range variants {
		fmt.Printf("\n %d -----------\n", variant)
		table, tableNtn := data(variant)

		bigTable := floatToBigTable(table)
		lPolinome := Lagrange(bigTable)

		bigTableNtn := floatToBigTableN(tableNtn)
		nPolinome := NewTon(bigTableNtn)

		poli := map[string]Polinome{
			"Lagrange": lPolinome,
			"Newton":   nPolinome,
		}

		for name, polinome := range poli {
			fmt.Printf("\n%s's Polinome:\nP = %s", name, polinome)

			switch variant {
			case 1:
				xIlyana := big.NewRat(2010, 1)
				exact := polinome.Exact(xIlyana).FloatString(0)
				fmt.Printf("\nx = %v; result: %v", xIlyana, exact)
			case 5:
				x := big.NewRat(25, 1)
				fmt.Printf("\nx = %v; result: %v", x, polinome.Exact(x))
				fmt.Printf("\nx = %v; approx: %v", x, polinome.Exact(x).FloatString(5))
			case 6:
				derivative := polinome.Derivative()
				xVar6 := big.NewRat(5, 1)
				fmt.Printf("\nx = %v; result: %v", xVar6, derivative.Exact(xVar6))
			case 7:
				derivative := polinome.Derivative()
				secondDerivative := derivative.Derivative()
				x := big.NewRat(3, 10)
				fmt.Printf("\nx = %v; result: %v", x, secondDerivative.Exact(x))
				fmt.Printf("\nx = %v; approx: %v", x, secondDerivative.Exact(x).FloatString(5))
			case 666:
				fmt.Println()
				derivative := polinome.Derivative()
				fmt.Printf("First derivative: %v\n", derivative)
				secondDerivative := derivative.Derivative()
				fmt.Printf("Second derivative: %v\n", secondDerivative)
				// x := big.NewRat(3, 10)
				// fmt.Printf("\nx = %v; result: %v", x, secondDerivative.Exact(x))
				// fmt.Printf("\nx = %v; approx: %v", x, secondDerivative.Exact(x).FloatString(5))

			}
			fmt.Println()
		}
	}
}

func floatToBigTableN(table [][]float64) [][]*big.Rat {
	bigTable := make([][]*big.Rat, len(table))
	for i, xy := range table {
		newPoint := make([]*big.Rat, 2)

		newX := big.NewRat(int64(xy[0]*10), 10)
		newPoint[0] = newX

		newY := big.NewRat(int64(xy[1]*10), 10)
		newPoint[1] = newY

		bigTable[i] = newPoint
	}
	return bigTable
}

func NewTon(bigTable [][]*big.Rat) Polinome {
	degree := len(bigTable) - 1
	polinome := Polinome{
		Degree: degree,
		Coefs:  make([]*big.Rat, degree+1),
	}
	for i := 0; i <= degree; i++ {
		polinome.Coefs[i] = big.NewRat(0, 1)
	}

	for i := 1; i <= degree; i++ {
		roots := make([]*big.Rat, i)
		for rootIndex := range roots {
			roots[rootIndex] = bigTable[rootIndex][0]
		}

		a := DivDiff(bigTable, i)

		interPoli := Polinome{
			Degree: i,
			Roots:  roots,
			A:      a,
		}
		interPoli.Parse()

		polinome.Add(interPoli)
	}

	polinome.Coefs[len(polinome.Coefs)-1].Add(polinome.Coefs[len(polinome.Coefs)-1], bigTable[len(bigTable)-1][1])

	return polinome
}

func DivDiff(bigTable [][]*big.Rat, n int) *big.Rat {
	result := big.NewRat(0, 1)

	for j := 0; j <= n; j++ {
		inter := big.NewRat(0, 1)
		inter.Set(bigTable[j][1])

		denominator := big.NewRat(1, 1)
		for i := 0; i <= n; i++ {
			if i != j {
				devisionResult := big.NewRat(0, 1)
				devisionResult.Set(bigTable[j][0])
				devisionResult.Sub(devisionResult, bigTable[i][0])
				denominator.Mul(denominator, devisionResult)
			}
		}
		denominator.Inv(denominator)

		inter.Mul(inter, denominator)
		result.Add(result, inter)
	}

	fmt.Printf("\nDivdiff %d: %v", n, result)

	return result
}

func Lagrange(bigTable map[*big.Rat]*big.Rat) Polinome {
	lagrangeCoefMap := make(map[*big.Rat]*big.Rat)

	for x, y := range bigTable {
		lagrangeCoef := big.NewRat(1, 1)
		multi := big.NewRat(1, 1)
		w := wPolinome(bigTable, x)
		inverted := multi.Inv(w)
		lagrangeCoef.Mul(y, inverted)
		// fmt.Printf("x: %s\t coef: %s\n", x, lagrangeCoef)
		lagrangeCoefMap[x] = lagrangeCoef
	}

	degree := len(bigTable) - 1

	polinome := Polinome{
		Degree: degree,
		Coefs:  make([]*big.Rat, degree+1),
	}
	for i := 0; i < degree+1; i++ {
		polinome.Coefs[i] = big.NewRat(0, 1)
	}

	for x, coef := range lagrangeCoefMap {
		roots := make([]*big.Rat, 0)
		for xOrig := range bigTable {
			if x.Cmp(xOrig) != 0 {
				xCopy := big.NewRat(1, 1)
				xCopy.Set(xOrig)
				roots = append(roots, xCopy)
			}
		}
		interPoli := Polinome{
			Degree: degree,
			A:      coef,
			Roots:  roots,
		}
		interPoli.Parse()

		// fmt.Printf("roots: %v", interPoli.Roots)
		fmt.Printf("x: %v; Intermediate polinome: %v\n", x, interPoli)

		polinome.Add(interPoli)

	}

	return polinome
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

func (p *Polinome) Exact(x *big.Rat) *big.Rat {
	result := big.NewRat(0, 1)
	for i, coef := range p.Coefs {
		if i != len(p.Coefs)-1 {
			power := p.Degree - i
			powered := big.NewRat(1, 1)
			powered.Set(x)
			for j := 0; j < power-1; j++ {
				powered.Mul(powered, x)
			}
			inter := big.NewRat(0, 1)
			inter.Mul(coef, powered)
			result.Add(result, inter)
		}
	}
	result.Add(result, p.Coefs[len(p.Coefs)-1])
	return result
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
	if p.Degree < other.Degree {
		log.Panic("no")
	} else if p.Degree > other.Degree {
		zerosNumber := p.Degree - other.Degree
		newSlice := make([]*big.Rat, zerosNumber)
		for i := range newSlice {
			newSlice[i] = big.NewRat(0, 1)
		}

		newSlice = append(newSlice, other.Coefs...)
		other.Coefs = newSlice
		other.Degree = p.Degree
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
