package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

func main() {

	userInput := getUserInput()

	fmt.Println("User input ")
	fmt.Println(userInput)

	fmt.Println("Sorted result")

	BubbleSort(userInput)

	fmt.Println(userInput)

}

func BubbleSort(inputSlice []int) {

	sliceLen := len(inputSlice)

	for i := 0; i < sliceLen-1; i++ {

		for j := 0; j < sliceLen-i-1; j++ {

			if inputSlice[j] > inputSlice[j+1] {
				swap(inputSlice, j)
			}
		}
	}

}

func swap(inputSlice []int, index int) {

	inputSlice[index], inputSlice[index+1] = inputSlice[index+1], inputSlice[index]

}

func getUserInput() []int {

	fmt.Println("Enter integer numbers separated by a space")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	values := strings.Split(scanner.Text(), " ")

	userInput := make([]int, len(values))

	for i, j := range values {

		inputInteger, err := strconv.Atoi(j)

		if err == nil {
			userInput[i] = inputInteger
		}

	}

	return userInput

}
