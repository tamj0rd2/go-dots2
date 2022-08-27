package game_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/tamj0rd2/go-dots2/game"
	"github.com/tamj0rd2/go-dots2/game/points"
	"github.com/tamj0rd2/go-dots2/specifications"
)

func TestPlayingATwoPlayerGame(t *testing.T) {
	t.Run("connecting dots to make a square", func(t *testing.T) {
		for _, tc := range []struct {
			name        string
			connectDots func(g *game.Game)
			isSquare    bool
		}{
			{
				name:        "no lines is not a square",
				connectDots: func(g *game.Game) {},
				isSquare:    false,
			},
			{
				name: "1 line is not a square",
				connectDots: func(g *game.Game) {
					g.Connect(points.Coord{X: 0, Y: 0}, points.Right)
				},
				isSquare: false,
			},
			{
				name: "2 lines is not a square",
				connectDots: func(g *game.Game) {
					topLeft := points.Coord{X: 0, Y: 0}
					g.Connect(topLeft, points.Right)
					g.Connect(topLeft, points.Down)
				},
				isSquare: false,
			},
			{
				name: "3 lines is not a square",
				connectDots: func(g *game.Game) {
					topLeft := points.Coord{X: 0, Y: 0}
					bottomRight := topLeft.Translate(points.Down).Translate(points.Right)

					g.Connect(topLeft, points.Right)
					g.Connect(topLeft, points.Down)
					g.Connect(bottomRight, points.Left)
				},
				isSquare: false,
			},
			{
				name: "4 lines is a square",
				connectDots: func(g *game.Game) {
					topLeft := points.Coord{X: 0, Y: 0}
					bottomRight := topLeft.Translate(points.Down).Translate(points.Right)

					g.Connect(topLeft, points.Right)
					g.Connect(topLeft, points.Down)
					g.Connect(bottomRight, points.Left)
					g.Connect(bottomRight, points.Up)
				},
				isSquare: true,
			},
		} {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				g := game.New(1)
				tc.connectDots(g)
				assert.Equal(t, tc.isSquare, g.IsSquare(points.Coord{X: 0, Y: 0}))
			})
		}
	})

	specifications.Game{
		NewSubject: func(size int) specifications.GameDriver {
			return game.New(size)
		},
	}.Test(t)
}
