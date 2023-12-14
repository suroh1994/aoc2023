package common

import "fmt"

func DebugPrintSmokeMap(smokeMap []string, mirrorRowIdx int) {
	for i, row := range smokeMap {
		fmt.Println(row)
		if i == mirrorRowIdx {
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
	}
	fmt.Println()
}
