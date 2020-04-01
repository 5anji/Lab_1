package main

import (
	"fmt"
	"math"
)

func exercise9() {
	var x, y, b float64 = 0.1, 2.2, 3
	var Q, P float64

	Q = 1 + (2*b + math.Log10(math.Pow(x, 2)))/(2*x + math.Sqrt(math.Abs(2*b - math.Pow(y, 2))))
	P = 1/math.Tan(math.Pow(math.E, 2*math.Cos(y) + b) - math.Pow(x, 3))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", Q)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", P)
	fmt.Printf("************************************\n")
}