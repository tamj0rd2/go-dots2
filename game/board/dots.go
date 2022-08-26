package board

import "fmt"

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Translate(translation Translation) Coordinate {
	switch translation {
	case Up:
		return Coordinate{X: c.X, Y: c.Y - 1}
	case Right:
		return Coordinate{X: c.X + 1, Y: c.Y}
	case Down:
		return Coordinate{X: c.X, Y: c.Y + 1}
	case Left:
		return Coordinate{X: c.X - 1, Y: c.Y}
	default:
		panic(fmt.Errorf("unrecognised translation %v", translation))
	}
}

type Translation string

const (
	Up    Translation = "up"
	Right Translation = "right"
	Down  Translation = "down"
	Left  Translation = "left"
)

func (c Coordinate) ID() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
