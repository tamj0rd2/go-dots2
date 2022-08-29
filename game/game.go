package game

import (
	"fmt"

	"github.com/tamj0rd2/go-dots2/game/points"
)

type Game struct {
	grid                [][]dot
	width, height       int
	dotWidth, dotHeight int
}

// by the end of the game, it's expected that every dot will be connected to its direct neighbours. This is 2 for
// dots at the edge of the board, or 4 for dots nearer to the middle.
const maxConnections = 4

func New(size int) *Game {
	dotPerimeter := size + 1

	grid := make([][]dot, dotPerimeter)
	for y := 0; y < dotPerimeter; y++ {
		grid[y] = make([]dot, dotPerimeter)

		for x := 0; x < dotPerimeter; x++ {
			grid[y][x] = make(map[points.Translation]bool, maxConnections)
		}
	}

	return &Game{
		grid:      grid,
		dotWidth:  dotPerimeter,
		dotHeight: dotPerimeter,
		width:     size,
		height:    size,
	}
}

func (g *Game) Connect(a points.Coord, translation points.Translation) {
	g.getDot(a).connect(translation)
	g.getDot(a.Translate(translation)).connect(translation.Opposite())
}

func (g *Game) getDot(coord points.Coord) dot {
	if !coord.IsWithinBounds(g.dotWidth, g.dotHeight) {
		panic(fmt.Errorf("dot %v is out of bounds of the %dx%d grid", coord, g.width, g.height))
	}
	return g.grid[coord.Y][coord.X]
}

type dot map[points.Translation]bool

func (c dot) connect(translation points.Translation) {
	c[translation] = true
}

func (c dot) isConnected(translation points.Translation) bool {
	return c[translation]
}
