package main

import (
	"aoc2023/2/common"
	"aoc2023/aoccommon"
)

func main() {
	set := common.Cubes{Red: 12, Green: 13, Blue: 14}

	inputLines := aoccommon.ReadInputLineByLine("input")

	summedIds := 0
	for _, line := range inputLines {
		id, draws := common.ParseLine(line)
		isGamePossible := true
		for _, draw := range draws {
			if !isGamePossibleWithSetOfCubes(set, draw) {
				isGamePossible = false
				break
			}
		}

		if isGamePossible {
			summedIds += id
		}
	}

	err := aoccommon.WriteOutputLineByLine("output", summedIds)
	if err != nil {
		panic("oh crap")
	}
}

func isGamePossibleWithSetOfCubes(set, draw common.Cubes) bool {
	return set.Red >= draw.Red && set.Green >= draw.Green && set.Blue >= draw.Blue
}
