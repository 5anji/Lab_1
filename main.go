package main

import "fmt"

func main() {
	var selector int8 = -1;

	for selector != 0 {
		fmt.Printf("# 0 - exit\n")
		fmt.Printf("Select an exercise (1-15): ")
		fmt.Scanf("%d", &selector)

		switch selector {
		case 0:
		case 1: exercise1()
		case 2: exercise2()
		case 3: 
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
		case 9:
		case 10:
		case 11:
		case 12:
		case 13:
		case 14:
		case 15: 
		}
	}
	
}