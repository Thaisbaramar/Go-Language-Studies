package main

import "fmt"

func main() {

	var i, j, rows int

	rows = 5

	fmt.Println("\nInverted Pyramid Star Pattern")
	for i = rows; i > 0; i-- {
		for j = 0; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
