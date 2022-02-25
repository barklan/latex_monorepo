package main

import (
	"fmt"
	"math"
	nm "num/numerical"
)

func fn(x float64) float64 {
	return 3*x + 4*math.Pow(x, 3) - 12*math.Pow(x, 2) - 5
}

// Re-writing f(x)=0 to x = g(x)
func fnG(x float64) float64 {
	intermediate := 3*math.Pow(x, 2) - 0.75*x + 1.25
	return math.Cbrt(intermediate)
}

func fnPrime(x float64) float64 {
	return 3 + 12*math.Pow(x, 2) - 24*x
}

func fn2(x float64) float64 {
	return 2*math.Tan(x) - x/2 + 1
}

// todo
func fn2G(x float64) float64 {
	return 0.0
}

func fn2Prime(x float64) float64 {
	cos := math.Cos(x)
	return -1/2 + 2/math.Pow(cos, 2)
}

func main() {
	fmt.Println("First equation: 3x + 4x^3 - 12x^2 - 5 = 0")
	fmt.Println("\nBisection:")
	fmt.Println(nm.Bisection(fn, 2.5, 3.5, 0.0001, 1000000000))
	fmt.Println("\nSecant:")
	fmt.Println(nm.Secant(fn, 2.5, 3.5, 0.0001, 1000))
	fmt.Println("\nFixed point iteration:")
	nm.FixedPointIteration(fn, fnG, 2.5, 0.0001, 100)
	fmt.Println("\nNewton's method:")
	nm.NewtonsMethod(fn, fnPrime, 2.5, 0.0001, 0.000000000001, 20)

	fmt.Println()
	fmt.Println("Second equation: 2tg(x) - x/2 + 1 = 0")
	fmt.Println("\nBisection:")
	fmt.Println(nm.Bisection(fn2, -1.0, 1.0, 0.0001, 1000000000))
	fmt.Println("\nSecant:")
	fmt.Println(nm.Secant(fn2, -1.0, 1.0, 0.0001, 100))
	// fmt.Println("\nFixed point iteration:")
	// nm.FixedPointIteration(fn, fnG, 2.5, 0.0001, 100)
	fmt.Println("\nNewton's method:")
	nm.NewtonsMethod(fn2, fn2Prime, -1.0, 0.0001, 0.000000000001, 20)
}
