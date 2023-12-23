package main

import (
	"fmt"
	"sync"
)

var i int = 0
var mut sync.Mutex
var wg sync.WaitGroup

func inc() {
	mut.Lock()
	i = i + 1
	mut.Unlock()
	wg.Done()
}

func main() {
	//Mutex ensures mutual exclusion
	//Mutex uses binary semiphore

	//Lock() method puts the flag up
	//Lock() will create a mutually exclusive zone
	//Lock() will block the next go routine

	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(i)
}
