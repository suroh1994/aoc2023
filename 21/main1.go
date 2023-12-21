package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"sync"
	"time"
)

var (
	maxDepthReached = 0
	startTime       time.Time
)

func main() {
	lines := aoc.ReadInputLineByLine("input")
	runesMap := aoc.Map(lines, aoc.StringToRuneSlice)

	startPos := aoc.Position2D{}
	for y := range runesMap {
		for x := range runesMap[y] {
			if runesMap[y][x] == 'S' {
				startPos.X = x
				startPos.Y = y
			}
		}
	}

	stepCount := 64
	reachablePositions := make(map[aoc.Position2D]*int)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	startTime = time.Now()
	for _, dir := range aoc.AllDirections {
		newPos := startPos.Add(dir)
		if runesMap[newPos.Y][newPos.X] == '#' {
			continue
		}

		wg.Add(1)
		go walkMapThread(&wg, &runesMap, newPos, stepCount-1, &mutex, &reachablePositions)
	}
	wg.Wait()

	for pos, stepsLeft := range reachablePositions {
		if (stepCount-(*stepsLeft))%2 == 0 {
			runesMap[pos.Y][pos.X] = 'O'
		}
	}

	for _, line := range aoc.Map(runesMap, aoc.RuneSliceToString) {
		fmt.Println(line)
	}
}

func walkMapThread(group *sync.WaitGroup, runesMap *[][]rune, startPos aoc.Position2D, stepsLeft int, mutex *sync.Mutex, reachablePositions *map[aoc.Position2D]*int) {
	walkMap(runesMap, startPos, stepsLeft, mutex, reachablePositions)
	group.Done()
}

func walkMap(runesMap *[][]rune, startPos aoc.Position2D, stepsLeft int, mutex *sync.Mutex, reachablePositions *map[aoc.Position2D]*int) {
	mutex.Lock()
	(*reachablePositions)[startPos] = &stepsLeft
	mutex.Unlock()

	// no more steps to take
	if stepsLeft == 0 {
		return
	}

	// walk in every direction that hasn't been touched yet
	for _, dir := range aoc.AllDirections {
		newPos := startPos.Add(dir)
		mutex.Lock()
		stepsLeftOnPrevVisit, alreadyVisited := (*reachablePositions)[newPos]
		mutex.Unlock()
		if (*runesMap)[newPos.Y][newPos.X] == '#' || (alreadyVisited && *stepsLeftOnPrevVisit >= stepsLeft) {
			continue
		}
		walkMap(runesMap, newPos, stepsLeft-1, mutex, reachablePositions)
	}

	if stepsLeft > maxDepthReached {
		maxDepthReached = max(maxDepthReached, stepsLeft)
		fmt.Printf("depth %d completed for the first time after %s\n", maxDepthReached, time.Now().Sub(startTime).String())
	}
}
