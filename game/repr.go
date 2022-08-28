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

func (g *Game) Grid() string {
	// TODO: this probably shouldn't live on Game. Seems like a responsibility of the presentation layer
	var out []string
	lastYIndex := g.height

	for squareY := 0; squareY <= g.height; squareY++ {
		var (
			tops  []string
			sides []string
		)

		for squareX := 0; squareX < g.width; squareX++ {
			isLastXIndex := squareX == g.width-1
			topLeftCorner := g.grid.get(points.Coord{X: squareX, Y: squareY})

			cornerChar := getCornerChar(squareY, lastYIndex)
			top := getHorizontal(topLeftCorner, cornerChar)
			side := getVertical(topLeftCorner)

			if isLastXIndex {
				top += cornerChar

				topRightCorner := g.grid.get(points.Coord{X: squareX + 1, Y: squareY})
				if topRightCorner.isConnected(points.Down) {
					side += connectedSideChar
				} else {
					side += " "
				}
			}

			tops = append(tops, top)
			sides = append(sides, side)
		}

		out = append(out, strings.Join(tops, ""), strings.Join(sides, ""))
	}

	return strings.Join(out, "\n")
}

func getHorizontal(topLeftCorner dot, cornerChar string) string {
	topChar := " "
	if topLeftCorner.isConnected(points.Right) {
		topChar = connectedTopChar
	}
	return cornerChar + strings.Repeat(topChar, charCount)
}

func getVertical(topLeftCorner dot) string {
	sideChar := " "
	if topLeftCorner.isConnected(points.Down) {
		sideChar = connectedSideChar
	}
	return sideChar + strings.Repeat(" ", charCount)
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
