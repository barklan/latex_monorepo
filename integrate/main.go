package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
)

const precision = 1e-4

type Report struct {
	ToAdd float64
	Cnt   int
}

type Job struct {
	Method func(func(float64) float64, float64, float64) float64
	Target func(float64) float64
	Low    float64
	High   float64
	N      int64
}

type JobH struct {
	Method func(func(float64) float64, float64, float64) float64
	Target func(float64) float64
	Low    float64
	High   float64
	H      float64
}

func Pretty(c <-chan Report, wg *sync.WaitGroup, n int64) {
	defer wg.Done()

	var total float64

	bar := progressbar.Default(n)
	for report := range c {
		total += report.ToAdd
		bar.Add(report.Cnt)
		if bar.State().CurrentPercent > 0.999999 {
			fmt.Printf("result: %.10f\n", total)
			// fmt.Printf("state: %.10f", bar.State().CurrentPercent)
			break
		}
	}
}

func Spread(job Job, workers int) {
	c := make(chan Report, workers)
	h := (job.High - job.Low) / float64(job.N)

	partN := float64(job.N) / float64(workers)
	toAdd := h * partN

	var wg sync.WaitGroup
	wg.Add(workers + 1)

	go Pretty(c, &wg, job.N)

	for i := 0; i < workers; i++ {
		low := job.Low + (float64(i) * toAdd)
		high := low + toAdd
		if i == workers-1 {
			high = job.High
		}
		go integrate(c, &wg, JobH{
			Method: job.Method,
			Target: job.Target,
			Low:    low,
			High:   high,
			H:      h,
		})
	}

	wg.Wait()
}

func integrate(
	report chan<- Report,
	wg *sync.WaitGroup,
	job JobH,
) {
	defer wg.Done()
	xNext := job.Low + job.H
	small := 0.0
	toAdd := 0.0
	cnt := 0
	for x := job.Low; xNext <= job.High; xNext += job.H {
		small = job.Method(job.Target, x, xNext)
		// small = job.Method(job.Target, x, job.H)
		toAdd += small
		x = xNext
		cnt++
		if cnt%10_000_000 == 0 {
			report <- Report{
				ToAdd: toAdd,
				Cnt:   cnt,
			}
			cnt = 0
			toAdd = 0.0
		}
	}

	report <- Report{
		ToAdd: toAdd,
		Cnt:   cnt,
	}
}

func ascend(
	correction float64,
	method func(func(float64) float64, float64, float64) float64,
	target func(float64) float64,
	low,
	high float64,
) float64 {
	// integralN :=
	// for n := 1; ; n *= 2 {
	// }
	return 0.0
}

// func interactive(
// 	// wg *sync.WaitGroup,
// 	method func(func(float64) float64, float64, float64) float64,
// 	target func(float64) float64,
// 	low,
// 	high float64,
// 	n float64,
// ) {
// 	// defer wg.Done()
// 	res := integrate(method, target, low, high, n)
// 	fmt.Printf("n: %f; res: %.10f\n", n, res)
// }

func main() {
	cpus := runtime.NumCPU()
	a := 1.0
	b := 4.141593
	// b := 21000.0
	// var wg sync.WaitGroup
	fmt.Printf("Precise: %.10f\n", IlyanaPecise())
	var nSimpson int64 = 2.1e8
	fmt.Printf("simpson\n")
	now := time.Now()
	Spread(
		Job{
			Method: simpson,
			Target: fnGleb,
			Low:    a,
			High:   b,
			N:      nSimpson,
		},
		cpus,
	)
	elapsed := time.Since(now)
	fmt.Printf("Elapsed: %s\n", elapsed)

	var nLeft int64 = 1.2e8
	fmt.Printf("leftRect\n")
	Spread(
		Job{
			Method: leftRect,
			Target: fnGleb,
			Low:    a,
			High:   b,
			N:      nLeft,
		},
		cpus,
	)
	// interactive(leftRect, fnIlyana, a, b, nLeft)
	// wg.Wait()
}
