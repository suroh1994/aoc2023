package main

import (
	"aoc2023/aoccommon"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	records := parseInput(inputLines)
	fmt.Println(records)

	productRecordBreakers := 1
	for _, record := range records {
		productRecordBreakers *= calculateNumberOfPossibleRecordBreakers(record)
	}

	err := aoccommon.WriteOutputLineByLine("output", productRecordBreakers)
	if err != nil {
		panic("oh crap")
	}
}

type RaceRecord struct {
	Time     int
	Distance int
}

/*
Time:      7  15   30
Distance:  9  40  200
*/
func parseInput(lines []string) []RaceRecord {
	if len(lines) != 2 {
		panic("somethings off, we got more than two lines!")
	}

	expr := regexp.MustCompile(`\d+`)
	timeStrings := expr.FindAllString(lines[0], -1)
	distanceStrings := expr.FindAllString(lines[1], -1)

	if len(timeStrings) != len(distanceStrings) {
		panic(fmt.Sprintf("something was mismatched; len(times)=%d but len(distances)=%d", len(timeStrings), len(distanceStrings)))
	}

	records := make([]RaceRecord, len(timeStrings))
	for idx := range timeStrings {
		time, err := strconv.Atoi(timeStrings[idx])
		if err != nil {
			panic("time didn't parse: " + timeStrings[idx])
		}
		distance, err := strconv.Atoi(distanceStrings[idx])
		if err != nil {
			panic("distance didn't parse: " + timeStrings[idx])
		}
		records[idx] = RaceRecord{Time: time, Distance: distance}
	}

	return records
}

func calculateNumberOfPossibleRecordBreakers(record RaceRecord) int {
	count := 0
	for i := 0; i < record.Time; i++ {
		if calculateDistanceTraveled(i, record.Time) > record.Distance {
			count++
		}
	}
	return count
}

func calculateDistanceTraveled(timeHeldDown, totalRaceDuration int) int {
	return (totalRaceDuration - timeHeldDown) * timeHeldDown
}
