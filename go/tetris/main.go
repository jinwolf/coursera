package main

import (
	"fmt"
	"github.com/jinwolf/tetris/tetris"
	"github.com/jinwolf/tetris/screen"
)

func main() {
	fmt.Println("Tetris")
	game := tetris.New()
	scr := screen.New()

	scr.RenderAsci(game.GetBoard())
}
