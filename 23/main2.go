package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
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
	// create a map with distances between crossings
	// finally test all options for traversing the map based on the known distances

	fmt.Println(longestPath)
}

func IsCrossing(pos aoc.Position2D, charMap [][]rune) bool {
	numOfPathsFromPos := 0

	for _, dir := range aoc.AllDirections {
		if newPos := pos.Add(dir); newPos.InBounds(len(charMap[0]), len(charMap)) && charMap[newPos.Y][newPos.X] != '#' {
			numOfPathsFromPos++
		}
	}

	return numOfPathsFromPos > 2
}
