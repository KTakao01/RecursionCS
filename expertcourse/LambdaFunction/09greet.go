// https://recursionist.io/dashboard/problems/433
package main

import (
	"fmt"
	"strings"
)

func loud(text string) string {
	return strings.ToUpper(text)
}

func quiet(text string) string {
	return strings.ToLower(text)
}

func reverse(text string) string {
	runes := []rune(text)
	reverse := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		reverse = append(reverse, runes[i])
	}
	return string(reverse)
}

func repeat(text string) string {
	rText := []rune(text)
	tCopy := make([]rune, len(rText))
	copy(tCopy, rText)
	rText = append(rText, []rune(" ")...)
	rText = append(rText, tCopy...)
	return string(rText)
}

func greet(time int, text string, f func(string) string) {
	if time >= 0 && time < 12 {
		fmt.Printf("Good Morning %s\n", f(text))
	} else if time >= 12 && time < 18 {
		fmt.Printf("Good Afternoon %s\n", f(text))
	} else {
		fmt.Printf("Good Evening %s\n", f(text))
	}
}

func main() {
	greet(1, "John", loud)
	greet(2, "John", quiet)
	greet(13, "John", reverse)
	greet(19, "John", repeat)
	greet(13, "Leslie Emmanuel Beadon", loud)
	greet(19, "Leslie Emmanuel Beadon", quiet)
	greet(5, "Leslie Emmanuel Beadon", reverse)
	greet(1, "Leslie Emmanuel Beadon", repeat)
}
