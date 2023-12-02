package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ food, locomotion, noise string }

type Bird struct{ food, locomotion, noise string }

type Snake struct{ food, locomotion, noise string }

func (cow Cow) Eat() {
	fmt.Println("\nIt eats", cow.food)
}

func (bird Bird) Eat() {
	fmt.Println("\nIt eats", bird.food)
}

func (snake Snake) Eat() {
	fmt.Println("\nIt eats", snake.food)
}

func (cow Cow) Move() {
	fmt.Println("\nIt", cow.locomotion)
}

func (bird Bird) Move() {
	fmt.Println("\nIt", bird.locomotion)
}

func (snake Snake) Move() {
	fmt.Println("\nIt", snake.locomotion)
}

func (cow Cow) Speak() {
	fmt.Println("\nIt", cow.noise)
}

func (bird Bird) Speak() {
	fmt.Println("\nIt", bird.noise)
}

func (snake Snake) Speak() {
	fmt.Println("\nIt", snake.noise)
}

func main() {

	var input string
	var animalName string

	createAnimal := make(map[string]string)
	var choice int

	for {
		fmt.Printf("\n> 1. Create New Animal \n> 2. Query Animal \n> 3. Exit")
		fmt.Println("\nEnter your choice number")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("\nEnter Animal Name and Animal Type [cow/bird/snake] (separated by comma)")
			fmt.Scanln(&input)

			values := strings.Split(input, ",")
			animalName = values[0]
			animalType := values[1]

			_, check := createAnimal[animalName]
			if !check {
				createAnimal[strings.ToLower(animalName)] = strings.ToLower(animalType)
				fmt.Println("Animal Created!!!", animalName+"-"+animalType)
			} else {
				fmt.Println("Animal with this type already created")
			}
		case 2:
			fmt.Println("Enter Animal Name and Animal Query [food/locomotion/Speak] (separated by comma)")
			fmt.Scanln(&input)

			values := strings.Split(input, ",")
			animalName = values[0]
			info := values[1]

			queryAnimal(strings.ToLower(animalName), strings.ToLower(info), createAnimal)
		case 3:
			return
		default:
			fmt.Println("Invalid choice pls try again")
			return
		}
	}
}

func queryAnimal(animalName string, info string, createdAnimal map[string]string) {

	switch strings.ToLower(createdAnimal[animalName]) {

	case "cow":

		cow := Cow{"grass", "walk", "moo"}

		switch info {
		case "food":
			cow.Eat()
		case "locomotion":
			cow.Move()
		case "speak":
			cow.Speak()
		default:
			fmt.Printf("Invalid Query")
			return
		}
	case "bird":

		bird := Bird{"worms", "fly", "peep"}

		switch info {
		case "food":
			bird.Eat()
		case "locomotion":
			bird.Move()
		case "speak":
			bird.Speak()
		default:
			fmt.Printf("Invalid Query pls try again")
			return
		}
	case "snake":

		snake := Snake{"mice", "slither", "hsss"}

		switch info {
		case "food":
			snake.Eat()
		case "locomotion":
			snake.Move()
		case "speak":
			snake.Speak()
		default:
			fmt.Printf("Invalid Query pls try again")
			return
		}
	default:
		fmt.Printf("No such animal found")
		return
	}
}
