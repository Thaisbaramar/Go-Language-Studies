package main

import (
	"fmt"
	"time"
)

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}

func main() {
	// Input: positive integer for which factorial needs to be calculated
	n := 10

	// Channel to collect the result from the goroutine
	resultCh := make(chan int)

	// Measure the start time
	startTime := time.Now()

	// Launch the goroutine to calculate factorial
	go factorial(n, resultCh)

	// Collect the result from the goroutine
	totalResult := <-resultCh

	// Close the result channel
	close(resultCh)

	// Measure the end time
	endTime := time.Now()

	// Display the result and execution time
	fmt.Printf("Factorial of %d is: %d\n", n, totalResult)
	fmt.Printf("Execution time: %v\n", endTime.Sub(startTime))
}