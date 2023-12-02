package main

import "fmt"

type IAnimal interface {
	Eat()
	Move()
	Speak()
}

type Dog struct {
	Name string
	Food string
	Locomotion string
	Sound string
}
	
func (d Dog) Eat() {
	fmt.Println(d.Name, "is a dog and eats", d.Food)
}

func (d Dog) Move() {
	fmt.Println(d.Name, "is a dog and move by", d.Locomotion)
}

func (d Dog) Speak() {
	fmt.Println(d.Name, "is a dog and speak like", d.Sound)
}

func RunEat(animal IAnimal) {
	animal.Eat()
}

func main() {
	iAnimal := Dog { "Bug", "Chicken", "walk", "Woof Woof",}
	RunEat(iAnimal)
}
