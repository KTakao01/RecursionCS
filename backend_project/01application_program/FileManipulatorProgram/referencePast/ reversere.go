package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run file_manipulator.go reverse <inputfile> <outputfile>")
		return
	}
	//リファクタリング。分割するだけ

	command := os.Args[1]
	inputFile := os.Args[2]
	outputFile := os.Args[3]

	switch command {
	case "reverse":
		err := reverse(inputFile, outputFile)
		if err != nil {
			fmt.Println("Error processing reverse:", err)
		}
	}
}

// reverse ファイル内容を反転して、指定された出力ファイルに保存します
func reverse(inputFile, outputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("Error reading file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Error reading content: %w", err)
	}

	reversedData := reverseString(string(content))

	err = os.WriteFile(outputFile, []byte(reversedData), 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %w", err)
	}
	return nil
}

// reverseString は文字列を逆にして返します
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
