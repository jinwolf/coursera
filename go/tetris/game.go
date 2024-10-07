package tetris

type game struct {
	board [][] int
}

func New() *game {
	return &game{}
}
