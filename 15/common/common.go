package common

func Hash(in string) int {
	currentValue := 0
	for idx := range in {
		currentValue = currentValue + int(in[idx])
		currentValue = currentValue * 17
		currentValue = currentValue % 256
	}
	return currentValue
}
