package main

import (
	"fmt"
	"math"
)

func exercise11() {
	var x, y, m float64 = 1.7, 0.5, 3
	var z, f float64

	z = (math.Sin(x))/(math.Sqrt(1 + math.Pow(m, 2)*math.Pow(math.Sin(x), 2))) - 2*m*math.Log10(y*x)
	f = math.Pow(math.E, -2*x)*math.Cbrt(x + 1) + math.Pow(math.E, x*y)*math.Sqrt(x + 1.5)

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", z)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", f)
	fmt.Printf("************************************\n")
}