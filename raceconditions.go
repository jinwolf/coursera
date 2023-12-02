package main

import (
	"fmt"
	"time"
)

var i int

/*
A race condition is when two goroutines are reading or writing
to the same variable at the same time. This would be a race
condition because the end result could be different each time 
based on the order of the goroutines completing.

A race condition occurs here because multiple goroutines are
reading and writing to the same variable at the same time.
*/
func main() {
	go subOne()
	go subTwo()
	time.Sleep(1 * time.Second)
}

func subOne() {
	i++
	fmt.Println("Race conditon", i)
}

func subTwo() {
	i--
	fmt.Println("Race conditon", i)
}
