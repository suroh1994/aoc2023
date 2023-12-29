package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"slices"
)

var (
	startPos aoc.Position2D
	endPos   aoc.Position2D
)

func main() {
	lines := aoc.ReadInputLineByLine("input")

	charMap := aoc.Map(lines, aoc.StringToRuneSlice)

	// determine start & end positions
	for x, char := range charMap[0] {
		if char == '.' {
			startPos.X = x
		}
	}
	for x, char := range charMap[len(charMap)-1] {
		if char == '.' {
			endPos.X = x
			endPos.Y = len(charMap) - 1
		}
	}

	// walk map recursively
	connectionsAtlas := map[aoc.Position2D]map[aoc.Position2D]int{}
	crossings := []aoc.Position2D{startPos, endPos}

	// identify all crossings
	for y := range charMap {
		for x := range charMap[y] {
			pos := aoc.Position2D{X: x, Y: y}
			if IsCrossing(pos, charMap) {
				crossings = append(crossings, pos)
			}
		}
	}

	// for each crossing, start walking in all directions until you reach a crossing/deadend/loop back onto yourself
	for _, crossing := range crossings {
		for _, dir := range aoc.AllDirections {
			// check validity
			neighbour := crossing.Add(dir)
			if !neighbour.InBounds(len(charMap[0]), len(charMap)) ||
				charMap[neighbour.Y][neighbour.X] == '#' {
				continue
			}

			path := [][]aoc.Position2D{{crossing, neighbour}}
			nextCrossing, steps := FollowTrailUntilNextCrossing(path, charMap)

			// deadend/loop/whatever, skip
			if steps == -1 {
				continue
			}

			// update connection atlas
			innerMap, exists := connectionsAtlas[crossing]
			if !exists {
				innerMap = make(map[aoc.Position2D]int)
			}
			innerMap[nextCrossing] = steps
			connectionsAtlas[crossing] = innerMap

			innerMap, exists = connectionsAtlas[nextCrossing]
			if !exists {
				innerMap = make(map[aoc.Position2D]int)
			}
			innerMap[crossing] = steps

		}
	}

	//for crossing1 := range connectionsAtlas {
	//	fmt.Printf("(%d, %d)\t↓\n", crossing1.Y, crossing1.X)
	//	for crossing2, steps := range connectionsAtlas[crossing1] {
	//		fmt.Printf("\t(%d, %d) in %d\n", crossing2.Y, crossing2.X, steps)
	//	}
	//}

	// finally test all options for traversing the map based on the known distances
	longestPath := FindPath([]aoc.Position2D{startPos}, 0, connectionsAtlas)

	fmt.Println(longestPath)
}

func IsCrossing(pos aoc.Position2D, charMap [][]rune) bool {
	if !pos.InBounds(len(charMap[0]), len(charMap)) || charMap[pos.Y][pos.X] == '#' {
		return false
	}

	numOfPathsFromPos := 0

	for _, dir := range aoc.AllDirections {
		if newPos := pos.Add(dir); newPos.InBounds(len(charMap[0]), len(charMap)) && charMap[newPos.Y][newPos.X] != '#' {
			numOfPathsFromPos++
		}
	}

	return numOfPathsFromPos > 2
}

func FollowTrailUntilNextCrossing(paths [][]aoc.Position2D, charMap [][]rune) (aoc.Position2D, int) {
	for len(paths) > 0 {
		// grab path from paths slice
		currentPath := paths[0]
		paths = paths[1:]

		// check if latest position == endpos and update longestPath
		lastPosition := currentPath[len(currentPath)-1]
		if IsCrossing(lastPosition, charMap) || lastPosition == startPos || lastPosition == endPos {
			return lastPosition, len(currentPath) - 1
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
			if char == '#' {
				continue
			}

			// create a copy of the current path and add it to queue
			newPath := make([]aoc.Position2D, len(currentPath)+1)
			copy(newPath, currentPath)
			newPath[len(newPath)-1] = nextPosition

			paths = append(paths, newPath)
		}
	}

	// deadend/loop
	return aoc.Position2D{}, -1
}

func FindPath(crossingsVisited []aoc.Position2D, distanceTraveled int, connectionAtlas map[aoc.Position2D]map[aoc.Position2D]int) int {
	connections := connectionAtlas[crossingsVisited[len(crossingsVisited)-1]]

	longestDistance := 0
	for crossing, distance := range connections {
		// crossing already visited
		if slices.Contains(crossingsVisited, crossing) {
			continue
		}

		// end reached
		if crossing == endPos {
			//fmt.Printf("%d: ", distanceTraveled+distance)
			//for _, pos := range crossingsVisited {
			//	fmt.Printf("(%d, %d) -> ", pos.X, pos.Y)
			//}
			//fmt.Printf("(%d, %d)\n", crossing.X, crossing.Y)

			longestDistance = max(longestDistance, distanceTraveled+distance)
		}

		path := make([]aoc.Position2D, len(crossingsVisited)+1)
		copy(path, crossingsVisited)
		path[len(path)-1] = crossing

		longestDistance = max(longestDistance, FindPath(path, distanceTraveled+distance, connectionAtlas))
	}

	return longestDistance
}
