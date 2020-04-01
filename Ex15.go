package main

import (
	"fmt"
	"math"
)

func exercise15() {
	var y, x, c float64 = 0.5, 5.1, 3
	var A, B float64

	A = math.Pow(math.E, x)*math.Pow(x + 1, 1/5) + math.Pow(math.E, x*y)*math.Sqrt(x + 1.5)
	B = 2*y + 7*x*(math.Log10(y*x) + math.Sin(x)) / (math.Sqrt(math.Abs(x - 6*y))) + 2*c*x

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", A)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", B)
	fmt.Printf("************************************\n")
}