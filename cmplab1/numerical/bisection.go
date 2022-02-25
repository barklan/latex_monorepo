package numerical

// https://en.wikipedia.org/wiki/Bisection
func Bisection(f func(float64) float64,
	a, b, tolerance float64,
	maxIterations int) (float64, int, error) {
	if f(a)*f(b) >= 0 {
		panic("oh no")
	}

	for n := 1; n <= maxIterations; n++ {
		c := (a + b) / 2
		if f(c) == 0 || (b-a)/2 < tolerance {
			return c, n, nil
		}
		if f(c)*f(a) > 0 {
			a = c
		} else {
			b = c
		}
	}

	panic("maxIterations exceeded")
}
