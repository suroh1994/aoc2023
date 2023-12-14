package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")

	// rotate 90°  so that rocks fall to the right for easier parsing
	linesAsRunes := aoc.Map(lines, aoc.StringToRuneSlice)
	linesAsRunes = aoc.RotateRight(linesAsRunes)
	rotatedLines := aoc.Map(linesAsRunes, aoc.RuneSliceToString)

	totalWeight := 0
	fallingRocks := 0
	for _, line := range rotatedLines {
		for idx, char := range line {
			if char == 'O' {
				fallingRocks++
				continue
			}

			if char == '#' {
				for i := 0; i < fallingRocks; i++ {
					totalWeight += idx - i
				}
				fallingRocks = 0
			}
		}

		for i := 0; i < fallingRocks; i++ {
			totalWeight += len(line) - i
		}
		fallingRocks = 0
	}

	fmt.Println(totalWeight)
}
