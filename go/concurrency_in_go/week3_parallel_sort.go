/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
	"sync"
)

func getUserInput() []int {
	fmt.Println("Enter integer numbers separted by a space")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	values := strings.Split(scanner.Text(), " ")

	userInput := make([]int, len(values))

	for i, item  := range values {
		inputInteger, err := strconv.Atoi(item)

		if err == nil {
			userInput[i] = inputInteger
		}
	}

	return userInput
}
	
func mergeSlices(slices ...[]int) []int {
	var result [] int
	for _, s := range slices {
		result = append(result, s...)
	}
	sort.Ints(result)
	return result
}

func main() {
	// prompt the user to enter a series of integers
	originalSlice := getUserInput()

	// split the array into 4 parts
	n := len(originalSlice) / 4

	var wg sync.WaitGroup
	slices := make([][]int, 4)

	// create 4 go routines and each go routine sorting each array
	for i := 0; i < 4; i++ {
		start := i * n
		end := start + n
		slices[i] = originalSlice[start:end]

		// Sorting each slice in a goroutine
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sort.Ints(slices[i])
		}(i)
	}

	wg.Wait()

	// main function waits for all go routines to finish and merge them into one single sorted array

	//quick sort and merge sort?
	sortedSlice := mergeSlices(slices[0],slices[1],slices[2],slices[3])
	fmt.Println("Original input numbers:", originalSlice)	
	fmt.Println("Sorted   input numbers:", sortedSlice)
}
