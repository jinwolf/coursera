package main

import (
	"fmt"
	"sync"
)

var on sync.Once
var wg sync.WaitGroup

func setup() {
	fmt.Println("Init")
}

func dostuff() {
	on.Do(setup)
	fmt.Println("hello")
	wg.Done()
}

func main() {
	wg.Add(3)
        go dostuff()
        go dostuff()
        go dostuff()
	wg.Wait()
}

