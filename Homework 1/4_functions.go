package main

import("fmt")

//Function to do factorial
func factorial(x int) int {
	if x == 1 {
		return 1
	}
	factorialNumber := x * factorial(x-1)
	return factorialNumber 
}
// Calling the function and printing
func main() {
	var x int
	fmt.Print("Enter the number: ")
	fmt.Scan(&x)
	fmt.Println("The factorial of", x, "is", factorial(x))
}

