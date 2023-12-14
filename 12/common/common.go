package common

import (
	aoc "aoc2023/aoccommon"
	"slices"
)

func IsPossible(intermediateSolution string, currentIdx int, expectedLengths []int) bool {
	actualLengths := make([]int, 1)
	for i := range intermediateSolution[:currentIdx] {
		if intermediateSolution[i] == '#' {
			actualLengths[len(actualLengths)-1]++
		} else {
			if actualLengths[len(actualLengths)-1] != 0 {
				actualLengths = append(actualLengths, 0)
			}
		}
	}
	if actualLengths[len(actualLengths)-1] == 0 {
		actualLengths = actualLengths[:len(actualLengths)-1]
	}

	// check if what we found matches so far
	if len(actualLengths) > len(expectedLengths) {
		return false
	}

	for i := range actualLengths {
		if actualLengths[i] > expectedLengths[i] ||
			(i < len(actualLengths)-1 && actualLengths[i] != expectedLengths[i]) {
			return false
		}
	}

	// TODO lookahead to see the num of # you have to use and the max num of # you can place.
	maximumNumHashes := aoc.Reduce(expectedLengths, func(a int, total int) int {
		return a + total
	}, 0)
	hashesPlaced := aoc.Reduce([]rune(intermediateSolution), func(char rune, total int) int {
		if char == '#' {
			return total + 1
		}
		return total
	}, 0)

	return hashesPlaced <= maximumNumHashes

}

/** Idea:
Treat groups with ? and # as if they only contain ?
Then see how many segments you can fit in this group.
Then check which positions are occupied by # and drop all combinations that require a . in that position.
Move on to the next group and check starting at the last group used.

*/

func MatchesExpectations(states string, segmentLengths []int) bool {
	actualSegmentLength := make([]int, 1)
	for i := range states {
		if states[i] == '#' {
			actualSegmentLength[len(actualSegmentLength)-1]++
		} else {
			if actualSegmentLength[len(actualSegmentLength)-1] != 0 {
				actualSegmentLength = append(actualSegmentLength, 0)
			}
		}
	}
	if actualSegmentLength[len(actualSegmentLength)-1] == 0 {
		actualSegmentLength = actualSegmentLength[:len(actualSegmentLength)-1]
	}

	return slices.Equal(segmentLengths, actualSegmentLength)
}

func IsDone(solution string) bool {
	return !aoc.Reduce([]rune(solution), func(char rune, questionMarkEncountered bool) bool {
		return questionMarkEncountered || char == '?'
	}, false)
}

func CalculatePossibleSolutions(states string, segmentLengths []int) int {
	idx := 0
	for ; states[idx] != '?'; idx++ {
	}

	option1 := states[:idx] + "." + states[idx+1:]
	option2 := states[:idx] + "#" + states[idx+1:]

	return calculatePossibleSolutionsRecursive(option1, idx, segmentLengths) +
		calculatePossibleSolutionsRecursive(option2, idx, segmentLengths)

}

func calculatePossibleSolutionsRecursive(states string, currentIdx int, segmentLengths []int) int {
	if !IsPossible(states, currentIdx, segmentLengths) {
		return 0
	}

	if IsDone(states) {
		if MatchesExpectations(states, segmentLengths) {
			//fmt.Printf("Possible solution: %s\n", states)
			return 1
		} else {
			return 0
		}
	}

	idx := currentIdx
	for ; states[idx] != '?'; idx++ {
	}

	option1 := states[:idx] + "." + states[idx+1:]
	option2 := states[:idx] + "#" + states[idx+1:]

	return calculatePossibleSolutionsRecursive(option1, idx, segmentLengths) +
		calculatePossibleSolutionsRecursive(option2, idx, segmentLengths)

}
