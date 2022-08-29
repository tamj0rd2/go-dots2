package points

import "fmt"

type Coords struct {
	X, Y int
}

type ID string

func (c Coords) ID() ID {
	return ID(fmt.Sprintf("%d,%d", c.X, c.Y))
}

func (c Coords) IsWithinBounds(width, height int) bool {
	return c.X < width && c.Y < height
}

func (c Coords) Translate(translation Translation) Coords {
	switch translation {
	case Up:
		return Coords{X: c.X, Y: c.Y - 1}
	case Right:
		return Coords{X: c.X + 1, Y: c.Y}
	case Down:
		return Coords{X: c.X, Y: c.Y + 1}
	case Left:
		return Coords{X: c.X - 1, Y: c.Y}
	default:
		panic(fmt.Errorf("unrecognised translation %v", translation))
	}
}

func (c Coords) String() string {
	return string(c.ID())
}

type Translation string

const (
	Up    Translation = "up"
	Right Translation = "right"
	Down  Translation = "down"
	Left  Translation = "left"

	TranslationCount = 4
)

func (t Translation) Opposite() Translation {
	switch t {
	case Up:
		return Down
	case Right:
		return Left
	case Down:
		return Up
	case Left:
		return Right
	default:
		panic(fmt.Errorf("unrecognised translation %v", t))
	}
}
