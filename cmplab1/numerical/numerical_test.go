package numerical

import (
	"math"
	"testing"
)

func fn(x float64) float64 {
	return math.Pow(x, 2) - 2
}

func TestBisection(t *testing.T) {
	got, _ := Bisection(fn, 0, 2, 0.01, 1000000000)
	low := 1.41
	high := 1.5

	if got < low || got > high {
		t.Errorf("got %v", got)
	}
}

func TestSecant(t *testing.T) {
	got, _ := Secant(fn, 0, 2, 0.01, 10)
	low := 1.41
	high := 1.5

	if got < low || got > high {
		t.Errorf("got %v", got)
	}
}
