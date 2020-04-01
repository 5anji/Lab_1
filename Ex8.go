package main

import (
	"fmt"
	"math"
)

func exercise8() {
	var x, t, z float64 = 0.1, 0.5, 2.5
	var D, C float64

	D = math.Sin(math.Pow(math.E, 2*math.Cos(t) + z) - math.Pow(x, 3)) + math.Log10(z)
	C = (math.Pow(x, 3) + t - (math.Pow(math.Sin(math.Pow(x, 3)), 3))/(2*t + z)) / (math.Sqrt(math.Abs(x - math.Pow(t, 3)*z)))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", D)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", C)
	fmt.Printf("************************************\n")
}