// 1. design func merge_sort(int []) int [] 
// 2 design swap 
// 3.design merge
// 4.write tests

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"merge_sort/sortlib"
)

func takeUserInput() []int {
	// prompt user with a list of integers separated with a space
	// parse the string into a slice of integers
	fmt.Println("Enter a series of integers separated by space:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim the new line
	substrings := strings.Fields(strings.TrimSpace(input))

	var integers []int
	for _, str := range substrings {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("error converting '%s' to integer: %s\n", str, err)
			continue
		}
		integers = append(integers, num)
	}
	return integers
}

func main() {
	input := takeUserInput() 
	sorted := sortlib.MergeSort(input)
	fmt.Println(sorted)
	
}
