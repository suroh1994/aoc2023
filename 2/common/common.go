package common

import (
	"strconv"
	"strings"
)

type Cubes struct {
	Red   int
	Blue  int
	Green int
}

func ParseLine(line string) (int, []Cubes) {
	// split at :
	gameAndDraws := strings.Split(line, ":")
	// get id (skip "Game ")
	gameId, err := strconv.ParseInt(gameAndDraws[0][5:], 10, 64)
	if err != nil {
		panic("failed to parse game ID from string: '" + gameAndDraws[0] + "'")
	}

	games := gameAndDraws[1]
	// split at ;
	gamesSlice := strings.Split(games, ";")
	draws := make([]Cubes, len(gamesSlice))
	for idx, game := range gamesSlice {
		// split at ,
		drawStrings := strings.Split(game, ",")

		drawnCubes := Cubes{}
		// get counts by splitting at " "
		for _, drawString := range drawStrings {
			draw := strings.Split(strings.TrimSpace(drawString), " ")

			// get count
			count, err := strconv.ParseInt(draw[0], 10, 64)
			if err != nil {
				panic("failed to parse count of drawn cubes from string: '" + draw[0] + "'")
			}

			switch draw[1] {
			case "red":
				drawnCubes.Red = int(count)
			case "green":
				drawnCubes.Green = int(count)
			case "blue":
				drawnCubes.Blue = int(count)
			default:
				panic("WHO PUT THIS HERE?! => " + draw[1] + " <=")
			}
		}

		draws[idx] = drawnCubes
	}
	return int(gameId), draws
}
