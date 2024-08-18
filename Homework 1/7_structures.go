package main

import "fmt"

// Define a struct representing a book
type Book struct {
	Title           string
	Author          string
	PublicationYear int
}

// Define a method for the Book struct to print its details
func (b Book) PrintDetails() {
	fmt.Printf("Title: %s\nAuthor: %s\nPublication Year: %d\n", b.Title, b.Author, b.PublicationYear)
}

func main() {
	// Create instances of the book struc
	books := []Book{
		{
			Title:           "Throne of Glass",
			Author:          "Sarah J. Maas",
			PublicationYear: 2012,
		},
		{
			Title:           "The Assassin's Blade",
			Author:          "Sarah J. Maas",
			PublicationYear: 2013,
		},
		{
			Title:           "Crown of Midnight",
			Author:          "Sarah J. Maas",
			PublicationYear: 2013,
		},
		{
			Title:           "Heir of Fire",
			Author:          "Sarah J. Maas",
			PublicationYear: 2014,
		},
		{
			Title:           "A Court of Thorns and Roses",
			Author:          "Sarah J. Maas",
			PublicationYear: 2015,
		},
		{
			Title:           "Queen of Shadows",
			Author:          "Sarah J. Maas",
			PublicationYear: 2015,
		},
		{
			Title:           "A Court of Mist and Fury",
			Author:          "Sarah J. Maas",
			PublicationYear: 2016,
		},
		{
			Title:           "Empire of Storms",
			Author:          "Sarah J. Maas",
			PublicationYear: 2016,
		},
		{
			Title:           "A Court of Wings and Ruin",
			Author:          "Sarah J. Maas",
			PublicationYear: 2017,
		},
		{
			Title:           "Tower of Dawn",
			Author:          "Sarah J. Maas",
			PublicationYear: 2017,
		},
		{
			Title:           "A Court of Frost and Starlight",
			Author:          "Sarah J. Maas",
			PublicationYear: 2018,
		},
		{
			Title:           "Catwoman: Soulstealer",
			Author:          "Sarah J. Maas",
			PublicationYear: 2018,
		},
		{
			Title:           "Kingdom of Ash",
			Author:          "Sarah J. Maas",
			PublicationYear: 2018,
		},
		{
			Title:           "House of Earth and Blood",
			Author:          "Sarah J. Maas",
			PublicationYear: 2020,
		},
		{
			Title:           "A Court of Silver Flames",
			Author:          "Sarah J. Maas",
			PublicationYear: 2021,
		},
		{
			Title:           "House of Sky and Breath",
			Author:          "Sarah J. Maas",
			PublicationYear: 2022,
		},
		{
			Title:           "House of Flame and Shadow",
			Author:          "Sarah J. Maas",
			PublicationYear: 2024,
		},
	}

	// Call the method to print the details of the books
	for i, book := range books {
		fmt.Printf("Book %d Details:\n", i+1)
		book.PrintDetails()
		fmt.Println()
	}
}





