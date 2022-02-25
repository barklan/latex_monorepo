package main

import "math"

func fnGleb(x float64) float64 {
	return math.Log(x) * math.Abs(math.Cos(128*x))
}

func fnIlyana(x float64) float64 {
	return 1 / (x*(3*x+2) - 1)
}

func IlyanaPecise() float64 {
	return math.Log(3.0) / 4.0
}

func x3(x float64) float64 {
	return x * x * x
}

func simpson(fn func(float64) float64, a, b float64) float64 {
	return ((b - a) / 6) * (fn(a) + 4*fn((a+b)/2) + fn(b))
}

func leftRect(fn func(float64) float64, a, b float64) float64 {
	return fn(a) * (b - a)
}

func leftRectOptim(fn func(float64) float64, a float64, h float64) float64 {
	return fn(a) * h
}
