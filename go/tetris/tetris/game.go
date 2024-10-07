package tetris

const BOARD_HEIGHT = 16
const BOARD_WIDTH = 10

type game struct {
	board [][] int
}

func (g *game) GetBoard() [][]int {
	return g.board
}

func (g *game) init() {
	g.board = make([][]int, BOARD_HEIGHT)

	for i := 0; i < BOARD_HEIGHT; i++ {
		g.board[i] = make([]int, BOARD_WIDTH)
		
		for j := 0; j < BOARD_WIDTH; j++ {
			g.board[i][j] = 0
		}
	}
}

func New() *game {
	g := &game{}
	g.init()

	return g
}
