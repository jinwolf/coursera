package main

import "fmt"

/*
Race condition - when multiple threads are trying to work with the same variable (changing it value).
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
It means that result can be deferent every time
*/

func increment(v *int) {
	*v++
	fmt.Println(*v)
}

func decrement(v *int) {
	*v--
	fmt.Println(*v)
}
func main() {
	i := 0
	go increment(&i)
	go decrement(&i)
	i++
	fmt.Println()
}
