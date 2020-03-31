package main

import (
	"fmt"
	"math"
)

func exercise4() {
	var x, y, a, t = 0.3, 1.1, 2, 3
	var R, U float64

	R = (1 + (2*x - math.Pow(y, 3)) / (math.Log10(float64(a)))) / (2 + math.Pow(math.Pow(x, 2), 3))
	U = math.Pow(math.E, float64(-t)*x)*math.Tan(float64(a)*x + y) - math.Sqrt(math.Abs(y*float64(t) + x))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", R)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", U)
	fmt.Printf("************************************\n")
}