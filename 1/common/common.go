package common

func IsCharDigit(char uint8) bool {
	return char > 47 && char < 58
}

func Sum(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
