package main

import ("fmt")

func main() {
	var x int
	// User input
	fmt.Print("Enter the number: ")
	fmt.Scan(&x)
	//Check if odd or even
	if x%2 == 0 {
		fmt.Println(x, "is an even number")
	} else { fmt.Println(x, "is an odd number") 
	} 
	//Loop to print
	for i := 1; i <= x; i++ {
		fmt.Println(i)	
	}
}