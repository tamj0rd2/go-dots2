package game

import (
	"github.com/tamj0rd2/go-dots2/game/points"
)

type Game struct {
	grid
}

func New(size int) *Game {
	dotCount := size * 4
	return &Game{grid: make(grid, dotCount)}
}

func (g *Game) Connect(a points.Coord, position points.Translation) {
	b := a.Translate(position)
	g.grid.get(a).connectTo(b)
	g.grid.get(b).connectTo(a)
}

func (g *Game) IsSquare(coordinate points.Coord) bool {
	topLeft := points.Coord{X: coordinate.X, Y: coordinate.Y}
	topRight := topLeft.Translate(points.Right)
	bottomRight := topRight.Translate(points.Down)
	bottomLeft := bottomRight.Translate(points.Left)

	return g.grid.areConnected(topLeft, topRight) &&
		g.grid.areConnected(topRight, bottomRight) &&
		g.grid.areConnected(bottomRight, bottomLeft) &&
		g.grid.areConnected(bottomLeft, topLeft)
}

type connections map[string]bool

func (c connections) isConnectedTo(dot points.Coord) bool {
	return c[dot.ID()]
}

func (c connections) connectTo(b points.Coord) {
	c[b.ID()] = true
}

type grid map[string]connections

func (c grid) get(coordinate points.Coord) connections {
	cnx, ok := c[coordinate.ID()]
	if !ok {
		cnx = connections{}
		c[coordinate.ID()] = cnx
	}

	return cnx
}

func (c grid) areConnected(a, b points.Coord) bool {
	isAConnected := c.get(a).isConnectedTo(b)
	isBConnected := c.get(b).isConnectedTo(a)

	if isAConnected != isBConnected {
		panic("connections state of a and b are out of sync")
	}

	return isAConnected && isBConnected
}
