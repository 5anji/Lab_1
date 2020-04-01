package main

import (
	"fmt"
	"math"
)

func exercise14() {
	var y, b, x float64 = 0.5, 3.2, 5.1
	var d, c float64

	d = math.Pow(y, 2*x) + math.Pow(b, -x)*math.Pow(math.Cos(x + math.Pow(y, 2)), 2) - math.Log10(b*x)
	c = (math.Pow(b, 2)*x + math.Pow(math.E, -x)*math.Cos(b*x)) / (b*x - math.Pow(math.E,  -x)*math.Sin(b*x) + 1)

	fmt.Printf("************************************\n")
	fmt.Printf("*           CALC RESULTS           *\n")
	fmt.Printf("************************************\n\v")
	fmt.Printf("Result 1 = %f\n", d)
	fmt.Printf("************************************\n\v")
	fmt.Printf("\t\tResult 2 = %f\n", c)
	fmt.Printf("************************************\n")
}