package game

import (
	"fmt"

	"github.com/tamj0rd2/go-dots2/game/points"
)

type grid struct {
	grid                [][]dot
	width, height       int
	dotWidth, dotHeight int
}

func newGrid(width, height int) *grid {
	dotWidth := width + 1
	dotHeight := height + 1

	dots := make([][]dot, dotHeight)
	for y := 0; y < dotHeight; y++ {
		dots[y] = make([]dot, dotWidth)

		for x := 0; x < dotWidth; x++ {
			dots[y][x] = make(map[points.Translation]bool, maxConnections)
		}
	}

	return &grid{
		grid:      dots,
		width:     width,
		height:    height,
		dotWidth:  dotWidth,
		dotHeight: dotHeight,
	}
}

func (g *grid) getDot(coords points.Coords) dot {
	if !coords.IsWithinBounds(g.dotWidth, g.dotHeight) {
		panic(fmt.Errorf("dot %v is out of bounds of the %dx%d grid", coords, g.width, g.height))
	}
	return g.grid[coords.Y][coords.X]
}
