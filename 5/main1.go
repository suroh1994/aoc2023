package main

import (
	"aoc2023/5/common"
	"aoc2023/aoccommon"
	"slices"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	seeds, mappings := common.ParseInput(inputLines)
	locations := common.FindLocationsForSeeds(seeds, mappings)
	slices.Sort(locations)

	err := aoccommon.WriteOutputLineByLine("output", locations[0])
	if err != nil {
		panic("oh crap")
	}
}
