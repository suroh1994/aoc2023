package aoccommon

type Position2D struct {
	X, Y int
}

var (
	DirNorth      = Position2D{Y: -1}
	DirSouth      = Position2D{Y: 1}
	DirEast       = Position2D{X: 1}
	DirWest       = Position2D{X: -1}
	AllDirections = []Position2D{DirNorth, DirEast, DirSouth, DirWest}
)

func (p Position2D) Add(q Position2D) Position2D {
	return Position2D{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

func (p Position2D) InBounds(width, height int) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

type SortedQueue[T any] struct {
	sortFunction func(a, b T) int
	content      []T
}

func NewSortedQueue[T any](sortFunc func(a, b T) int) SortedQueue[T] {
	return SortedQueue[T]{
		sortFunction: sortFunc,
		content:      make([]T, 0),
	}
}

func (s *SortedQueue[T]) Insert(element T) {
	for i := 0; i < len(s.content); i++ {
		if s.sortFunction(s.content[i], element) > 0 {
			s.content = append(s.content[:i+1], s.content[i:]...)
			s.content[i] = element
			return
		}
	}
	s.content = append(s.content, element)
}

func (s *SortedQueue[T]) Pop() (T, bool) {
	if len(s.content) == 0 {
		// create a default value
		var empty T
		return empty, false
	}

	head := s.content[0]
	s.content = s.content[1:]
	return head, true
}

func (s *SortedQueue[T]) Clean(cleanFunc func(T) bool) {
	for idx := 0; idx < len(s.content); {
		if cleanFunc(s.content[idx]) {
			s.content = append(s.content[:idx], s.content[idx+1:]...)
		} else {
			idx++
		}
	}
}
