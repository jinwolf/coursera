package main

import ( "fmt"
)

type Animal struct {
      food        string
      locomotion  string
      noise       string
}


func (a Animal) Eat() {
              fmt.Println(a.food)
}

func (a Animal) Move() {
              fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
              fmt.Println(a.noise)
}


func main() {
    cow   := Animal{ "grass", "walk", "moo"}
    bird  := Animal{ "worms", "fly", "peep"}
    snake := Animal{ "mice", "slither", "hsss"}
    var animal string
    var what   string
    for {
       fmt.Printf(">")
       n, _ := fmt.Scanf("%s %s", &animal, &what)
       if n == 2 {
       switch  what{
       case "eat":
           if animal == "cow" {
               cow.Eat()
           } else if animal == "bird" {
               bird.Eat()
           } else if animal == "snake" {
               snake.Eat()
           } else {
               panic("unknown animal")
           }
           animal = ""
       case "move":
           if animal == "cow" {
               cow.Move()
           } else if animal == "bird" {
               bird.Move()
           } else if animal == "snake" {
               snake.Move()
           } else {
               panic("unknown animal")
           }
           animal = ""
       case "speak":
           if animal == "cow" {
               cow.Speak()
           } else if animal == "bird" {
               bird.Speak()
           } else if animal == "snake" {
               snake.Speak()
           } else {
               panic("unknown animal")
           }
           animal = ""
       }
    } else {
        continue
    }
}
}
