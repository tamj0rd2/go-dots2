package board

import "fmt"

type Coordinate struct {
	X, Y int
}

func (c Coordinate) ID() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
