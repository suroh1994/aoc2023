package main

import (
	"aoc2023/1/common"
	"aoc2023/aoccommon"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")
	calibrationValues := calculateCalibrationValues(inputLines)
	calibrationValue := common.Sum(calibrationValues)
	err := common.WriteOutputLineByLine("output", calibrationValue)
	if err != nil {
		panic("oh crap")
	}
}

func calculateCalibrationValues(lines []string) []int {
	calibrationValues := make([]int, 0, len(lines))

	for _, line := range lines {
		indices := [2]int{len(line), -1}
		for i := range line {
			// exit out if we've searched the whole string from both ends
			if i > len(line)/2 {
				break
			}

			updateIndices(line, i, &indices)

			j := len(line) - 1 - i
			updateIndices(line, j, &indices)

			// exit out early if we found digits at both ends
			if indices[0] < len(line)/2 && indices[1] > len(line)/2 {
				break
			}
		}
		calibrationValues = append(calibrationValues, int((line[indices[0]]-48)*10+(line[indices[1]]-48)))
	}

	return calibrationValues
}

func updateIndices(line string, j int, indices *[2]int) {
	if common.IsCharDigit(line[j]) {
		if j < indices[0] {
			indices[0] = j
		}

		if j > indices[1] {
			indices[1] = j
		}
	}
}
