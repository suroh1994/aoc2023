package main

import (
	"aoc2023/2/common"
	"aoc2023/aoccommon"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	summedPowers := 0
	for _, line := range inputLines {
		_, draws := common.ParseLine(line)
		minimalSet := calculateMinimalSetOfCubes(draws)
		summedPowers += minimalSet.Red * minimalSet.Green * minimalSet.Blue
	}

	err := aoccommon.WriteOutputLineByLine("output", summedPowers)
	if err != nil {
		panic("oh crap")
	}
}

func calculateMinimalSetOfCubes(draws []common.Cubes) common.Cubes {
	minimum := common.Cubes{}
	for _, draw := range draws {
		minimum.Red = max(draw.Red, minimum.Red)
		minimum.Green = max(draw.Green, minimum.Green)
		minimum.Blue = max(draw.Blue, minimum.Blue)
	}
	return minimum
}
