package main

import (
	"math/rand"
	"os"
)

var letters = []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789\n.,")

const (
	len = 342190
	fp  = "/Users/jayson/Bucket/wc/text_maker/text.txt"
)

// CreateTestFile is a function that creates a new text file or erases all content if the file already exists.
// Parameters:
//
//	fp (string): The file path where the file will be created.
//	len (int): The number of random characters to be written into the file.
//
// Returns:
//
//	error: An error object that describes the error, if any occurred. Otherwise, it returns nil.
//
// Example:
//
//	err := CreateTestFile("/tmp/testfile.txt", 100)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Note:
//
//	The function writes random characters into the file. The characters are chosen from a predefined 'letters' slice.
//	If the function encounters an error while creating the file or writing to it, it will return the error immediately.
//	The function ensures that the file is properly closed before it returns, even if an error occurs.
func main() {
	// Create the txt file or erase all content if it exits
	f, err := os.Create(fp)
	if err != nil {
		panic(err)
	}
	defer func() error {
		if err := f.Close(); err != nil {
			panic(err)
		}
		return nil
	}()

	for i := 0; i < len; i++ {
		l := letters[rand.Intn(64)] // Len of letters
		_, err := f.WriteString(string(l))
		if err != nil {
			panic(err)
		}
	}
}
