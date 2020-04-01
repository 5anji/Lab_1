package main

import (
	"fmt"
	"math"
)

func exercise6() {
	var r, d, a float64 = 1.5, 2, 0.2
	var V, T float64

	V = r*d*(1 + (r*math.Cos(a))/(math.Sqrt(math.Pow(r, 2) + math.Pow(math.Sin(a), 2)))) + math.Pow(1/math.Tan(a), 3)
	T = (math.Pow(math.Sin(math.Pow(a, 3)), 2))/(math.Log10(a) + math.Pow(math.E, r)) + 1 + math.Cbrt(d+4)

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", V)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", T)
	fmt.Printf("************************************\n")
}