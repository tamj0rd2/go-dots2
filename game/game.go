package game

import "github.com/tamj0rd2/go-dots2/game/points"

type Game struct {
	*grid
	*scoring
}

func New(size int) *Game {
	return &Game{
		grid:    newGrid(size, size),
		scoring: &scoring{},
	}
}

func (g *Game) DrawLine(from points.Coords, translation points.Translation) {
	g.grid.drawLine(from, translation)
	newSquaresCompleted := g.grid.newSquaresCompleted(from, from.Translate(translation))
	g.scoring.RecordNewSquaresCompleted(newSquaresCompleted)
}

