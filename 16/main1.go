package main

import (
	"aoc2023/16/common"
	aoc "aoc2023/aoccommon"
	"fmt"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")
	mirrorMap := aoc.Map(lines, aoc.StringToRuneSlice)
	startPos := aoc.Position2D{}
	startDir := aoc.Position2D{X: 1}

	visitedMap := map[aoc.Position2D][]aoc.Position2D{
		startPos: {startDir},
	}
	common.FollowBeam(mirrorMap, startPos, startDir, &visitedMap)

	keyCount := 0
	for range visitedMap {
		keyCount++
	}
	fmt.Println(keyCount)
}
