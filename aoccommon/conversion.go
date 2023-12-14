package aoccommon

import "strconv"

func MustParseInt(in string) int {
	value, err := strconv.Atoi(in)
	if err != nil {
		panic("didn't parse to an int: " + in)
	}
	return value
}

func StringToRuneSlice(in string) []rune {
	return []rune(in)
}

func RuneSliceToString(in []rune) string {
	return string(in)
}
