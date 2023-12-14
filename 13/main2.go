package main

import (
	"aoc2023/13/common"
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
	for i, smokeMap := range smokeMaps {
		multiplier := 100
		mirrorRow := findMirrorRowV2(smokeMap)
		if mirrorRow == -1 {
			multiplier = 1
			// flip map
			runeMap := aoc.Map(smokeMap, aoc.StringToRuneSlice)
			transposedRuneMap := aoc.Transpose(runeMap)
			smokeMap = aoc.Map(transposedRuneMap, aoc.RuneSliceToString)
			// find row
			mirrorRow = findMirrorRowV2(smokeMap)
		}
		score := (mirrorRow + 1) * multiplier
		fmt.Printf("#%d: %d\n", i, score)
		common.DebugPrintSmokeMap(smokeMap, mirrorRow)
		sum += score
	}

	fmt.Println(sum)
}

func findMirrorRowV2(smokeMap []string) int {
	for i := 0; i < len(smokeMap)-1; i++ {
		differences := countDifferences(smokeMap[i], smokeMap[i+1])
		if differences < 2 {
			allRowsMatch := true
			smudgeFixed := differences == 1
			// if the next is identical, start going backwards (i=1; i < idx) and compare idx-i with idx+1+i until i == 0 or idx+1+i == len-1
			for j := 1; 0 <= i-j && i+1+j < len(smokeMap) && allRowsMatch; j++ {
				differences = countDifferences(smokeMap[i-j], smokeMap[i+1+j])

				allRowsMatch = differences == 0 || (differences == 1 && !smudgeFixed)
				if differences == 1 {
					smudgeFixed = true
				}
			}
			// if mirror plane found, return idx
			if allRowsMatch && smudgeFixed {
				return i
			}
		}
	}
	return -1
}

func countDifferences(lineA, lineB string) int {
	count := 0
	for i := range lineA {
		if lineA[i] != lineB[i] {
			count++
		}
	}
	return count
}
