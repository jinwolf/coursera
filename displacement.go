package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Write a program which first prompts the user
// to enter values for acceleration, initial velocity, and initial displacement.
// Then the program should prompt the user to enter a value for time and the
// program should compute the displacement after the entered time.
// For example, let’s say that I want to assume
// the following values for acceleration, initial velocity, and initial
// displacement: a = 10, vo = 2, so = 1. I can use the
// following statement to call GenDisplaceFn() to
// generate a function fn which will compute displacement as a function of time.
func main() {

	fmt.Println("Insert (separated by spaces) acceleration, initial velocity, and initial displacement")
	paramsInput := ReadLine()
	// read a string with 3 values
	paramsSlice := strings.Split(paramsInput, " ")
	if len(paramsSlice) > 3 {
		fmt.Println("Invalid input, more than 3 elements inserted")
		return
	}
	// convert the input strings to float values
	floatParams, err := ToFloatSlice(paramsSlice)
	if err != nil {
		log.Fatal(err)
	}
	// read a string with a single value
	fmt.Println("Insert an amount of time in seconds")
	timeInput := ReadLine()
	timeSlice := strings.Split(timeInput, " ")
	if len(timeSlice) != 1 {
		fmt.Println("Insert one and only one value")
		return
	}
	// get the time value as a float
	floatTime, err := ToFloatSlice(timeSlice)
	if err != nil {
		log.Fatal(err)
	}

	fn := GenDisplaceFn(floatParams[0], floatParams[1], floatParams[2])

	fmt.Printf("Displacement after %f seconds is %f\n", floatTime[0], fn(floatTime[0]))

}

//You will need to define and use a function
//called GenDisplaceFn() which takes three float64
//arguments, acceleration a, initial velocity v0, and initial
//displacement s0. GenDisplaceFn()
//should return a function which computes displacement as a function of time,
//assuming the given values acceleration, initial velocity, and initial
//displacement.
// The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
// float64 argument which is the displacement travelled after time t.

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	//s = ½ * a * tˆ2 + v0 * t + s0
	return func(t float64) float64 {
		return (a * math.Pow(t, 2) / 2) + (v0 * t) + s0
	}
}

func ReadLine() string {
	// needs bufio to handle strings containing spaces
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	return scanner.Text()
}

func ToFloatSlice(sli []string) ([]float64, error) {
	floatSlice := make([]float64, 0, len(sli))

	for _, v := range sli {
		num, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return []float64{}, errors.New("Invalid input element: " + v)
		}
		floatSlice = append(floatSlice, num)
	}

	return floatSlice, nil
}
