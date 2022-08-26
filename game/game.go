package game

import "github.com/tamj0rd2/go-dots2/game/board"

type Game struct {
	dots
}

func New(size int) *Game {
	dotCount := size * 4
	return &Game{dots: make(dots, dotCount)}
}

func (g *Game) ConnectDots(a, b board.Coordinate) {
	g.dots.get(a).connectTo(b)
	g.dots.get(b).connectTo(a)
}

func (g Game) IsSquareOwned(coordinate board.Coordinate) bool {
	leftEdge := coordinate.X
	rightEdge := coordinate.X + 1
	topEdge := coordinate.Y
	bottomEdge := coordinate.Y + 1

	topLeft := board.Coordinate{X: leftEdge, Y: topEdge}
	topRight := board.Coordinate{X: rightEdge, Y: topEdge}
	bottomRight := board.Coordinate{X: rightEdge, Y: bottomEdge}
	bottomLeft := board.Coordinate{X: leftEdge, Y: bottomEdge}

	return g.dots.areDotsConnected(topLeft, topRight) &&
		g.dots.areDotsConnected(topRight, bottomRight) &&
		g.dots.areDotsConnected(bottomRight, bottomLeft) &&
		g.dots.areDotsConnected(bottomLeft, topLeft)
}

type connections map[string]bool

func (c connections) isConnectedTo(dot board.Coordinate) bool {
	return c[dot.ID()]
}

func (c connections) connectTo(b board.Coordinate) {
	c[b.ID()] = true
}

type dots map[string]connections

func (c dots) get(coordinate board.Coordinate) connections {
	cnx, ok := c[coordinate.ID()]
	if !ok {
		cnx = connections{}
		c[coordinate.ID()] = cnx
	}

	return cnx
}

func (c dots) areDotsConnected(a, b board.Coordinate) bool {
	isAConnected := c.get(a).isConnectedTo(b)
	isBConnected := c.get(b).isConnectedTo(a)

	if isAConnected != isBConnected {
		panic("connections state of a and b are out of sync")
	}

	return isAConnected && isBConnected
}
