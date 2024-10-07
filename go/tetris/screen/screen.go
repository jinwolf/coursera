package screen

import "fmt"

type gameScreen struct{}

func (g *gameScreen) RenderAsci(board [][]int) {
	fmt.Println("\n============")

	for _, e := range board {
		for _, num := range e {
			if num > 0 {
				fmt.Println("X")
			} else {
				fmt.Println(" ")
			}
		}
	}
	fmt.Println("")
}

func New() *gameScreen {
	return &gameScreen{} 
}

