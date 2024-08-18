package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

// Book represents a book with title and author.
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

// ReadJSONFile reads a JSON file and returns the content as a byte slice.
func ReadJSONFile(filePath string) ([]byte, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// WriteJSONFile writes data to a JSON file.
func WriteJSONFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ParseJSONData parses JSON data into a slice of Book structs.
func ParseJSONData(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// FilterBooks filters books based on a given condition.
func FilterBooks(books []Book, condition func(Book) bool) []Book {
	var filteredBooks []Book
	for _, book := range books {
		if condition(book) {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return filteredBooks
}

// SortBooks sorts books alphabetically by title.
func SortBooks(books []Book) {
	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})
}

// ModifyBooks modifies the author of each book.
func ModifyBooks(books []Book, newAuthor string) {
	for i := range books {
		books[i].Author = newAuthor
	}
}

func main() {
	// Read the JSON file
	jsonFilePath := "books.json"
	jsonData, err := ReadJSONFile(jsonFilePath)
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Parse JSON data into Go structs
	books, err := ParseJSONData(jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}

	// Perform operations (filtering, sorting, modification)
	filteredBooks := FilterBooks(books, func(book Book) bool {
		return book.Author == "John Doe"
	})

	SortBooks(filteredBooks)
	ModifyBooks(filteredBooks, "New Author")

	// Convert modified data back to JSON
	modifiedData, err := json.MarshalIndent(filteredBooks, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON data:", err)
		return
	}

	// Write the modified data to a new JSON file
	newJSONFilePath := "modified_books.json"
	err = WriteJSONFile(newJSONFilePath, modifiedData)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Modified data written to", newJSONFilePath)
}
