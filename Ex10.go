package main

import (
	"fmt"
	"math"
)

func exercise10() {
	var a, x, y float64 = 2.1, 3, 3.2
	var U, T float64

	U = math.Pow(2.7*math.E, -2*x) + math.Cbrt(5*a - math.Sqrt(math.Abs(x - 7))) + math.Sqrt(math.Abs(a - x))
	T = (1 + (2*x - math.Pow(a, 3))/(math.Log10(y)))/(2 + math.Pow(math.Cos(math.Pow(x, 2)), 3)) + (1/math.Tan(x) + math.Pow(math.Tan(x), 2))/(3*math.Tan(x))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", U)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", T)
	fmt.Printf("************************************\n")
}