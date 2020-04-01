package main

import (
	"fmt"
	"math"
)

func exercise12() {
	var x, y, c float64 = 2, 1.1, 5.6
	var P, Q float64

	P = math.Cbrt(math.Pow(1/math.Tan(x), 2) + y) + math.Pow(math.E, x*y) + math.Sqrt(10 + math.Pow(x + y, 2/3))
	Q = 2*y + 7*x*(math.Log10(y*x) + math.Sin(x)) / (math.Sqrt(math.Abs(x - 6*y))) + 2*c*x

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", P)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", Q)
	fmt.Printf("************************************\n")
}