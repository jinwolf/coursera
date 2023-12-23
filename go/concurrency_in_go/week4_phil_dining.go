package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

type Chopstick struct { 
	sync.Mutex
}

type Philosopher struct {
	id int
	leftChopstick, rightChopstick *Chopstick
	eatCount int
}

func (p Philosopher) eat(permChan chan bool, wg *sync.WaitGroup) {
	for p.eatCount < 3 {
		<-permChan // Request permission to eat

		// Pick up chopsticks in random order
		if rand.Intn(2) == 0 {
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()
		} else {
			p.rightChopstick.Lock()
			p.leftChopstick.Lock()
		}

		fmt.Printf("%d: starting to eat \n", p.id)
		p.eatCount++
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Printf("%d: finishing eating \n", p.id)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		permChan <- true // Release permission
	}
	wg.Done()

}

func host(permChan chan bool) {
	for i := 0; i < 2; i++ {
		permChan <- true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	chopsticks := make([]*Chopstick, 5)

	for i :=0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			id: i + 1,
			leftChopstick: chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
		}
	}

	permChan := make(chan bool, 2)
	go host(permChan)

	var wg sync.WaitGroup
	for _, philosopher := range philosophers {
		wg.Add(1)
		go philosopher.eat(permChan, &wg)
	}
	wg.Wait()
}

