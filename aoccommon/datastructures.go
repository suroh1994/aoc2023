package aoccommon

type Position2D struct {
	X, Y int
}

func (p1 Position2D) Add(p2 Position2D) Position2D {
	return Position2D{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}
