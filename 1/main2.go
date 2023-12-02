package main

import (
	"aoc2023/1/common"
	"aoc2023/aoccommon"
	"strings"
)

var (
	letterWords = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")
	calibrationValues := calculateCalibrationValuesV2(inputLines)
	calibrationValue := common.Sum(calibrationValues)

	err := aoccommon.WriteOutputLineByLine("output", calibrationValue)
	if err != nil {
		panic("oh crap")
	}
}

func calculateCalibrationValuesV2(lines []string) []int {
	updatedInput := make([]int, len(lines))

	for i, line := range lines {
		leftMostDigit := -1
		rightMostDigit := -1
		for j := range line {
			if leftMostDigit == -1 {
				if common.IsCharDigit(line[j]) {
					leftMostDigit = int(line[j] - 48)
				} else {
					leftMostDigit = containsDigitWord(line[:j+1])
				}
			}

			if rightMostDigit == -1 {
				k := len(line) - 1 - j
				if common.IsCharDigit(line[k]) {
					rightMostDigit = int(line[k] - 48)
				} else {
					rightMostDigit = containsDigitWord(line[k:])
				}
			}

			if leftMostDigit != -1 && rightMostDigit != -1 {
				break
			}
		}

		updatedInput[i] = leftMostDigit*10 + rightMostDigit
	}
	return updatedInput
}

// Either returns the value of the digit word contained or -1
func containsDigitWord(input string) int {
	for idx, word := range letterWords {
		if strings.Contains(input, word) {
			return idx + 1
		}
	}
	return -1
}
