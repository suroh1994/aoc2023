package main

import (
	"aoc2023/16/common"
	aoc "aoc2023/aoccommon"
	"fmt"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")
	mirrorMap := aoc.Map(lines, aoc.StringToRuneSlice)

	maxKeyCount := 0
	// top & bottom
	for x := range mirrorMap[0] {
		// starting at the top
		startPos := aoc.Position2D{X: x}
		startDir := common.DirSouth

		visitedMap := map[aoc.Position2D][]aoc.Position2D{
			startPos: {startDir},
		}

		common.FollowBeam(mirrorMap, startPos, startDir, &visitedMap)

		keyCount := 0
		for range visitedMap {
			keyCount++
		}
		maxKeyCount = max(keyCount, maxKeyCount)

		// starting at the bottom
		startPos = aoc.Position2D{X: x, Y: len(mirrorMap) - 1}
		startDir = common.DirNorth

		visitedMap = map[aoc.Position2D][]aoc.Position2D{
			startPos: {startDir},
		}

		common.FollowBeam(mirrorMap, startPos, startDir, &visitedMap)

		keyCount = 0
		for range visitedMap {
			keyCount++
		}
		maxKeyCount = max(keyCount, maxKeyCount)
	}

	// left & right
	for y := range mirrorMap {
		// starting on the left
		startPos := aoc.Position2D{Y: y}
		startDir := common.DirEast

		visitedMap := map[aoc.Position2D][]aoc.Position2D{
			startPos: {startDir},
		}

		common.FollowBeam(mirrorMap, startPos, startDir, &visitedMap)

		keyCount := 0
		for range visitedMap {
			keyCount++
		}
		maxKeyCount = max(keyCount, maxKeyCount)

		// starting on the right
		startPos = aoc.Position2D{X: len(mirrorMap[0]) - 1, Y: y}
		startDir = common.DirWest

		visitedMap = map[aoc.Position2D][]aoc.Position2D{
			startPos: {startDir},
		}

		common.FollowBeam(mirrorMap, startPos, startDir, &visitedMap)

		keyCount = 0
		for range visitedMap {
			keyCount++
		}
		maxKeyCount = max(keyCount, maxKeyCount)
	}
	fmt.Println(maxKeyCount)
}
