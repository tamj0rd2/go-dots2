package game_test

import (
	"testing"

	"github.com/tamj0rd2/go-dots2/game"
	"github.com/tamj0rd2/go-dots2/specifications"
)

func TestPlayingATwoPlayerGame(t *testing.T) {
	specifications.Game{
		NewSubject: func(size int) specifications.GameDriver {
			return game.New(size)
		},
	}.Test(t)
}
