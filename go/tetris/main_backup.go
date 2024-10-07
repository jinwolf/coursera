package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	red := color.RGBA{R: 255, G:0, B: 0, A: 255}
	ebitenutil.DrawRect(screen, 50, 50, 30, 30,red)
}
	func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
		return 320, 240
	}

	func main() {
		ebiten.SetWindowSize(640, 480)
		ebiten.SetWindowTitle("Tetris")
		if err := ebiten.RunGame(&Game{}); err != nil {
			log.Fatal(err)
		}
	}
