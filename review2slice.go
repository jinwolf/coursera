package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	mySlice := make([]int, 3)

	for {
		fmt.Println("Please enter an integer: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if input == "x" {
			break
		}

		integer, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		mySlice = append(mySlice, integer)
		sort.Ints(mySlice)
		fmt.Println("sorted slice: ", mySlice)
	}
	fmt.Println("final slice: ")
	fmt.Println(mySlice[3:])
}
