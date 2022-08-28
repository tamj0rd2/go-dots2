package game

import (
	"github.com/tamj0rd2/go-dots2/game/points"
)

type Game struct {
	grid dots
}

func New(size int) *Game {
	dotCount := size * 4
	return &Game{grid: make(dots, dotCount)}
}

func (g *Game) Connect(a points.Coord, translation points.Translation) {
	g.grid.get(a).connect(translation)
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
		cnx = dot{}
		c[coordinate.ID()] = cnx
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
