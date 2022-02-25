package numerical

import (
	"fmt"
	"math"
)

func FixedPointIteration(fn, fng func(float64) float64,
	x0, tolerance float64,
	maxIterations int) {
	step := 1
	flag := 1
	condition := true
	var x1 float64
	for condition {
		x1 = fng(x0)
		fmt.Printf("Iteration-%d, x1 = %0.6f and fn(x1) = %0.6f\n",
			step, x1, fn(x1))
		x0 = x1

		step = step + 1

		if step > maxIterations {
			flag = 0
			break
		}

		condition = math.Abs(fn(x1)) > tolerance
	}

	if flag == 1 {
		fmt.Println("\nSolution: ", x1)
	} else {
		fmt.Printf("\nNot Convergent.")
	}
}
