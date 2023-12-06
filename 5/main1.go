package main

import (
	"aoc2023/5/common"
	"aoc2023/aoccommon"
	"fmt"
	"slices"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	seeds, mappings := common.ParseInput(inputLines)

	fmt.Println("~~~~~~~~~~~~~~~~~")
	locations := seeds
	for _, mapping := range mappings {
		fmt.Printf("%s: %d elements = %v\n", mapping.Source, len(locations), locations)
		locations = mapping.MapValuesToDestValues(locations)
	}
	fmt.Printf("%s: %d elements = %v\n", mappings[len(mappings)-1].Destination, len(locations), locations)
	fmt.Println("~~~~~~~~~~~~~~~~~")
	slices.Sort(locations)

	err := aoccommon.WriteOutputLineByLine("output", locations[0])
	if err != nil {
		panic("oh crap")
	}
}
