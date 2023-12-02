package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	pie := make([]int, 3)
	pointer := 0
	var userInput string

	for i := 0; i < len(pie)+1; i = i + 1 {

		fmt.Printf("Please enter a value to add to the slice: ")
		fmt.Scanln(&userInput)

		// check if the slice is at capacity and increase it to keep the loop from ending when needed
		if pointer == len(pie)-1 {
			pie = append(pie, 0)
		}

		//check if the user provided value is numeric and store it to the slice
		if userNum, err := strconv.Atoi(userInput); err == nil {
			pie[pointer] = userNum
		}

		// check for x to break
		if userInput == "x" {
			break
		}

		//Store the values to a new slice for sorting and display pposes without manuipulating the slice that manages data collection
		displayPie := make([]int, len(pie))
		copy(displayPie, pie)
		sort.Ints(displayPie)

		for ps := range displayPie {
			fmt.Printf("After sorting, position %d of the slice now has a value of %d\n", ps, displayPie[ps])
		}
		pointer++
	}

}
