package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"math"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")

	var galaxies []aoc.Position2D
	columnsWithGalaxies := make([]bool, len(lines[0]))
	rowsWithGalaxies := make([]bool, len(lines))

	for idxRow, line := range lines {
		for idxCol, char := range line {
			if char != '#' {
				continue
			}

			columnsWithGalaxies[idxCol] = true
			rowsWithGalaxies[idxRow] = true
			galaxies = append(galaxies, aoc.Position2D{X: idxCol, Y: idxRow})
		}
	}

	fmt.Println(columnsWithGalaxies)
	fmt.Println(rowsWithGalaxies)
	fmt.Println(galaxies)

	expansionFactor := 999999

	colOffset := make([]int, len(columnsWithGalaxies))
	for idx := range columnsWithGalaxies {
		if columnsWithGalaxies[idx] {
			continue
		}

		for i := idx; i < len(colOffset); i++ {
			colOffset[i] += expansionFactor
		}
	}
	fmt.Println(colOffset)

	rowOffset := make([]int, len(rowsWithGalaxies))
	for idx := range rowsWithGalaxies {
		if rowsWithGalaxies[idx] {
			continue
		}

		for i := idx; i < len(rowOffset); i++ {
			rowOffset[i] += expansionFactor
		}
	}
	fmt.Println(rowOffset)

	sumOfDistances := 0
	for i, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies[i+1:] {
			offsetX1 := galaxy1.X + colOffset[galaxy1.X]
			offsetX2 := galaxy2.X + colOffset[galaxy2.X]
			sumOfDistances += int(math.Abs(float64(offsetX2 - offsetX1)))

			offsetY1 := galaxy1.Y + rowOffset[galaxy1.Y]
			offsetY2 := galaxy2.Y + rowOffset[galaxy2.Y]
			sumOfDistances += int(math.Abs(float64(offsetY2 - offsetY1)))
		}
	}

	fmt.Println(sumOfDistances)
}
