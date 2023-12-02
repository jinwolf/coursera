package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal definition
type Animal struct {
	food, locomotion, noise string
}

//Eat definition
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

//Move definition
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

//Speak definition
func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {

	var animal Animal
	for {
		fmt.Println("Choose an animal from this list {cow, bird, snake} and the information you need {eat, move, speak}: ->")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		
		animalName := strings.Split(scanner.Text(), " ")[0]
		information := strings.Split(scanner.Text(), " ")[1]

		switch animalName {
			case "cow":
				animal = Animal{"grass", "walk", "moo"}
			case "bird":
				animal = Animal{"worms", "fly", "peep"}
			case "snake":
				animal = Animal{"mice", "slither", "hsss"}
			default:
				fmt.Println("Please choose animals from the list")
				
		}

		if animalName == "cow" || animalName == "bird" || animalName == "snake"{
			switch information {
				case "eat": 
				animal.Eat();
			case "move":
				animal.Move();
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Please choose information from the list")	
				
			}
		}
	}
	
}