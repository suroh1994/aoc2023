package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
)

func main() {
	// read input
	lines := aoc.ReadInputLineByLine("input")

	start := -1
	smokeMaps := make([][]string, 0)
	for i, line := range lines {
		if line != "" && start == -1 {
			start = i
		}

		if line == "" {
			smokeMaps = append(smokeMaps, lines[start:i])
			start = -1
		}
	}
	smokeMaps = append(smokeMaps, lines[start:])

	sum := 0
	for _, smokeMap := range smokeMaps {
		multiplier := 100
		mirrorRow := findMirrorRow(smokeMap)
		if mirrorRow == -1 {
			multiplier = 1
			// flip map
			runeMap := aoc.Map(smokeMap, aoc.StringToRuneSlice)
			transposedRuneMap := aoc.Transpose(runeMap)
			smokeMap = aoc.Map(transposedRuneMap, aoc.RuneSliceToString)
			// find row
			mirrorRow = findMirrorRow(smokeMap)
		}
		score := (mirrorRow + 1) * multiplier
		fmt.Println(score)
		debugPrintSmokeMap(smokeMap, mirrorRow)
		sum += score
	}

	fmt.Println(sum)
}

func debugPrintSmokeMap(smokeMap []string, mirrorRowIdx int) {
	for i, row := range smokeMap {
		fmt.Println(row)
		if i == mirrorRowIdx {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
	}
	fmt.Println()
}

func findMirrorRow(smokeMap []string) int {
	for i := 0; i < len(smokeMap)-1; i++ {
		if smokeMap[i] == smokeMap[i+1] {
			allRowsMatch := true
			// if the next is identical, start going backwards (i=1; i < idx) and compare idx-i with idx+1+i until i == 0 or idx+1+i == len-1
			for j := 0; 0 <= i-j && i+1+j < len(smokeMap) && allRowsMatch; j++ {
				allRowsMatch = smokeMap[i-j] == smokeMap[i+1+j]
			}
			// if mirror plane found, return idx
			if allRowsMatch {
				return i
			}
		}
	}
	return -1
}
