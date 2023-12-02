package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Swap function swaps the position of two adjacent elements in the slice
// if the former is greater than the later. After swapping it recurses into
// index = index - 1, if possible.
func Swap(slice []int, index int) {
	// Return if slice sorting has been exausted.
	if index+1 >= len(slice) {
		return
	}

	// If next value of the slice is smaller than the current one swap postions.
	if slice[index+1] < slice[index] {
		tempPrevSlice := slice[index]
		slice[index] = slice[index+1]
		slice[index+1] = tempPrevSlice
		if index != 0 {
			Swap(slice, index-1)
		}
	}
}

// BubbleSort iters the slice while calling Swap function to perform bubble
// sort algorithm.
func BubbleSort(slice []int) {
	for index := range slice {
		Swap(slice, index)
	}
}

func main() {

	// Request and parse input.
	fmt.Printf("Gimmie da slice! (up to 10 integers separated by a space):\n")
	reader := bufio.NewReader(os.Stdin)
	barr, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	rawInput := bytes.NewBuffer(barr).String()
	stringSliceInput := strings.Split(rawInput, " ")
	processedSlice := make([]int, 0)
	for _, value := range stringSliceInput {
		intValue, _ := strconv.Atoi(value)
		processedSlice = append(processedSlice, intValue)
	}

	// Sort slice and print result.
	BubbleSort(processedSlice)
	fmt.Printf("%d", processedSlice)
}
