package main

import "fmt"

type AnimalType struct {
	Food string
	Locomotion string
	Sound string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func main() {
	animalTypes := map[string]AnimalType{
		"cow": AnimalType{"cow", "grass", "walk", "moo"},
		"bird": AnimalType{"bird", "worms", "fly", "peep"},
		"snake": AnimalType{"snake", "mice", "slither", "hsss"},
	}

	animals := map[string]AnimalType{}
		

	// Get user input
	for {
		var command, name, aniType string

		fmt.Print("> ")
		fmt.Scanln(&command, &name, &aniType)
		
		switch command {
			case "newanimal":
				// find the type
				definedType := animalTypes[aniType]
				animals[name] = definedType
				fmt.Println("Created it!")
			case "query":
				ani := animals[name]
				fmt.Println(ani)
				// call a method
				

			default:
				fmt.Println("Invalid command")
				continue
			}
		


	}
}
