package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestPolinomeParse(t *testing.T) {
	p := Polinome{
		Degree: 4,
		Roots: []*big.Rat{
			big.NewRat(2, 1),
			big.NewRat(7, 1),
			big.NewRat(13, 1),
			big.NewRat(14, 1),
		},
		A: big.NewRat(5, 1),
	}

	p.Parse()

	got := p.Coefs

	want := []*big.Rat{
		big.NewRat(5, 1),
		big.NewRat(-180, 1),
		big.NewRat(2195, 1),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
