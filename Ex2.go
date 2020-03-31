package main

import (
	"fmt"
	"math"
)

func exercise2() {
	var a, x, b = 2.2, 0.5, 2
	var Y, D float64

	Y = a * math.Pow(math.Cos(math.Pow(x, 3/4)), 2) + math.Sqrt(math.Abs(float64(b)*x+a))
	D = math.Pow(math.Tan(math.Pow(x+float64(b), 4)), 3) + math.Log10(a+math.Pow(x, 2))

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", Y)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", D)
	fmt.Printf("************************************\n")
}