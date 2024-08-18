package main

import (
	"fmt"
)

func main() {
	// Prompt the user to enter the number and 
	// read and store the first number, handle potential scanning errors
	var x, y int
	fmt.Print("Enter the first number: ")
	_, err1 := fmt.Scan(&x)
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}

	fmt.Print("Enter the second number: ")
	_, err2 := fmt.Scan(&y)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	// Check for division by zero
	if y == 0 {
		fmt.Println("Error: Division by zero")
		return
	}

	// Check if the first number is divisible by the second number
	if x%y == 0 {
		fmt.Printf("%d is divisible by %d\n", x, y)
	} else {
		fmt.Printf("%d is not divisible by %d\n", x, y)
	}
}