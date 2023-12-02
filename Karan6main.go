package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Prompt the user to enter a string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// Convert the input to lowercase for case-insensitive comparison
	input = strings.ToLower(input)

	startsWithI := strings.HasPrefix(input, "i")
	endsWithN := strings.HasSuffix(input, "n")
	containsA := strings.Contains(input, "a")

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
