package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"math"
)

func main() {
	// parse input
	lines := aoc.ReadInputLineByLine("input")
	heatLossRunes := aoc.Map(lines, aoc.StringToRuneSlice)
	heatLossMap := aoc.Map(heatLossRunes, func(digits []rune) []int {
		numbers := make([]int, len(digits))
		for i, digit := range digits {
			numbers[i] = int(digit - 48)
		}
		return numbers
	})
	fmt.Println(heatLossMap)

	// for each direction, check if
	// new position is inbounds
	// direction hasn't been chosen three times already
	// then calculate heatloss and store
	fmt.Println(min(walkMap(heatLossMap, aoc.Position2D{}, aoc.DirEast), walkMap(heatLossMap, aoc.Position2D{}, aoc.DirSouth)))
}

type MoveHistory struct {
}

func walkMap(heatLossMap [][]int, startPos, startDir aoc.Position2D) int {

	return math.MaxInt
}
