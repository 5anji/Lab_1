package main

import (
	"fmt"
	"math"
)

func exercise13() {
	var x, y, b float64 = 2, 3.5, 0.5
	var S, Y float64

	S = math.Abs(math.Pow(x, y/x) - math.Cbrt(y/x))
	Y = b*math.Pow(math.Tan(x), 2) - y / (math.Pow(1/math.Tan(x/y), 2))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", S)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", Y)
	fmt.Printf("************************************\n")
}