package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/tamj0rd2/go-dots2/game/board"
)

type GameDriver interface {
	Connect(dot board.Coordinate, position board.Translation)
	IsSquare(coordinate board.Coordinate) bool
}

type Game struct {
	NewSubject func(size int) GameDriver
}

func (spec Game) Test(t *testing.T) {
	t.Run("connecting dots in a single player game", func(t *testing.T) {
		game := spec.NewSubject(1)
		assert.False(t, game.IsSquare(board.Coordinate{X: 0, Y: 0}))

		topLeft := board.Coordinate{X: 0, Y: 0}
		bottomRight := board.Coordinate{X: 1, Y: 1}
		game.Connect(topLeft, board.Right)
		game.Connect(topLeft, board.Down)
		game.Connect(bottomRight, board.Up)
		game.Connect(bottomRight, board.Left)

		assert.True(t, game.IsSquare(board.Coordinate{X: 0, Y: 0}), "expected square to be owned")
	})
}
