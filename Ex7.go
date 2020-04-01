package main

import (
	"fmt"
	"math"
)

func exercise7() {
	var x, y, z float64 = 1, 1, 2.7
	var A, B float64

	A = math.Pow(math.Sqrt(math.Sin(math.Pow(x, 3)) + 2*z - 3*y), 1/5)
	B = (math.Pow(math.E, x + math.Sin(y)) + math.Pow(y, 4)) / (2*math.Log10(x) + math.Cos(math.Pow(y, 3)))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", A)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", B)
	fmt.Printf("************************************\n")
}