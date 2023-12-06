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

func (M MappingTriple) SourceEnd() int {
	return M[1] + M[2]
}

func (M MappingTriple) DestinationStart() int {
	return M[0]
}

func (M MappingTriple) DestinationEnd() int {
	return M[0] + M[2]
}

func (M MappingTriple) RangeLength() int {
	return M[2]
}

type Mapping struct {
	Source      string
	Destination string
	Triples     []MappingTriple
}

func (M Mapping) MapValuesToDestValues(sourceValues []int) []int {
	input := make([][2]int, len(sourceValues))

	for idx, value := range sourceValues {
		input[idx] = [2]int{value, 1}
	}

	output := M.MapValueRangesToDestValueRanges(input)
	result := make([]int, len(output))
	for idx, out := range output {
		result[idx] = out[0]
	}

	return result
}

func (M Mapping) MapValueRangesToDestValueRanges(sourceValues [][2]int) [][2]int {
	output := make([][2]int, 0, len(sourceValues))
	slices.SortFunc(M.Triples, MappingTripleSortBySourceStart)

	for _, sourceValue := range sourceValues {
		tripleIdx := 0
		for offset := 0; offset < sourceValue[1]; {
			startValue := 0
			count := 0
			// no more mappings, either no more triples or the next triple starts after the end of this group, keep the rest
			if (tripleIdx >= len(M.Triples)) || (sourceValue[0]+offset < M.Triples[tripleIdx].SourceStart()) {
				startValue = sourceValue[0] + offset
				count = sourceValue[1] - offset
			} else {
				startValue = M.Triples[tripleIdx].DestinationStart() + (sourceValue[0] - M.Triples[tripleIdx].SourceStart()) + offset
				// include up to the end of the mapping or all remaining numbers, depending on what's less
				count = min(
					M.Triples[tripleIdx].SourceEnd()-(sourceValue[0]+offset),
					sourceValue[1]-offset,
				)
			}

			// only add if count is positive (which means the current offset value is inside the mapped range)
			if count > 0 {
				output = append(output, [2]int{startValue, count})
				offset += count
			}
			tripleIdx++
		}
	}

	return output
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

func MappingTripleSortBySourceStart(a, b MappingTriple) int {
	return a.SourceStart() - b.SourceStart()
}
