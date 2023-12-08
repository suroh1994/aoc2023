package main

import (
	"aoc2023/8/common"
	"aoc2023/aoccommon"
	"fmt"
)

func main() {
	lines := aoccommon.ReadInputLineByLine("input")
	directions := common.ParseDirections(lines[0])
	nodes := common.ParseMap(lines[2:])

	stepCount := traversMap(directions, nodes)
	fmt.Println(stepCount)
}

func traversMap(directions []int, nodes map[string][2]string) int {
	stepCount := 0
	dirIdx := 0
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		currentNode = nodes[currentNode][directions[dirIdx]]
		dirIdx = (dirIdx + 1) % len(directions)
		stepCount++
	}
	return stepCount
}
