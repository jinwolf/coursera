package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

var counter int = 0

func increment() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	temp := counter // temp could be 0 or 1
	temp++ // temp could 1 or 2

	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	counter = temp
}

/*
This program creates two goroutines, each running the function "increment."
This function increments a counter with random time delays to amplify the race condition.
Since there is no guarantee which goroutine will start or finish first, the final
counter value is nondeterministic.
I put this in a loop to see the different final values in each loop
*/
func main() {
	rand.Seed(time.Now().UnixNano())

	for i:=0; i < 10; i++ {
		// resets counter
		counter = 0


		var wg sync.WaitGroup

		wg.Add(2)
		go func() {
			increment()
			wg.Done()
		}()

		go func() {
			increment()
			wg.Done()
		}()

		wg.Wait()
	
		// counter will be either 1 or 2
		fmt.Printf("Run %d: Final counter value: %d\n",i+1,  counter)
	}
}
