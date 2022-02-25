package numerical

import (
	"fmt"
	"math"
)

// x0 - the initial guess
// fn - the function whose root we are trying to find
// fnPrime - the derivative of the function
// tolerance - tolerance
// epsilon - do not divide by a number smaller than this
// todo make it a library function
func NewtonsMethod(fn, fnPrime func(float64) float64,
	x0, tolerance, epsilon float64, maxIterations int) {
	solutionFound := false
	var y, yPrime, x1 float64

	for i := 1; i <= maxIterations; i++ {
		fmt.Printf("Iteration-%d, x1 = %0.6f and fn(x1) = %0.6f\n",
			i, x1, fn(x1))
		y = fn(x0)
		yPrime = fnPrime(x0)

		// Stop if the denominator is too small
		if math.Abs(yPrime) < epsilon {
			break
		}

		// Do Newton's computation
		x1 = x0 - y/yPrime

		// Stop when the result is within the desired tolerance
		if math.Abs(x1-x0) <= tolerance {
			solutionFound = true
			break
		}

		// Update x0 to start the process again
		x0 = x1
	}

	if solutionFound {
		// x1 is a solution within tolerance and maximum number of iterations
		fmt.Println("\nSolution: ", x1)
	} else {
		// Newton's method did not converge
		fmt.Println("Did not converge")
	}
}
