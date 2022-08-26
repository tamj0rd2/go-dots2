package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/tamj0rd2/go-dots2/game/board"
)

type GameDriver interface {
	ConnectDots(a, b board.Coordinate)
	IsSquareOwned(coordinate board.Coordinate) bool
}

type Game struct {
	NewSubject func(size int) GameDriver
}

func (spec Game) Test(t *testing.T) {
	t.Run("connecting dots in a single player game", func(t *testing.T) {
		game := spec.NewSubject(1)
		assert.False(t, game.IsSquareOwned(board.Coordinate{X: 0, Y: 0}))

		topLeft := board.Coordinate{X: 0, Y: 0}
		topRight := board.Coordinate{X: 1, Y: 0}
		bottomLeft := board.Coordinate{X: 0, Y: 1}
		bottomRight := board.Coordinate{X: 1, Y: 1}

		game.ConnectDots(topLeft, topRight)
		game.ConnectDots(topRight, bottomRight)
		game.ConnectDots(bottomRight, bottomLeft)
		game.ConnectDots(bottomLeft, topLeft)

		assert.True(t, game.IsSquareOwned(board.Coordinate{X: 0, Y: 0}), "expected square to be owned")
	})
}
