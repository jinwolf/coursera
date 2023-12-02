package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	truncated := int(math.Trunc(input))
	fmt.Println("Truncated integer:", truncated)
}
