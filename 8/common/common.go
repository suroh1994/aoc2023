package common

func ParseDirections(line string) []int {
	directions := make([]int, len(line))
	for idx := range line {
		if line[idx] == 'L' {
			directions[idx] = 0
		} else {
			directions[idx] = 1
		}
	}
	return directions
}

func ParseMap(lines []string) map[string][2]string {
	output := make(map[string][2]string)
	for idx := range lines {
		output[lines[idx][:3]] = [2]string{lines[idx][7:10], lines[idx][12:15]}
	}
	return output
}
