// Each animal has a name, type:cow, bird or snake
// The user can create a new animal of one of the three types
// The user can request (query) info about an animal that already exists
// the animal types are restricted
package main

import "fmt"

type AnimalType struct {
	Animal string
	Food string
	Locomotion string
	Sound string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Animal struct {
	name string
	aniType AnimalType
}

func main() {
	animalTypes := map[string]AnimalType{
		"cow": AnimalType{"cow", "grass", "walk", "moo"},
		"bird": AnimalType{"bird", "worms", "fly", "peep"},
		"snake": AnimalType{"snake", "mice", "slither", "hsss"},
	}

	animals := map[string]Animal{}
		

	// Get user input
	for {
		var command, name, aniType string

		fmt.Print("> ")
		fmt.Scanln(&command, &name, &aniType)
		
		switch command {
			case "newanimal":
				// find the type
				definedType := animalTypes[aniType]
				animals[name] = Animal{name, definedType}
				fmt.Println("Created it!")
			case "query":
				ani := animals[name]
				fmt.Println(ani)

			default:
				fmt.Println("Invalid command")
				continue
			}
		


	}
}
