package main

import "fmt"

type AnimalType struct {
	Food string
	Locomotion string
	Sound string
}

type Cow struct {
	AnimalType
}

type Bird struct {
	AnimalType
}

type Snake struct {
	AnimalType
}

func NewCow() Cow {
	return Cow {
		AnimalType {
			Food: "grass",
			Locomotion: "walk",
			Sound: "moo",
		},
	}
}

func NewBird() Bird {
	return Bird {
		AnimalType {
			Food: "worms", 
			Locomotion: "fly",
			Sound: "peep",
		},
	}
}

func NewSnake() Snake {
	return Snake {
		AnimalType {
			Food: "mice",
			Locomotion: "slither",
			Sound: "hsss",
		},
	}
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

func (a AnimalType) Eat() {
	fmt.Println(a.Food)
}

func(a AnimalType) Move() {
	fmt.Println(a.Locomotion)
}

func(a AnimalType) Speak() {
	fmt.Println(a.Sound)
}

func CreateAnimal(kind string) Animal {
	switch kind {
	case "cow":
		return NewCow()
	case "bird":
		return NewBird()
	case "snake":
		return NewSnake()
	default:
		fmt.Println("Invalid Animal Type")
		return nil
	}
}

func AnimalAction(animal Animal, kind string) {
	switch kind {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Print("Invalid action")
		fmt.Scanln() // clear the buffer
	}
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command, name, kind string
		fmt.Println("- newanimal [name] [type]")
		fmt.Println("- query [name] [action]")
		fmt.Print("> ")
		fmt.Scan(&command, &name, &kind)

		switch command {
		case "newanimal":
			animals[name] = CreateAnimal(kind)
			fmt.Println("Created it!")
		case "query":
			animal, ok := animals[name]
			if !ok {
				fmt.Println("No such animal")
				continue
			}

			AnimalAction(animal, kind)
		default:
			fmt.Println("Invalid command")
			fmt.Scanln() // clear buffer
		}

	}
}
