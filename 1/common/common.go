package common

import (
	"os"
	"strconv"
)

func IsCharDigit(char uint8) bool {
	return char > 47 && char < 58
}

func WriteOutputLineByLine(outputFileName string, result int) error {
	err := os.WriteFile(outputFileName, []byte(strconv.Itoa(result)), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func Sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
