package sortlib

func MergeSort(arr []int) []int {
	// implement the merge sort algorithm
	if len(arr) == 0 {
		return []int{}
	}

	if len(arr) == 1 {
		return arr
	}

	return []int{1,2} 
}

func MergeInPlace(list []int) {
	i, j, k := 0, 0, 0

	// for loop that starts from 2, 4, 8, 16 until it rearches the half of the query 
	h := 2

	for h < len(list)/2 {
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

		h *= 2
	}

	return merged
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
