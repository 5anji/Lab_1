package main

import (
	"fmt"
	"math"
)

func exercise1() {
	const e = 2.71828
	var x, y, a, b = 2.2, 0.5, 2, 3
	var P, Q float64

	P = math.Cbrt(2*x*math.Sqrt((1 + y)/math.Abs(math.Pow(math.Sin(x), 3))))
	Q = math.Pow(e, (float64(-a)*x)) * math.Sqrt(x+1) + math.Pow(e, (float64(-b)*x)) + math.Sqrt(x+1.5)

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", P)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", Q)
	fmt.Printf("************************************\n")
}