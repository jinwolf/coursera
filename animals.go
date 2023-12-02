package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	Food string
	Locomotion string
	Sound string
}

func(a Animal) Eat() string {
	return a.Food
}

func(a Animal) Move() string {
	return a.Locomotion
}

func(a Animal) Speak() string {
	return a.Sound
}

func main() {
	animals := map[string]Animal{
		"cow": Animal{"grass", "walk", "moo"},
		"bird": Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	// Get user input
	for {
		fmt.Print("> ")
		var animalName, action string
		fmt.Scanln(&animalName, &action)
		
		animal, ok := animals[animalName]
		if !ok {
			fmt.Println("Animal not found")
			continue
		}

		var attr string
		
		switch strings.ToLower(action) {
			case "eat":
				attr = animal.Eat()
			case "move":
				attr =  animal.Move()
			case "speak":
				attr = animal.Speak()
			default:
				fmt.Println("Action not recognized")
		}

		fmt.Println(attr)
	}
}
