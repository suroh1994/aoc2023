package main

import (
	"aoc2023/7/common"
	"aoc2023/aoccommon"
	"slices"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	hands := common.ParseInput(inputLines, false)

	slices.SortFunc(hands, func(a, b common.CamelCardHand) int {
		if a.HandType != b.HandType {
			return a.HandType - b.HandType
		}

		for idx := 0; idx < 5; idx++ {
			if a.Cards[idx] == b.Cards[idx] {
				continue
			}
			// desc order!
			return a.Cards[idx] - b.Cards[idx]
		}
		return 0
	})

	totalWinnings := 0
	for idx := range hands {
		totalWinnings = totalWinnings + hands[idx].Bid*(idx+1)
	}

	err := aoccommon.WriteOutputLineByLine("output", totalWinnings)
	if err != nil {
		panic("oh crap")
	}
}
