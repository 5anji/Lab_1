package main

import "fmt"

func main() {
	var selector int8 = -1;

	for selector != 0 {
		fmt.Printf("# 0 - exit\n")
		fmt.Printf("# Select an exercise (1-15): ")
		fmt.Scanf("%d", &selector)

		switch selector {
		case 0:
		case 1: exercise1()
		case 2: exercise2()
		case 3: exercise3()
		case 4: exercise4()
		case 5: exercise5()
		case 6: exercise6()
		case 7: exercise7()
		case 8: exercise8()
		case 9: exercise9()
		case 10: exercise10()
		case 11: exercise11()
		case 12: exercise12()
		case 13: exercise13()
		case 14: exercise14()
		case 15: exercise15()
		default: fmt.Printf("Invalid expression. Try again.\n")
		}
	}
	
}