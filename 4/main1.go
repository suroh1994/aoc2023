package main

import (
	"aoc2023/4/common"
	"aoc2023/aoccommon"
)

func main() {

	inputLines := aoccommon.ReadInputLineByLine("input")

	summedPoints := 0
	for _, line := range inputLines {
		card := common.ParseLine(line)

		summedPoints += common.CalculatePoints(common.CountWinningNumbers(card))
	}

	err := aoccommon.WriteOutputLineByLine("output", summedPoints)
	if err != nil {
		panic("oh crap")
	}
}
