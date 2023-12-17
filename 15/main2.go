package main

import (
	"aoc2023/15/common"
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

type LabeledLense struct {
	Label    string
	Strength int
}

func main() {
	initSeq := aoc.ReadInputLineByLine("input")[0]
	segments := strings.Split(initSeq, ",")

	hashMap := make(map[int][]LabeledLense)

	for _, segment := range segments {
		opIdx := 0
		for ; segment[opIdx] != '=' && segment[opIdx] != '-'; opIdx++ {
		}
		mapKey := common.Hash(segment[:opIdx])

		// parse lens info
		lensLabel := segment[:opIdx]
		operation := segment[opIdx]
		lensStrength := -1
		if operation == '=' {
			lensStrength = aoc.MustParseInt(segment[opIdx+1:])
		}

		// get box
		box, exists := hashMap[mapKey]
		if !exists && operation == '=' {
			box = make([]LabeledLense, 1)
			box[0] = LabeledLense{Label: lensLabel, Strength: lensStrength}
		} else {
			// search through box
			lensIdx := 0
			for ; lensIdx < len(box); lensIdx++ {
				// lens found
				if box[lensIdx].Label == lensLabel {
					// remove entry
					if operation == '-' {
						if lensIdx+1 < len(box) {
							box = append(box[:lensIdx], box[lensIdx+1:]...)
						} else {
							box = append(box[:lensIdx])
						}
						break
					}
					// update entry
					box[lensIdx].Strength = lensStrength
					break
				}
			}

			// finally add lens if new
			if lensIdx == len(box) && operation == '=' {
				box = append(box, LabeledLense{Label: lensLabel, Strength: lensStrength})
			}

		}

		// update hashmap
		hashMap[mapKey] = box
	}

	total := 0
	for mapKey, box := range hashMap {
		for lensSlot, lens := range box {
			total += (mapKey + 1) * (lensSlot + 1) * lens.Strength
		}
	}
	fmt.Println(total)
}
