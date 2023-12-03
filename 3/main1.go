package main

import (
	"aoc2023/3/common"
	"aoc2023/aoccommon"
	"unicode"
)

func main() {
	partNumbers := map[common.Position2D]*common.PartNumber{}
	symbols := make([]common.Position2D, 0)
	// parse the file line by line
	lines := aoccommon.ReadInputLineByLine("input")

	for x, line := range lines {
		// go character by character:
		var currentPartNumber *common.PartNumber
		for y, char := range line {
			// if symbol = . (terminate current number and) continue
			// if symbol = \n (terminate current number and) continue
			if char == '.' || char == '\n' {
				currentPartNumber = nil
				continue
			}

			// if symbol = \d start number or add to current number (current * 10 + new)
			if unicode.IsDigit(char) {
				if currentPartNumber == nil {
					currentPartNumber = &common.PartNumber{}
				}

				currentPartNumber.Value = currentPartNumber.Value*10 + common.CharToDigit(char)
				partNumbers[common.Position2D{X: x, Y: y}] = currentPartNumber
				continue
			}

			// if symbol = anything else (terminate current number and) add a marker to map
			currentPartNumber = nil
			symbols = append(symbols, common.Position2D{X: x, Y: y})
		}

	}

	// iterate over symbols
	for _, symbolPosition := range symbols {
		// check all neighbouring cells and mark numbers as adjacent
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				partNumber, exists := partNumbers[common.Position2D{X: symbolPosition.X + i, Y: symbolPosition.Y + j}]
				if exists {
					partNumber.AdjacentToSymbol = true
				}
			}
		}
	}

	// sum all part numbers marked adjacent
	sum := 0
	for _, partNumber := range partNumbers {
		if partNumber.AdjacentToSymbol && !partNumber.IncludedInSum {
			sum += partNumber.Value
			partNumber.IncludedInSum = true
		}
	}

	err := aoccommon.WriteOutputLineByLine("output", sum)
	if err != nil {
		panic("oh crap")
	}
}
