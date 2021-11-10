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
		big.NewRat(-10080, 1),
		big.NewRat(12740, 1),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCombinations(t *testing.T) {
	set := []*big.Rat{
		big.NewRat(1, 1),
		big.NewRat(2, 1),
		big.NewRat(3, 1),
	}

	got := Combinations(set, 2)

	want := [][]*big.Rat{
		{big.NewRat(1, 1), big.NewRat(2, 1)},
		{big.NewRat(1, 1), big.NewRat(3, 1)},
		{big.NewRat(2, 1), big.NewRat(3, 1)},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMultiSumCombinations(t *testing.T) {
	combinations := [][]*big.Rat{
		{big.NewRat(1, 1), big.NewRat(2, 1)},
		{big.NewRat(1, 1), big.NewRat(3, 1)},
		{big.NewRat(2, 1), big.NewRat(3, 1)},
	}

	got := MultiSumCombinations(combinations)

	want := big.NewRat(11, 1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
