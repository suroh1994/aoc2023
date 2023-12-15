package main

import (
	"aoc2023/15/common"
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

func main() {
	initSeq := aoc.ReadInputLineByLine("input")[0]
	segments := strings.Split(initSeq, ",")

	total := 0
	for _, segment := range segments {
		total += common.Hash(segment)
	}
	fmt.Println(total)
}
