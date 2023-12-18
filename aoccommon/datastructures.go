package aoccommon

type Position2D struct {
	X, Y int
}

var (
	DirNorth = Position2D{Y: -1}
	DirSouth = Position2D{Y: 1}
	DirEast  = Position2D{X: 1}
	DirWest  = Position2D{X: -1}
)

func (p1 Position2D) Add(p2 Position2D) Position2D {
	return Position2D{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}
