package main

import (
	"aoc2023/10/common"
	aoc "aoc2023/aoccommon"
	"fmt"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")
	pipeMap := make([][]rune, len(lines))

	var startPosition aoc.Position2D
	for idxRow := range lines {
		pipeMap[idxRow] = make([]rune, len(lines[idxRow]))
		for idxCol, char := range lines[idxRow] {
			pipeMap[idxRow][idxCol] = char

			if char == 'S' {
				startPosition = aoc.Position2D{
					X: idxCol,
					Y: idxRow,
				}
			}
		}
	}

	stepCount := walkMap(pipeMap, startPosition)
	fmt.Println(stepCount / 2)
}

func walkMap(pipeMap [][]rune, startPos aoc.Position2D) int {
	if steps, looped := tryDirection(pipeMap, startPos.Add(common.Direction2OffsetMap[common.North]), common.North, 1); looped {
		return steps
	}
	if steps, looped := tryDirection(pipeMap, startPos.Add(common.Direction2OffsetMap[common.East]), common.East, 1); looped {
		return steps
	}
	if steps, looped := tryDirection(pipeMap, startPos.Add(common.Direction2OffsetMap[common.South]), common.South, 1); looped {
		return steps
	}
	panic("algorithm is wrong, loop not found!")
}

func tryDirection(pipeMap [][]rune, pos aoc.Position2D, forward common.Direction, stepCount int) (int, bool) {
	direction2DirectionMap, exist := common.PipeTypeAndDirection2Direction[pipeMap[pos.Y][pos.X]]
	if !exist {
		return -1, false
	}

	newForward, exists := direction2DirectionMap[forward]
	if !exists {
		return -1, false
	}

	newPos := pos.Add(common.Direction2OffsetMap[newForward])
	if newPos.X < 0 || newPos.Y < 0 || newPos.X > len(pipeMap[0])-1 || newPos.Y > len(pipeMap) {
		return -1, false
	}

	if pipeMap[newPos.Y][newPos.X] == 'S' {
		return stepCount + 1, true
	}

	return tryDirection(pipeMap, newPos, newForward, stepCount+1)
}
