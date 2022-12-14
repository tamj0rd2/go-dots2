package game_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/tamj0rd2/go-dots2/game"
	"github.com/tamj0rd2/go-dots2/game/points"
	"github.com/tamj0rd2/go-dots2/specifications"
	"github.com/tamj0rd2/go-dots2/testutils"
)

func TestPlayingATwoPlayerGame(t *testing.T) {
	t.Run("connecting dots to make a square", func(t *testing.T) {
		for _, tc := range []struct {
			name        string
			connectDots func(g *game.Game)
			repr        string
			isSquare    bool
		}{
			{
				name:        "no lines",
				connectDots: func(g *game.Game) {},
				repr: `.   .

					   ˙   ˙`,
			},
			{
				name: "1 line",
				connectDots: func(g *game.Game) {
					g.DrawLine(points.Coords{X: 0, Y: 0}, points.Right)
				},
				repr: `.---.

					   ˙   ˙`,
			},
			{
				name: "2 lines",
				connectDots: func(g *game.Game) {
					topLeft := points.Coords{X: 0, Y: 0}
					g.DrawLine(topLeft, points.Right)
					g.DrawLine(topLeft, points.Down)
				},
				repr: `.---.
					   |
					   ˙   ˙`,
			},
			{
				name: "3 lines",
				connectDots: func(g *game.Game) {
					topLeft := points.Coords{X: 0, Y: 0}
					bottomRight := topLeft.Translate(points.Down).Translate(points.Right)

					g.DrawLine(topLeft, points.Right)
					g.DrawLine(topLeft, points.Down)
					g.DrawLine(bottomRight, points.Left)
				},
				repr: `.---.
					   |
					   ˙---˙`,
			},
			{
				name: "4 lines connected to make a square",
				connectDots: func(g *game.Game) {
					topLeft := points.Coords{X: 0, Y: 0}
					bottomRight := topLeft.Translate(points.Down).Translate(points.Right)

					g.DrawLine(topLeft, points.Right)
					g.DrawLine(topLeft, points.Down)
					g.DrawLine(bottomRight, points.Left)
					g.DrawLine(bottomRight, points.Up)
				},
				repr: `.---.
					   |   |
					   ˙---˙`,
				isSquare: true,
			},
		} {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				g := game.New(1)
				tc.connectDots(g)
				testutils.AssertGridEquals(t, tc.repr, g.Grid())
				assert.Equal(t, tc.isSquare, g.Score() > 0)
			})
		}
	})

	specifications.Game{
		NewSubject: func(size int) specifications.GameDriver {
			return game.New(size)
		},
	}.Test(t)
}
