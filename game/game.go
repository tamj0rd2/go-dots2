package game

import (
	"github.com/tamj0rd2/go-dots2/game/points"
)

type Game struct {
	*grid
	width, height       int
	dotWidth, dotHeight int
}

// by the end of the game, it's expected that every dot will be connected to its direct neighbours. This is 2 for
// dots at the edge of the board, or 4 for dots nearer to the middle.
const maxConnections = 4

func New(size int) *Game {
	dotPerimeter := size + 1

	return &Game{
		grid:      newGrid(size, size),
		dotWidth:  dotPerimeter,
		dotHeight: dotPerimeter,
		width:     size,
		height:    size,
	}
}

func (g *Game) DrawLine(from points.Coords, translation points.Translation) {
	to := from.Translate(translation)
	g.grid.getDot(from).connect(translation)
	g.grid.getDot(to).connect(translation.Opposite())
}

type dot map[points.Translation]bool

func (c dot) connect(translation points.Translation) {
	c[translation] = true
}

func (c dot) isConnected(translation points.Translation) bool {
	return c[translation]
}
