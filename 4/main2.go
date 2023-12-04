package main

import (
	"aoc2023/4/common"
	"aoc2023/aoccommon"
	"fmt"
)

func main() {

	inputLines := aoccommon.ReadInputLineByLine("input")

	cards := make([]common.ScratchCard, len(inputLines))
	for idx, line := range inputLines {
		cards[idx] = common.ParseLine(line)
	}

	knownResults := make(map[int]int)
	sum := 0
	for idx := range cards {
		wonCards := checkWonCards(idx, &cards, &knownResults)
		sum += wonCards
	}

	err := aoccommon.WriteOutputLineByLine("output", sum)
	if err != nil {
		panic("oh crap")
	}
}

func checkWonCards(id int, cards *[]common.ScratchCard, knownResults *map[int]int) int {
	fmt.Printf("checking card #%d\n", id+1)
	if count, exists := (*knownResults)[id]; exists {
		return count
	}

	fmt.Printf("calculating card #%d\n", id+1)
	if len(*cards) <= id {
		(*knownResults)[id] = 1
		return 0
	}

	points := common.CountWinningNumbers((*cards)[id])

	cardsWon := 1
	for i := 1; i <= points; i++ {
		cardsWon += checkWonCards(id+i, cards, knownResults)

	}

	(*knownResults)[id] = cardsWon
	return cardsWon
}
