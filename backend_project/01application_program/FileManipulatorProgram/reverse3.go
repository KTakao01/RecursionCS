package main

import (
	"fmt"
	"os"
)

// FileReader interface abstracts the file reading dependency.
type FileReader interface {
	ReadFile(filename string) ([]byte, error)
}

// FileWriter interface abstracts the file writing dependency.
type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type defaultFileOperations struct{}

func (d *defaultFileOperations) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (d *defaultFileOperations) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run file_manipulator.go reverse <inputfile> <outputfile>")
		return
	}

	command := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	ops := &defaultFileOperations{}

	switch command {
	case "reverse":
		err := reverse(inputFile, outputFile, ops, ops)
		if err != nil {
			fmt.Println("Error processing reverse:", err)
		}
	}
}

func reverse(inputFile, outputFile string, reader FileReader, writer FileWriter) error {
	content, err := reader.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading content: %w", err)
	}

	reversedData := reverseString(string(content))

	err = writer.WriteFile(outputFile, []byte(reversedData), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}
	return nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
