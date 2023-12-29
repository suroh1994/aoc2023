package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"slices"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")

	charMap := aoc.Map(lines, aoc.StringToRuneSlice)

	// determine start & end positions
	startPos := aoc.Position2D{}
	for x, char := range charMap[0] {
		if char == '.' {
			startPos.X = x
		}
	}
	endPos := aoc.Position2D{Y: len(charMap) - 1}
	for x, char := range charMap[len(charMap)-1] {
		if char == '.' {
			endPos.X = x
		}
	}

	// walk map recursively
	longestPath := 0
	paths := [][]aoc.Position2D{{startPos}}
	pathsCompleted := 0

	for len(paths) > 0 {
		// grab path from paths slice
		currentPath := paths[0]
		paths = paths[1:]

		// check if latest position == endpos and update longestPath
		lastPosition := currentPath[len(currentPath)-1]
		if lastPosition == endPos {
			longestPath = max(longestPath, len(currentPath)-1)
			pathsCompleted++
			fmt.Printf("Completed %d paths. Newest path is %d steps long.\n", pathsCompleted, len(currentPath)-1)
			continue
		}
		// otherwise...

		// check all neighbouring positions
		for _, dir := range aoc.AllDirections {
			nextPosition := lastPosition.Add(dir)

			// if already in path, skip
			if slices.Contains(currentPath, nextPosition) {
				continue
			}

			// not a valid position
			if !nextPosition.InBounds(len(charMap[0]), len(charMap)) {
				continue
			}

			// check symbol allows move
			char := charMap[nextPosition.Y][nextPosition.X]
			if char == '#' ||
				char == 'v' && dir == aoc.DirNorth ||
				char == '^' && dir == aoc.DirSouth ||
				char == '>' && dir == aoc.DirWest ||
				char == '<' && dir == aoc.DirEast {
				continue
			}

			// create a copy of the current path and add it to queue
			newPath := make([]aoc.Position2D, len(currentPath)+1)
			copy(newPath, currentPath)
			newPath[len(newPath)-1] = nextPosition

			paths = append(paths, newPath)
		}
	}
	fmt.Println(longestPath)
}
