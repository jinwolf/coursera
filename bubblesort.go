// write a bubble sort algorithm in go
// the program prompts the user to type in a sequence of up to 10 integers
// the program will print out the integers on one line in a sorted order from the least to the greatest
// write a function, BubbleSort(nums []int) which modifies the slice so the elements will be sorted
// write a function, Swap(nums []int, index int) 
package main

import "fmt"

func main() {
	var input int
	var nums []int
	
	fmt.Print("Please enter each integer followed by the Enter key. When you're finished, or if you want to enter fewer than 10 integers, type any non-numeric character and press Enter to proceed with the sorting\n")
	for i:=0; i < 10; i++ {
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			break
		}

		nums = append(nums, input)
	}

	fmt.Println("User input: ", nums)
	BubbleSort(nums)
	fmt.Println("Sorted: ", nums)
}

func Swap(nums []int, index int) {
	tmp := nums[index]
	nums[index] = nums[index + 1]
	nums[index + 1] = tmp
}

func BubbleSort(nums []int) {
	n := len(nums)
	for i:=0; i < n; i++ {
		for j:=0; j < n-i-1; j++ {
			if nums[j] > nums[j + 1] {
				Swap(nums, j)
			}
		}
	}
}
