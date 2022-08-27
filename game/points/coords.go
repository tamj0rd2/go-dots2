package points

import "fmt"

type Coord struct {
	X, Y int
}

func (c Coord) ID() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

func (c Coord) Translate(translation Translation) Coord {
	switch translation {
	case Up:
		return Coord{X: c.X, Y: c.Y - 1}
	case Right:
		return Coord{X: c.X + 1, Y: c.Y}
	case Down:
		return Coord{X: c.X, Y: c.Y + 1}
	case Left:
		return Coord{X: c.X - 1, Y: c.Y}
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
