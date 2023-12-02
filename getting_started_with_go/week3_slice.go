// 1. prompts the user to enter intergers
// 2. stores the integers in a sorted slice
// 3. the program should be written as a loop
// 4. create an empty integer slice of size 3
// 5. During each loop, prompts to enter an integer
// 6. sort  the list and print it
// 7. the slice must grow in size
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var input string 
	slice := make([]int, 0, 3)

	for true {
		fmt.Print("Enter an integer: ")
		fmt.Scan(&input)

		if input == "X" || input == "x" {
			break
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid input. Please enter an integer")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice)
		fmt.Println("Sorted slice:", slice)
	}

	return 
}
