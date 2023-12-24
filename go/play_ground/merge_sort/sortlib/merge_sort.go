package sortlib

import (
	"fmt"
)

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	// Merge the two halves
	return MergeSlices(left, right)
}

func MergeInPlace(list []int) {
	// for loop that starts from 2, 4, 8, 16 until it rearches the half of the query 
	// let's visualize 8 elements in the int list
	// 3, 1, 2, 14, 35, 14, 10, 8
	// [1, 3] [2, 14], [35, 14], [10, 8]
	// [1, 2, 3, 14]
	
	h := 1

	fmt.Println("Sorting started", h)
	for h <= (len(list) + 1)/2 {
		i := h - 1 
		j := h

		fmt.Println("-------------------")
		for i < h && j < h * 2 {
			if list[i] <= list[j] {
				i++
			} else {
				Swap(&list[i], &list[j])
				j++
			}

		}
/*
		for i < len(a) {
			merged[k] = a[i]
			i++
			k++
		}

		for j < len(b) {
			merged[k] = b[j]
			j++
			k++
		}
		*/

		h *= 2
	}

}

func MergeSlices(a, b []int) []int {
	merged := make([]int, len(a) + len(b))
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}

		k++
	}

	for i < len(a) {
		merged[k] = a[i]
		i++
		k++
	}

	for j < len(b) {
		merged[k] = b[j]
		j++
		k++
	}

	return merged
}

func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}
