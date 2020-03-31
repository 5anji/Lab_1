package main

import (
	"fmt"
	"math"
)

func exercise5() {
	var a, b, x float64 = 2.0, 3.2, 0.5
	var U, F float64

	U = (2*a + math.Log10(b) + math.Abs(a/x)) / math.Pow(math.E, a*x) + (1/math.Tan(a) + (b/x))
	F = math.Sqrt(1 + math.Abs(math.Pow(math.Sin(x), 3))) + math.Pow(math.Cos(math.Pow(a*x, 3/5)), 3)

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", U)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", F)
	fmt.Printf("************************************\n")
}