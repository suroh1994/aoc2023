package common

import (
	aoc "aoc2023/aoccommon"
	"slices"
)

var (
	DirNorth = aoc.Position2D{Y: -1}
	DirSouth = aoc.Position2D{Y: 1}
	DirEast  = aoc.Position2D{X: 1}
	DirWest  = aoc.Position2D{X: -1}
)

func NextDirections(mirrorMap [][]rune, pos, dir aoc.Position2D) []aoc.Position2D {
	switch mirrorMap[pos.Y][pos.X] {
	case '-':
		{
			if dir.Y == 0 {
				return []aoc.Position2D{dir}
			}
			return []aoc.Position2D{DirEast, DirWest}
		}
	case '|':
		{
			if dir.X == 0 {
				return []aoc.Position2D{dir}
			}
			return []aoc.Position2D{DirNorth, DirSouth}
		}
	case '\\':
		{
			return []aoc.Position2D{
				{
					X: dir.Y,
					Y: dir.X,
				},
			}
		}
	case '/':
		{
			return []aoc.Position2D{
				{
					X: dir.Y * -1,
					Y: dir.X * -1,
				},
			}
		}
	default:
		{
			return []aoc.Position2D{dir}
		}
	}
}

func FollowBeam(mirrorMap [][]rune, pos, dir aoc.Position2D, visitedMap *map[aoc.Position2D][]aoc.Position2D) {
	directions := NextDirections(mirrorMap, pos, dir)
	for _, direction := range directions {
		newPos := pos.Add(direction)
		if newPos.X < 0 || newPos.X >= len(mirrorMap[0]) || newPos.Y < 0 || newPos.Y >= len(mirrorMap) {
			// we've gone out of bounds, no more work to do
			continue
		}

		if directionsGoing, exists := (*visitedMap)[newPos]; exists && slices.Contains(directionsGoing, direction) {
			// nothing new, we are likely going in a circle
			continue
		}

		(*visitedMap)[newPos] = append((*visitedMap)[newPos], direction)
		FollowBeam(mirrorMap, newPos, direction, visitedMap)
	}
}
