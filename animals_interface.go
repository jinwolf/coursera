package main

import (
	"fmt"
)

type AnimalInterface interface {
	Eat()
	Speak()
	Move()
}
type Animal struct {
	foodEaten        string
	locomotionMethod string
	spokenSound      string
}

func (animal Animal) Eat() {
	fmt.Println(animal.foodEaten)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotionMethod)
}

func (animal Animal) Speak() {
	fmt.Println(animal.spokenSound)
}

func main() {
	var cow, bird, snake Animal
	var animal AnimalInterface
	cow = Animal{"grass", "walk", "moo"}
	bird = Animal{"worms", "fly", "peep"}
	snake = Animal{"mice", "slither", "hsss"}

	mp := make(map[string]Animal)
	for {
		var command, name, commandType string
		fmt.Printf("> ")
		fmt.Scan(&command, &name, &commandType)
		if command == "newanimal" {
			switch commandType {
			case "cow":
				mp[name] = cow
			case "bird":
				mp[name] = bird
			case "snake":
				mp[name] = snake
			default:
				fmt.Printf("Invalid Animal Type::%s\n", commandType)
				continue
			}
			fmt.Println("Created it!")
		} else if command == "query" {
			_, ok := mp[name]
			if ok {
				animal = mp[name]
			} else {
				fmt.Printf("Animal %s does not exist. Please Create it first\n", name)
				continue
			}
			switch commandType {
			case "eat":
				animal.Eat()
			case "speak":
				animal.Speak()
			case "move":
				animal.Move()
			default:
				fmt.Printf("Invalid Action:: %s\n", commandType)
				continue
			}
		} else {
			fmt.Printf("Invalid Command:: %s\n", command)
		}

	}
}
