// read info from a file and represents it in a slice of structs
// a text file which contains a series of names (first and last names separated by a single space on the line)
// Create a name struct: fname and lname: string of size 20
// 1. prompt the name of the text file
// 2. read each line and create a struct and add it to a slice
// 3. print out the struct
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var fileName string
	var names []name

	fmt.Print("Enter the name of the text file: ")
	fmt.Scan(&fileName)

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 2)

		// check for correct format (first and last name)
		if len(parts) == 2 {
			var person name
			person.fname = parts[0]
			person.lname = parts[1]
			names = append(names, person)
		}
	}

	// check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Print names
	for _, person := range names {
		fmt.Printf("First Name: %-20s Last Name: %-20s\n", person.fname, person.lname)
	}
}
