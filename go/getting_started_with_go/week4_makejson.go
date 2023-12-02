// write a program which prompts the user to first enter a name, and the enter an address
// Your program should create a map and add the name and address to the map using the keys "name and "address" respectively
// The program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	userInfo := make(map[string]string)
	
	fmt.Print("Enter your name: ")
	scanner.Scan()
	userInfo["name"] = scanner.Text()

	fmt.Print("Enter your address: ")
	scanner.Scan()
	userInfo["address"] = scanner.Text()

	jsonInfo, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("Error creating JSON:", err)
		return
	}

	// Print JSON object
	fmt.Println("JSON output:", string(jsonInfo))
}

