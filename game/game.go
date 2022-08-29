package game

import (
	"fmt"

	"github.com/tamj0rd2/go-dots2/game/points"
)

type Game struct {
	grid          dots
	dotPerimeter  int
	width, height int
}

// by the end of the game, it's expected that every dot will be connected to its direct neighbours. This is 2 for
// dots at the edge of the board, or 4 for dots nearer to the middle.
const maxConnections = 4

func New(size int) *Game {
	dotPerimeter := size + 1
	dotCount := dotPerimeter * 2
	grid := make(dots, dotCount)

	for y := 0; y < dotPerimeter; y++ {
		for x := 0; x < dotPerimeter; x++ {
			grid[points.Coord{X: x, Y: y}.ID()] = make(map[points.Translation]bool, maxConnections)
		}
	}

	return &Game{
		grid:         grid,
		dotPerimeter: dotPerimeter,
		width:        size,
		height:       size,
	}
}

func (g *Game) Connect(a points.Coord, translation points.Translation) {
	coordToConnectTo := a.Translate(translation)
	if !coordToConnectTo.IsWithinBounds(g.dotPerimeter, g.dotPerimeter) {
		panic(fmt.Errorf("trying to connect %v to %v which would be out of bounds of the %dx%d grid", a, coordToConnectTo, g.width, g.height))
	}

	g.grid.get(a).connect(translation)
	g.grid.get(coordToConnectTo).connect(translation.Opposite())
}

func (g *Game) IsSquare(coordinate points.Coord) bool {
	topLeftDot := g.grid.get(coordinate)
	bottomRightDot := g.grid.get(coordinate.Translate(points.Down).Translate(points.Right))

	return topLeftDot.isConnected(points.Right) &&
		topLeftDot.isConnected(points.Down) &&
		bottomRightDot.isConnected(points.Left) &&
		bottomRightDot.isConnected(points.Up)
}

type dots map[points.ID]dot

func (c dots) get(coordinate points.Coord) dot {
	cnx, ok := c[coordinate.ID()]
	if !ok {
		panic(fmt.Errorf("unregistered coord %v", coordinate))
	}

	return cnx
}

type dot map[points.Translation]bool

func (c dot) connect(translation points.Translation) {
	c[translation] = true
}

func (c dot) isConnected(translation points.Translation) bool {
	return c[translation]
}
