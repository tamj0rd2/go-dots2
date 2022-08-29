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
	// by the end of the game, it's expected that every dot will be connected to its direct neighbours. This is 2 for
	// dots at the edge of the board, or 4 for dots nearer to the middle.
	const maxConnections = 4

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

func (g *grid) drawLine(from points.Coords, translation points.Translation) {
	to := from.Translate(translation)
	g.getDot(from).connect(translation)
	g.getDot(to).connect(translation.Opposite())
}

func (g *grid) getDot(coords points.Coords) dot {
	if !coords.IsWithinBounds(g.dotWidth, g.dotHeight) {
		panic(fmt.Errorf("dot %v is out of bounds of the %dx%d grid", coords, g.width, g.height))
	}
	return g.grid[coords.Y][coords.X]
}

func (g *grid) newSquaresCompleted(a, b points.Coords) int {
	newlyCompletedSquares := 0
	start, end, orientation := orderByDistanceFromOrigin(a, b)

	switch orientation {
	case vertical:
		if g.getDot(start).isConnected(points.Left) && g.getDot(end).isConnected(points.Left) {
			if g.getDot(start.Translate(points.Left)).isConnected(points.Down) {
				newlyCompletedSquares += 1
			}
		}

		if g.getDot(start).isConnected(points.Right) && g.getDot(end).isConnected(points.Right) {
			if g.getDot(start.Translate(points.Right)).isConnected(points.Down) {
				newlyCompletedSquares += 1
			}
		}
	case horizontal:
		if g.getDot(start).isConnected(points.Up) && g.getDot(end).isConnected(points.Up) {
			if g.getDot(start.Translate(points.Up)).isConnected(points.Right) {
				newlyCompletedSquares += 1
			}
		}

		if g.getDot(start).isConnected(points.Down) && g.getDot(end).isConnected(points.Down) {
			if g.getDot(start.Translate(points.Down)).isConnected(points.Right) {
				newlyCompletedSquares += 1
			}
		}
	default:
		panic(fmt.Errorf("unrecognised orientation %v", orientation))
	}

	return newlyCompletedSquares
}

func orderByDistanceFromOrigin(a, b points.Coords) (points.Coords, points.Coords, orientation) {
	if a.Y == b.Y {
		if a.X-b.X > 0 {
			return b, a, horizontal
		}
		return a, b, horizontal
	}

	if a.X == b.X {
		if a.Y-b.Y > 0 {
			return b, a, vertical
		}
		return a, b, vertical
	}

	panic(fmt.Errorf("the given points do not form a line: %s - %s", a.String(), b.String()))
}

type dot map[points.Translation]bool

func (c dot) connect(translation points.Translation) {
	c[translation] = true
}

func (c dot) isConnected(translation points.Translation) bool {
	return c[translation]
}

type orientation string

const (
	vertical   orientation = "vertical"
	horizontal orientation = "horizontal"
)
