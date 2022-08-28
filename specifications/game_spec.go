package specifications

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/tamj0rd2/go-dots2/game/points"
	"github.com/tamj0rd2/go-dots2/testutils"
)

type GameDriver interface {
	Connect(dot points.Coord, position points.Translation)
	IsSquare(coordinate points.Coord) bool
	Grid() string
}

type Game struct {
	NewSubject func(size int) GameDriver
}

func (spec Game) Test(t *testing.T) {
	t.Run("starting games of different sizes", func(t *testing.T) {
		for _, tc := range []struct {
			size int
			grid string
		}{
			{
				size: 1,
				grid: `.   .

					   ˙   ˙`,
			},
			{
				size: 2,
				grid: `.   .   .
        
				       :   :   :
        
                       ˙   ˙   ˙`,
			},
			{
				size: 3,
				grid: `.   .   .   .
        
        			   :   :   :   :
        
        			   :   :   :   :
        
        			   ˙   ˙   ˙   ˙`,
			},
			{
				size: 4,
				grid: ` .   .   .   .   .
        
                        :   :   :   :   :
        
                        :   :   :   :   :
        
                        :   :   :   :   :
        
                        ˙   ˙   ˙   ˙   ˙`,
			},
		} {
			tc := tc
			t.Run(fmt.Sprintf("%dx%d", tc.size, tc.size), func(t *testing.T) {
				game := spec.NewSubject(tc.size)
				testutils.AssertGridEquals(t, tc.grid, game.Grid())
			})
		}
	})

	t.Run("connecting dots in a grid", func(t *testing.T) {
		game := spec.NewSubject(2)
		testutils.AssertGridEquals(t, `
			.   .   .
			
			:   :   :

			˙   ˙   ˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 0, Y: 0}, points.Right)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:   :   :

			˙   ˙   ˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 0, Y: 1}, points.Down)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:   :   :
			|
			˙   ˙   ˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 1, Y: 1}, points.Right)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:   :---:
			|
			˙   ˙   ˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 0, Y: 1}, points.Right)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:---:---:
			|
			˙   ˙   ˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 2, Y: 1}, points.Down)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:---:---:
			|       |
			˙   ˙   ˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 2, Y: 2}, points.Left)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:---:---:
			|       |
			˙   ˙---˙ 
		`, game.Grid())

		game.Connect(points.Coord{X: 1, Y: 2}, points.Up)
		testutils.AssertGridEquals(t, `
			.---.   .
			
			:---:---:
			|   |   |
			˙   ˙---˙ 
		`, game.Grid())

		assert.False(t, game.IsSquare(points.Coord{X: 0, Y: 0}))
		assert.False(t, game.IsSquare(points.Coord{X: 1, Y: 0}))
		assert.False(t, game.IsSquare(points.Coord{X: 0, Y: 1}))
		assert.True(t, game.IsSquare(points.Coord{X: 1, Y: 1}))
	})
}
