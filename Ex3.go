package main

import (
	"fmt"
	"math"
)

func exercise3() {
	var t, c, x, m, a = 2.2, 0.5, 2, 1.5, 2
	var S, X float64

	S = (math.Pow(math.E, 2*float64(x)) + math.Abs(c*math.Sin(t)) + math.Cbrt(c*t + float64(x))) / (m*math.Cos(c*t*math.Sin(t) + c))
	X = c + math.Log10(float64(a) + math.Pow(t, 2)) + math.Pow(math.Tan(m/t), 2)

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", S)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", X)
	fmt.Printf("************************************\n")
}