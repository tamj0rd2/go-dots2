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

func (g *grid) getDot(coords points.Coords) dotWithCoords {
	if !coords.IsWithinBounds(g.dotWidth, g.dotHeight) {
		panic(fmt.Errorf("dot %v is out of bounds of the %dx%d grid", coords, g.width, g.height))
	}

	return dotWithCoords{
		dot:    g.grid[coords.Y][coords.X],
		Coords: coords,
	}
}

func (g *grid) newSquaresCompleted(a, b points.Coords) int {
	var (
		newlyCompletedSquares   = 0
		start, end, orientation = g.orderByDistanceFromOrigin(g.getDot(a), g.getDot(b))

		extremeEdgeOfSquare         points.Translation
		directionsNewSquaresCanBeIn [2]points.Translation
	)

	switch orientation {
	case vertical:
		directionsNewSquaresCanBeIn = [2]points.Translation{points.Left, points.Right}
		extremeEdgeOfSquare = points.Down
	case horizontal:
		directionsNewSquaresCanBeIn = [2]points.Translation{points.Up, points.Down}
		extremeEdgeOfSquare = points.Right
	default:
		panic(fmt.Errorf("unrecognised orientation %v", orientation))
	}

	for _, direction := range directionsNewSquaresCanBeIn {
		hasPerpendicularLines := start.isConnected(direction) && end.isConnected(direction)
		if !hasPerpendicularLines {
			continue
		}

		furthestCorner := g.getDot(start.Translate(direction))
		if furthestCorner.isConnected(extremeEdgeOfSquare) {
			newlyCompletedSquares += 1
		}
	}

	return newlyCompletedSquares
}

func (g *grid) orderByDistanceFromOrigin(a, b dotWithCoords) (dotWithCoords, dotWithCoords, orientation) {
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

type orientation string

const (
	vertical   orientation = "vertical"
	horizontal orientation = "horizontal"
)

type dotWithCoords struct {
	dot
	points.Coords
}

type dot map[points.Translation]bool

func (c dot) connect(translation points.Translation) {
	c[translation] = true
}

func (c dot) isConnected(translation points.Translation) bool {
	return c[translation]
}
