package common

import aoc "aoc2023/aoccommon"

type Direction int

const (
	North Direction = iota
	East  Direction = iota
	South Direction = iota
	West  Direction = iota
)

var (
	PipeTypeAndDirection2Direction = map[rune]map[Direction]Direction{
		'|': {
			North: North,
			South: South,
		},
		'-': {
			East: East,
			West: West,
		},
		'F': {
			West:  South,
			North: East,
		},
		'L': {
			West:  North,
			South: East,
		},
		'7': {
			East:  South,
			North: West,
		},
		'J': {
			South: West,
			East:  North,
		},
	}
	Direction2OffsetMap = map[Direction]aoc.Position2D{
		North: {X: 0, Y: -1},
		South: {X: 0, Y: 1},
		East:  {X: 1, Y: 0},
		West:  {X: -1, Y: 0},
	}
)
