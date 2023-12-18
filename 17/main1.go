package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
)

func main() {
	// parse input
	lines := aoc.ReadInputLineByLine("input")
	heatLossRunes := aoc.Map(lines, aoc.StringToRuneSlice)
	heatLossMap := aoc.Map(heatLossRunes, func(digits []rune) []int {
		numbers := make([]int, len(digits))
		for i, digit := range digits {
			numbers[i] = int(digit - 48)
		}
		return numbers
	})
	fmt.Println(heatLossMap)

	// for each direction, check if
	// new position is inbounds
	// direction hasn't been chosen three times already
	// then calculate heatloss and store
	fmt.Println(min(walkMap(heatLossMap, aoc.Position2D{}, aoc.DirEast), walkMap(heatLossMap, aoc.Position2D{}, aoc.DirSouth)))
}

type OpenMove struct {
	stepsLeft   int
	heatLost    int
	coordinates aoc.Position2D
	lastMoves   []aoc.Position2D
}

func walkMap(heatLossMap [][]int, startPos, startDir aoc.Position2D) int {
	knownCoordinates := make(map[aoc.Position2D]int)
	queue := aoc.NewSortedQueue[OpenMove](func(a, b OpenMove) int {
		return a.heatLost + a.stepsLeft - (b.heatLost + b.stepsLeft)
	})

	queue.Insert(OpenMove{
		stepsLeft: len(heatLossMap) + len(heatLossMap[0]),
		heatLost:  0,
		lastMoves: []aoc.Position2D{{}, {}, {}},
	})

	end := aoc.Position2D{
		X: len(heatLossMap[0]),
		Y: len(heatLossMap),
	}
	for _, exists := knownCoordinates[end]; !exists; {
		move, exists := queue.Pop()
		if !exists {
			panic("empty queue...Oh oh.")
		}

		knownCoordinates[move.coordinates] = move.heatLost
		for _, direction := range aoc.AllDirections {
			if !move.coordinates.Add(direction).InBounds(len(heatLossMap[0]), len(heatLossMap)) {
				continue
			}

			if allEqual(direction, move.lastMoves) {
				continue
			}

			newCoords := move.coordinates.Add(direction)

			queue.Insert(OpenMove{
				stepsLeft:   len(heatLossMap) - newCoords.Y + (len(heatLossMap[0]) - newCoords.X),
				heatLost:    move.heatLost + heatLossMap[newCoords.Y][newCoords.X],
				coordinates: newCoords,
				lastMoves:   append(move.lastMoves[1:], direction),
			})
		}
	}

	return knownCoordinates[end]
}

func allEqual(elem aoc.Position2D, compareTo []aoc.Position2D) bool {
	for _, comparator := range compareTo {
		if comparator != elem {
			return false
		}
	}
	return true
}
