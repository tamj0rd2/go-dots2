package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type GameDriver interface{}

type Game struct {
	NewSubject func() GameDriver
}

func (spec Game) Test(t *testing.T) {
	t.Run("playing a two player game", func(t *testing.T) {
		game := spec.NewSubject()

		assert.NotZero(t, game)
	})
}
