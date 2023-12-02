package main

import "fmt"

func main() {
	var a float64 //acceleration
	var vo float64 //initial velocity
	var so float64 //initial displacement
	var t float64 //time

	// prompt the uer to enter acceleration
	GetUserInput("Please enter the acceleration", &a)
	
	// prompt the user to enter initial velocity
	GetUserInput("Please enter the velocity", &vo)

	// prompt the user to enter initial displacement
	GetUserInput("Please enter the displacement", &so)

	fn := GenDisplaceFn(a, vo, so)

	// prompt the user to enter a value for time
	GetUserInput("Please enter the time", &t)
	
	fmt.Print("Displacement: ", fn(t))
}

func GetUserInput(msg string, num *float64) {
	fmt.Print(msg, ": ")
	if _, err := fmt.Scanf("%f", num); err != nil {
		fmt.Print("Enter a valid value")
		return
	}
}
func GenDisplaceFn(a, vo, so float64) func (float64) float64 {
	return func (t float64) float64 {
		return 0.5 * a * t * t + vo * t + so
	}
}
