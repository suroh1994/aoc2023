package main

import (
	"aoc2023/5/common"
	"aoc2023/aoccommon"
	"slices"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	seeds, mappings := common.ParseInput(inputLines)
	tuples := convertSeedsToTuples(seeds)

	for _, mapping := range mappings {
		tuples = mapping.MapValueRangesToDestValueRanges(tuples)
	}
	slices.SortFunc(tuples, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	err := aoccommon.WriteOutputLineByLine("output", tuples[0][0])
	if err != nil {
		panic("oh crap")
	}
}

func convertSeedsToTuples(seedsSlice []int) [][2]int {
	if len(seedsSlice)%2 != 0 {
		panic("uneven number of values for seeds found!")
	}

	output := make([][2]int, len(seedsSlice)/2)
	for idx := 0; idx < len(seedsSlice); idx += 2 {
		output[idx/2] = [2]int{seedsSlice[idx], seedsSlice[idx+1]}
	}

	return output
}
