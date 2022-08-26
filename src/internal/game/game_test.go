package game_test

import (
	"testing"

	"github.com/tamj0rd2/go-dots2/specifications"
	"github.com/tamj0rd2/go-dots2/src/internal/game"
)

func TestPlayingATwoPlayerGame(t *testing.T) {
	specifications.Game{
		NewSubject: func() specifications.GameDriver {
			return game.New()
		},
	}.Test(t)
}
