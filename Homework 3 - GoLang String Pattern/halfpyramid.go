package main

import "fmt"

func main() {

	var i, j, rows int

	rows = 5

	fmt.Println("\nRight Half Pyramid Pattern")
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Print("\n")
	}
}
