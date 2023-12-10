package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")

	predictions := 0
	for _, line := range lines {
		numberStrings := strings.Split(line, " ")
		seqs := [][]int{aoc.Map(numberStrings, aoc.MustParseInt)}

		// calc differences, repeat until all zeros
		for !aoc.Reduce(seqs[len(seqs)-1], func(a int, b bool) bool {
			return b && a == 0
		}, true) {
			newSequence := make([]int, len(seqs[len(seqs)-1])-1)

			for j := range newSequence {
				newSequence[j] = seqs[len(seqs)-1][j+1] - seqs[len(seqs)-1][j]
			}

			seqs = append(seqs, newSequence)
		}

		// extrapolate
		seqs[len(seqs)-1] = append(seqs[len(seqs)-1], 0)
		for idx := len(seqs) - 1; idx > 0; idx-- {
			// since we added a value to the last sequence, the last two sequences have the same length
			lastIdx := len(seqs[idx]) - 1
			seqs[idx-1] = append(seqs[idx-1], seqs[idx-1][lastIdx]+seqs[idx][lastIdx])
		}

		// store the predicted value for
		predictions += seqs[0][len(seqs[0])-1]
	}

	fmt.Println(predictions)
}
