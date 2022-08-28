package game

import (
	"strings"

	"github.com/tamj0rd2/go-dots2/game/points"
)

const bottomCornerChar = "˙"
const connectedSideChar = "|"

type Game struct {
	grid dots
	size int
}

func New(size int) *Game {
	dotCount := size * 4
	return &Game{grid: make(dots, dotCount), size: size}
}

func (g *Game) Grid() string {
	var lastIndex = g.size - 1
	var out []string

	for y := 0; y < g.size; y++ {
		var (
			tops    []string
			mids    []string
			bottoms []string

			topCornerChar = ":"
		)
		if y == 0 {
			topCornerChar = "."
		}

		for x := 0; x < g.size; x++ {
			var (
				top, mid, bottom string
			)
			top = topCornerChar + strings.Repeat("-", 3)
			mid = connectedSideChar + strings.Repeat(" ", 3)
			bottom = bottomCornerChar + strings.Repeat("-", 3)

			if x == lastIndex {
				top += topCornerChar
				mid += connectedSideChar
				bottom += bottomCornerChar
			}

			tops = append(tops, top)
			mids = append(mids, mid)
			bottoms = append(bottoms, bottom)
		}

		out = append(out, strings.Join(tops, ""), strings.Join(mids, ""))
		if y == lastIndex {
			out = append(out, strings.Join(bottoms, ""))
		}
	}

	return strings.Join(out, "\n")
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
