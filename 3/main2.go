package main

import (
	"aoc2023/3/common"
	"aoc2023/aoccommon"
	"unicode"
)

func main() {
	partNumbers := map[common.Position2D]*common.PartNumber{}
	potentialGears := make([]common.Position2D, 0)

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

			if char == '*' {
				potentialGears = append(potentialGears, common.Position2D{X: x, Y: y})
				currentPartNumber = nil
				continue
			}

			// if symbol = anything else (terminate current number and) add a marker to map
			currentPartNumber = nil
		}

	}

	// identify gears and calculate their gear ratio
	gearRatios := make([]int, 0)
	for _, potentialGear := range potentialGears {
		// check all neighbouring cells and count numbers
		numbers := make([]*common.PartNumber, 0)
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				partNumber, exists := partNumbers[common.Position2D{X: potentialGear.X + i, Y: potentialGear.Y + j}]
				if exists && !partNumber.AdjacentToSymbol {
					partNumber.AdjacentToSymbol = true
					numbers = append(numbers, partNumber)
				}
			}
		}

		if len(numbers) == 2 {
			gearRatios = append(gearRatios, numbers[0].Value*numbers[1].Value)
		}

		// reset partnumber adjecency tracking
		for _, number := range numbers {
			number.AdjacentToSymbol = false
		}
	}

	// sum all gear ratios
	sum := 0
	for _, gearRatio := range gearRatios {
		sum += gearRatio
	}

	err := aoccommon.WriteOutputLineByLine("output", sum)
	if err != nil {
		panic("oh crap")
	}
}
