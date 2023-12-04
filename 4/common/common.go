package common

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type ScratchCard struct {
	ID             int
	WinningNumbers []int
	DrawnNumbers   []int
}

func ParseLine(line string) ScratchCard {
	// split at :
	cardAndNumbers := strings.Split(line, ": ")
	// get id (skip "Card")
	cardId, err := strconv.ParseInt(strings.TrimSpace(cardAndNumbers[0][4:]), 10, 64)
	if err != nil {
		panic("failed to parse game ID from string: '" + cardAndNumbers[0] + "'")
	}

	// split at |
	winningAndOwnNumbers := strings.Split(cardAndNumbers[1], " | ")
	return ScratchCard{
		ID:             int(cardId),
		WinningNumbers: parseNumbers(winningAndOwnNumbers[0]),
		DrawnNumbers:   parseNumbers(winningAndOwnNumbers[1]),
	}
}

func parseNumbers(line string) []int {
	numberStrings := strings.Split(strings.TrimSpace(line), " ")
	numbers := make([]int, 0)
	for _, numberString := range numberStrings {
		if len(strings.TrimSpace(numberString)) == 0 {
			continue
		}

		number, err := strconv.ParseInt(strings.TrimSpace(numberString), 10, 64)
		if err != nil {
			panic("ah crap")
		}
		numbers = append(numbers, int(number))
	}

	return numbers
}

func CountWinningNumbers(card ScratchCard) int {
	slices.Sort(card.WinningNumbers)
	slices.Sort(card.DrawnNumbers)

	idxW := 0
	idxD := 0
	wins := 0
	for idxW < len(card.WinningNumbers) && idxD < len(card.DrawnNumbers) {
		if card.WinningNumbers[idxW] < card.DrawnNumbers[idxD] {
			idxW++
		} else if card.WinningNumbers[idxW] > card.DrawnNumbers[idxD] {
			idxD++
		} else if card.WinningNumbers[idxW] == card.DrawnNumbers[idxD] {
			wins++
			idxW++
			idxD++
		}
	}

	return wins
}

func CalculatePoints(wins int) int {
	if wins == 0 {
		return 0
	}

	return int(math.Pow(2, float64(wins-1)))
}
