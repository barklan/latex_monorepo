package main

import (
	"fmt"
	"math"
	"time"
)

const precision = 1e-4

type Fn func(float64) float64

func fnGleb(x float64) float64 {
	return math.Log(x) * math.Abs(math.Cos(128*x))
}

func tmpTask(x float64) float64 {
	return 2 / math.Sqrt(x*x+1)
}

func simpson(fn Fn, a, b float64) float64 {
	return ((b - a) / 6) * (fn(a) + 4*fn((a+b)/2) + fn(b))
}

func leftRect(fn Fn, a, b float64) float64 {
	return fn(a) * (b - a)
}

func rightRect(fn Fn, a, b float64) float64 {
	return fn(b) * (b - a)
}

func centerRect(fn Fn, a, b float64) float64 {
	return fn((b+a)/2) * (b - a)
}

func trapezoid(fn Fn, a, b float64) float64 {
	return (fn(a) + fn(b)) * (b - a) / 2
}

type Job struct {
	Method func(Fn, float64, float64) float64
	Target Fn
	Low    float64
	High   float64
	H      float64
	N      int64
	Teta   float64
}

func integrateSync(
	job Job,
) float64 {
	xNext := job.Low + job.H
	small := 0.0
	total := 0.0
	for x := job.Low; xNext <= job.High; xNext += job.H {
		small = job.Method(job.Target, x, xNext)
		total += small
		x = xNext
	}

	return total
}

func ascend(job Job) float64 {
	var n int64 = 1000
	var prev float64
	var res float64
	for {
		job.H = (job.High - job.Low) / float64(n)
		res = integrateSync(job)
		fmt.Printf("h = %0.10f; integral = %0.10f\n", job.H, res)
		if (job.Teta * math.Abs(res-prev)) < precision {
			fmt.Printf("FINAL RESULT: %.10f\n", res)
			fmt.Printf("final h = %.15f; final n = %d\n", job.H, n)
			break
		}
		prev = res
		n *= 2
	}
	return res
}

func main() {
	// a := 00000000000000000.1
	a := 0.0
	b := 1.0

	job := Job{
		Target: tmpTask,
		Low:    a,
		High:   b,
	}

	fmt.Printf("Precise: %.10f\n", 2*math.Log(1+math.Sqrt(2)))

	fmt.Printf("left\n")
	job.Method = leftRect
	n := 300_000_000
	job.H = (job.High - job.Low) / float64(n)
	res := integrateSync(job)
	fmt.Printf("n = %d, integral = %.10f\n\n", n, res)

	fmt.Printf("right\n")
	job.Method = rightRect
	n = 300_000_000
	job.H = (job.High - job.Low) / float64(n)
	res = integrateSync(job)
	fmt.Printf("n = %d, integral = %.10f\n\n", n, res)

	fmt.Printf("center\n")
	job.Method = centerRect
	n = 150_000
	job.H = (job.High - job.Low) / float64(n)
	res = integrateSync(job)
	fmt.Printf("n = %d, integral = %.10f\n\n", n, res)

	fmt.Printf("center\n")
	job.Method = centerRect
	n = 1_000
	job.H = (job.High - job.Low) / float64(n)
	res = integrateSync(job)
	fmt.Printf("n = %d, integral = %.10f\n\n", n, res)

	fmt.Printf("simpson\n")
	now := time.Now()
	job.Method = simpson
	job.Teta = 1.0 / 15.0
	ascend(job)
	elapsed := time.Since(now)
	fmt.Printf("Elapsed: %s\n\n", elapsed)
}

// func main() {
// 	a := 1.0
// 	b := 4.141593

// 	job := Job{
// 		Target: fnGleb,
// 		Low:    a,
// 		High:   b,
// 	}

// 	fmt.Printf("simpson\n")
// 	now := time.Now()
// 	job.Method = simpson
// 	job.Teta = 1.0 / 15.0
// 	ascend(job)
// 	elapsed := time.Since(now)
// 	fmt.Printf("Elapsed: %s\n\n", elapsed)

// 	fmt.Printf("left rect\n")
// 	now = time.Now()
// 	job.Method = leftRect
// 	job.Teta = 1.0 / 3.0
// 	ascend(job)
// 	elapsed = time.Since(now)
// 	fmt.Printf("Elapsed: %s\n\n", elapsed)

// 	fmt.Printf("right rect\n")
// 	now = time.Now()
// 	job.Method = rightRect
// 	job.Teta = 1.0 / 3.0
// 	ascend(job)
// 	elapsed = time.Since(now)
// 	fmt.Printf("Elapsed: %s\n\n", elapsed)

// 	fmt.Printf("center rect\n")
// 	now = time.Now()
// 	job.Method = centerRect
// 	job.Teta = 1.0 / 3.0
// 	ascend(job)
// 	elapsed = time.Since(now)
// 	fmt.Printf("Elapsed: %s\n\n", elapsed)

// 	fmt.Printf("trapezoid\n")
// 	now = time.Now()
// 	job.Method = trapezoid
// 	job.Teta = 1.0 / 3.0
// 	ascend(job)
// 	elapsed = time.Since(now)
// 	fmt.Printf("Elapsed: %s\n\n", elapsed)
// }
