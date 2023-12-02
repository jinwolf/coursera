package main

import "fmt"

func main() {
	fmt.Println("NUMBERS.")
	var number int
	fmt.Printf("Enter a decimal ...")
	fmt.Scan(&number)
	fmt.Print("The number you put in was : ", number)
}
