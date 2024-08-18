package main

import (
	"fmt"
)

// FileNotFoundError is a custom error type for representing a file not found error.
type FileNotFoundError struct {
	FileName string
}

// Error implements the error interface for FileNotFoundError.
func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("File not found: %s", e.FileName)
}

// simulateFileOperation simulates an operation that may result in a FileNotFoundError.
func simulateFileOperation(fileName string) error {
	// Simulate a condition where the file is not found
	if fileName == "" {
		return &FileNotFoundError{FileName: fileName}
	}

	// Simulate a successful operation if the file is found
	fmt.Println("File operation completed successfully.")
	return nil
}

func main() {
	// Simulate a file operation with a valid file name
	err := simulateFileOperation("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error, operation successful.")
	}

	// Simulate a file operation with an invalid (empty) file name
	err = simulateFileOperation("")
	if err != nil {
    		// Check if the error is of type FileNotFoundError
    		if _, ok := err.(*FileNotFoundError); ok {
        		fmt.Printf("Custom Error: %s\n", "Empty file name provided")
    		} else {
        	// Handle other types of errors
        	fmt.Println("Error:", err)
    		}
	} else {
    		fmt.Println("No error, operation successful.")
}



}
