package game

import (
	"strings"

	"github.com/tamj0rd2/go-dots2/game/points"
)

const (
	middleCorner      = ":"
	topCornerChar     = "."
	bottomCornerChar  = "˙"
	connectedSideChar = "|"
	connectedTopChar  = "-"
	charCount         = 3
)

func (g *grid) Grid() string {
	var out []string
	lastYIndex := g.height

	for squareY := 0; squareY <= g.height; squareY++ {
		var (
			tops  []string
			sides []string
		)

		for squareX := 0; squareX < g.width; squareX++ {
			var (
				isLastXIndex  = squareX == g.width-1
				topLeftCorner = g.getDot(points.Coords{X: squareX, Y: squareY})
			)

			cornerChar := getCornerChar(squareY, lastYIndex)
			top := getHorizontalRepresentation(topLeftCorner.dot, cornerChar)
			side := getVerticalRepresentation(topLeftCorner.dot) + strings.Repeat(" ", charCount)

			if isLastXIndex {
				top += cornerChar
				topRightCorner := g.getDot(topLeftCorner.Coords.Translate(points.Right))
				side += getVerticalRepresentation(topRightCorner.dot)
			}

			tops = append(tops, top)
			sides = append(sides, side)
		}

		out = append(out, strings.Join(tops, ""), strings.Join(sides, ""))
	}

	return strings.Join(out, "\n")
}

func getHorizontalRepresentation(topLeftCorner dot, cornerChar string) string {
	topChar := " "
	if topLeftCorner.isConnected(points.Right) {
		topChar = connectedTopChar
	}
	return cornerChar + strings.Repeat(topChar, charCount)
}

func getVerticalRepresentation(dot dot) string {
	if dot.isConnected(points.Down) {
		return connectedSideChar
	}
	return " "
}

func getCornerChar(y, lastYIndex int) string {
	if y == 0 {
		return topCornerChar
	}

	if y == lastYIndex {
		return bottomCornerChar
	}

	return middleCorner
}
