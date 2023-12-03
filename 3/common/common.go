package common

type Position2D struct {
	X int
	Y int
}

type PartNumber struct {
	Value            int
	AdjacentToSymbol bool
	IncludedInSum    bool
}

func CharToDigit(char int32) int {
	return int(char) - 48
}
