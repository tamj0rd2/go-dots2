package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/tamj0rd2/go-dots2/game/points"
)

type GameDriver interface {
	Connect(dot points.Coord, position points.Translation)
	IsSquare(coordinate points.Coord) bool
	// Grid() string
}

type Game struct {
	NewSubject func(size int) GameDriver
}

func (spec Game) Test(t *testing.T) {
	t.Run("connecting dots in a single player game", func(t *testing.T) {
		game := spec.NewSubject(1)
		assert.False(t, game.IsSquare(points.Coord{X: 0, Y: 0}))

		topLeft := points.Coord{X: 0, Y: 0}
		bottomRight := points.Coord{X: 1, Y: 1}
		game.Connect(topLeft, points.Right)
		game.Connect(topLeft, points.Down)
		game.Connect(bottomRight, points.Up)
		game.Connect(bottomRight, points.Left)

		assert.True(t, game.IsSquare(points.Coord{X: 0, Y: 0}), "expected square to be owned")

		/*bleh := `
		 _  . . . .
		|_| . . . .
		. . . . . .`
		*/
	})
}
