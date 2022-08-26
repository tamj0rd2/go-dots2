package game

import "github.com/tamj0rd2/go-dots2/game/board"

type Game struct {
	dots
}

func New(size int) *Game {
	dotCount := size * 4
	return &Game{dots: make(dots, dotCount)}
}

func (g *Game) Connect(dot board.Coordinate, position board.Translation) {
	dotToConnect := dot.Translate(position)
	g.dots.get(dot).connectTo(dotToConnect)
	g.dots.get(dotToConnect).connectTo(dot)
}

func (g Game) IsSquare(coordinate board.Coordinate) bool {
	topLeft := board.Coordinate{X: coordinate.X, Y: coordinate.Y}
	topRight := topLeft.Translate(board.Right)
	bottomRight := topRight.Translate(board.Down)
	bottomLeft := bottomRight.Translate(board.Left)

	return g.dots.areConnected(topLeft, topRight) &&
		g.dots.areConnected(topRight, bottomRight) &&
		g.dots.areConnected(bottomRight, bottomLeft) &&
		g.dots.areConnected(bottomLeft, topLeft)
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

func (c dots) areConnected(a, b board.Coordinate) bool {
	isAConnected := c.get(a).isConnectedTo(b)
	isBConnected := c.get(b).isConnectedTo(a)

	if isAConnected != isBConnected {
		panic("connections state of a and b are out of sync")
	}

	return isAConnected && isBConnected
}
