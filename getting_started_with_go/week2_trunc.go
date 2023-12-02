// Write a program which prompts the user to enter a floating point number 
// and prints the integer which is a truncated version of the floating point
// number that was entered. Truncation is the process of removing the digit
// to the right of the decimal place.

// Submit your source code for the program, “trunc.go”.

package main
import (
	"fmt"
	"math"
)

func main() {
	var floatNum float64

	fmt.Print("Enter a floating point number:")
	_, err := fmt.Scan(&floatNum)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	truncatedNum := int(math.Trunc(floatNum))
	fmt.Println("Truncated number:", truncatedNum)
}
