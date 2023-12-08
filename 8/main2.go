package main

import (
	"aoc2023/8/common"
	"aoc2023/aoccommon"
	"fmt"
	"math"
)

type mapLoop struct {
	ends   []End
	length int
}

func main() {
	lines := aoccommon.ReadInputLineByLine("input")
	directions := common.ParseDirections(lines[0])
	nodes := common.ParseMap(lines[2:])
	startingNodes := findAllStartingNodes(nodes)

	loops := make([]mapLoop, len(startingNodes))
	for idx := range startingNodes {
		_, length := findLoop(startingNodes[idx], nodes, directions)
		loops[idx] = mapLoop{
			length: length,
		}
	}

	primeMaps := make([]map[int]int, len(loops))
	for i, loop := range loops {
		primeMaps[i] = make(map[int]int)
		pfs := PrimeFactors(loop.length)
		for j := range pfs {
			primeMaps[i][pfs[j]]++
		}
	}

	minimizedPrimeMap := make(map[int]int)
	for _, pmap := range primeMaps {
		for k, v := range pmap {
			minimizedPrimeMap[k] = max(minimizedPrimeMap[k], v)

		}
	}

	var multiple float64 = 1
	for base, exp := range minimizedPrimeMap {
		multiple *= math.Pow(float64(base), float64(exp))
	}

	fmt.Println(int64(multiple))
}

// taken from https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
// Get all prime factors of a given number n
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func findAllStartingNodes(nodes map[string][2]string) []string {
	startingNodes := make([]string, 0)
	for node := range nodes {
		if isStartNode(node) {
			startingNodes = append(startingNodes, node)
		}
	}
	return startingNodes
}

type Visit struct {
	DirectionIndex int
	StepCount      int
}

type End struct {
	InLoop            bool
	StepsUntilReached int
}

func findLoop(startingNode string, nodes map[string][2]string, directions []int) (ends []End, loopLength int) {
	stepCount := 0
	dirIdx := 0

	visitedPlaces := map[string][]Visit{
		startingNode: {
			{
				DirectionIndex: stepCount,
				StepCount:      dirIdx,
			},
		},
	}
	for {
		// move to the next node
		startingNode = nodes[startingNode][directions[dirIdx]]
		dirIdx = (dirIdx + 1) % len(directions)
		stepCount++

		// see if we found a loop
		for _, visit := range visitedPlaces[startingNode] {
			if visit.DirectionIndex == dirIdx {
				// loop found. Extract all ending positions past the offset
				loopOffset := visit.StepCount
				loopLength = stepCount - visit.StepCount
				return getAllEnds(visitedPlaces, loopOffset), loopLength
			}
		}

		// mark current location as visited
		visitedPlaces[startingNode] = append(visitedPlaces[startingNode], Visit{
			DirectionIndex: dirIdx,
			StepCount:      stepCount,
		})
	}
}

func getAllEnds(visitedPlaces map[string][]Visit, loopOffset int) []End {
	endNodes := make([]End, 0)
	for node, visits := range visitedPlaces {
		if isEndNode(node) {
			for idx := range visits {
				endNodes = append(endNodes, End{
					InLoop:            visits[idx].StepCount > loopOffset,
					StepsUntilReached: visits[idx].StepCount,
				})
			}
		}
	}
	return endNodes
}

func isStartNode(node string) bool {
	return node[2] == 'A'
}

func isEndNode(node string) bool {
	return node[2] == 'Z'
}
