package numerical

func Secant(f func(float64) float64,
	a, b, tolerance float64,
	maxIterations int) (float64, int, error) {
	var a_n, b_n, m_n, f_m_n float64
	if f(a)*f(b) >= 0 {
		panic("oh no")
	}
	a_n = a
	b_n = b
	for n := 1; n <= maxIterations; n++ {
		m_n = a_n - f(a_n)*(b_n-a_n)/(f(b_n)-f(a_n))
		f_m_n = f(m_n)
		if f(a_n)*f_m_n < 0 {
			b_n = m_n
		} else if f(b_n)*f_m_n < 0 {
			a_n = m_n
		} else if f_m_n == 0 || (b_n-a_n)/2 < tolerance {
			return m_n, n, nil
		} else {
			panic("oh no")
		}
	}
	return a_n - f(a_n)*(b_n-a_n)/(f(b_n)-f(a_n)), maxIterations, nil
}
