package main

import (
	"aoc2023/12/common"
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")

	sumOfPossibleSolutions := 0
	for i, line := range lines {
		inputSegments := strings.Split(line, " ")
		segmentLengths := aoc.Map(strings.Split(inputSegments[1], ","), aoc.MustParseInt)
		numPossibleSolutions := common.CalculatePossibleSolutions(inputSegments[0], segmentLengths)
		fmt.Printf("Line #%d has %d solutions\n", i, numPossibleSolutions)
		sumOfPossibleSolutions += numPossibleSolutions
	}
	fmt.Println(sumOfPossibleSolutions)
}
