package main

import "fmt"

func main() {
	//Slice of integers
	x := []int{2, 4, 6, 8, 10, 11}
	sum := 0
	// Loop to calculate the sum
	for _, value := range x {
		sum += value
	}
	//Print it
	fmt.Println("The sum of the elements in the slice is:", sum)


	//  Create a map, Add key-value pairs to the map
	PercyJacksonBooksReadingOrderDictionary := map[string]int{
		"The Lightning Thief":      1,
		"The Sea of Monsters":       2,
		"The Titan's Curse":        3,
		"The Battle of the Labyrinth": 4,
		"The Last Olympian":        5,
		"The Lost Hero":            6,
		"The Son of Neptune":       7,
		"The Mark of Athena":       8,
		"The House of Hades":       9,
		"The Blood of Olympus":     10,
		"The Chalice of the Gods":   11,
	}

	// Retrieve and print 
	for book, order := range PercyJacksonBooksReadingOrderDictionary {
		fmt.Printf("%s: %d\n", book, order)
	}
}
