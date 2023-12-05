package common

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type MappingTriple [3]int

func (M MappingTriple) SourceStart() int {
	return M[1]
}

func (M MappingTriple) DestinationStart() int {
	return M[0]
}

func (M MappingTriple) RangeLength() int {
	return M[2]
}

type Mapping struct {
	Source      string
	Destination string
	Triples     []MappingTriple
}

var (
	regexpr = regexp.MustCompile("(\\w*)-to-(\\w*)\\smap:")
)

func ParseInput(lines []string) ([]int, []Mapping) {
	return getSeeds(lines[0]), getMappings(lines[2:])
}

func getSeeds(line string) []int {
	seedNumberStrings := strings.Split(line[7:], " ")
	seeds := make([]int, len(seedNumberStrings))
	for idx, seedNumberString := range seedNumberStrings {
		seedNumber, err := strconv.Atoi(seedNumberString)
		if err != nil {
			panic(fmt.Sprintf("ah crap: %v", err))
		}
		seeds[idx] = seedNumber
	}

	return seeds
}

func getMappings(lines []string) []Mapping {
	mappings := make([]Mapping, 0)
	currentMapping := &Mapping{}
	for _, line := range lines {
		// empty line terminates a mapping
		if strings.TrimSpace(line) == "" {
			mappings = append(mappings, *currentMapping)
			currentMapping = &Mapping{}
			continue
		}

		// if the mapping has no name, assume we are in a line with a mapping header
		if currentMapping.Source == "" {
			matches := regexpr.FindStringSubmatch(line)
			if len(matches) != 3 {
				panic(fmt.Sprintf("all good things come in three, not %d! (%q)", len(matches), lines))
			}
			currentMapping.Source = matches[1]
			currentMapping.Destination = matches[2]
			continue
		}

		mappingNumberStrings := strings.Split(line, " ")
		currentMapping.Triples = append(currentMapping.Triples, MappingTriple{
			parseNumber(mappingNumberStrings[0]),
			parseNumber(mappingNumberStrings[1]),
			parseNumber(mappingNumberStrings[2]),
		})
	}
	// add final mapping
	mappings = append(mappings, *currentMapping)

	return mappings
}

func parseNumber(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(fmt.Sprintf("failed to map %q to number: %v", a, err))
	}
	return i
}

func FindLocationsForSeeds(seeds []int, mappings []Mapping) []int {
	// sort all mappings by source
	for _, mapping := range mappings {
		slices.SortFunc(mapping.Triples, MappingTripleSortBySourceStart)
	}

	locations := make([]int, len(seeds))
	for idx, seed := range seeds {
		currentId := seed
		for _, mapping := range mappings {
			currentId = FindDestinationValue(currentId, mapping)
		}
		locations[idx] = currentId
	}
	return locations
}

func FindDestinationValue(sourceValue int, mapping Mapping) int {
	for idx := 0; idx < len(mapping.Triples); idx++ {
		// find the largest source smaller than seed (inc until source larger, go back one step)
		if mapping.Triples[idx].SourceStart() < sourceValue {
			continue
		}

		relevantTriple := mapping.Triples[idx-1]
		// check if within range
		if sourceValue <= relevantTriple.SourceStart()+relevantTriple.RangeLength() {
			// return mapped value
			offset := sourceValue - relevantTriple.SourceStart()
			return relevantTriple.DestinationStart() + offset
		}

		// we are between two mappings
		break
	}

	// we are either between mappings or beyond the greatest mapping, return unmapped value
	return sourceValue
}

func MappingTripleSortBySourceStart(a, b MappingTriple) int {
	return a.SourceStart() - b.SourceStart()
}

func MappingTripleSortByDestinationStart(a, b MappingTriple) int {
	return a.DestinationStart() - b.DestinationStart()
}
