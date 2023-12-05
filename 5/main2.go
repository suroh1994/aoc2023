package main

import (
	"aoc2023/5/common"
	"aoc2023/aoccommon"
	"fmt"
	"math"
	"slices"
)

func main() {
	inputLines := aoccommon.ReadInputLineByLine("input")

	seeds, mappings := common.ParseInput(inputLines)

	combined := mappings[0]
	addMissingMappings(&combined)
	for _, mapping := range mappings[1:] {
		combined = combineMappings(combined, mapping)
	}

	seedTuples := convertSeedsToTuples(seeds)
	location := findSmallestLocation(seedTuples, combined)

	err := aoccommon.WriteOutputLineByLine("output", location)
	if err != nil {
		panic("oh crap")
	}
}

func addMissingMappings(mapping *common.Mapping) {
	slices.SortFunc(mapping.Triples, common.MappingTripleSortBySourceStart)
	lastSourceValue := 0
	additionalTriples := make([]common.MappingTriple, 0)
	for _, triple := range mapping.Triples {
		if lastSourceValue < triple.SourceStart() {
			additionalTriples = append(additionalTriples, common.MappingTriple{lastSourceValue, lastSourceValue, triple.SourceStart() - lastSourceValue})
		}
		lastSourceValue = triple.SourceStart() + triple.RangeLength()
	}
	mapping.Triples = append(mapping.Triples, additionalTriples...)
}

/*
*
humidity2location: [

	{24, 5, 10}

]

0 => 0
1 => 1
2 => 2
3 => 3
4 => 4
5 => 24
6 => 25
7 => 26
8 => 27
9 => 28
10 => 29
11 => 30
12 => 31
13 => 32
14 => 33
15 => 15

temperature2humidity: [

	{3, 9, 6},
	{15, 21, 5}

]

_______________________________
0 => 0 => 0					   |
1 => 1 => 1					   |
2 => 2 => 2					   | Group 1
3 => 3 => 3					   |
4 => 4 => 4	___________________|
5 => 5 => 24				   |
6 => 6 => 25				   | Group 2
7 => 7 => 26				   |
8 => 8 => 27 __________________|
9 => 3 => 3					   | Group 3
10 => 4 => 4 __________________|
11 => 5 => 24				   |
12 => 6 => 25				   | Group 4
13 => 7 => 26 				   |
14 => 8 => 27 _________________|
15 => 15 => 15				   |
16 => 16 => 16				   | Group 5
...			  				   |
20 => 20 => 20 ________________|
21 => 15 => 15				   |
22 => 16 => 16				   | Group 6
23 => 17 => 17				   |
24 => 18 => 18				   |
25 => 19 => 19 ________________|
26 => 26 => 26				   |
...							   | Group 7

result: [

	{ 0,  0,  5},
	{24,  5,  4},
	{ 3,  9,  2}
	{24, 11,  4},
	{15, 15,  6},
	{15, 21,  5},

]

	{ 3,  9,  6} + {24,  5, 10} = { 0,  0,  5},
	{15, 21,  5}   { 0,  0,  5}	  {24,  5,  4},
								  { 3,  9,  2},
								  {24, 11,  4},
								  {15, 15,  6},
								  {15, 21,  5},
*/
func combineMappings(sourceMapping, destinationMapping common.Mapping) common.Mapping {
	if sourceMapping.Destination != destinationMapping.Source {
		panic(fmt.Sprintf("these don't match: %q != %q", sourceMapping.Destination, destinationMapping.Source))
	}

	slices.SortFunc(sourceMapping.Triples, common.MappingTripleSortBySourceStart)
	slices.SortFunc(destinationMapping.Triples, common.MappingTripleSortBySourceStart)

	// create an empty output mapping
	output := common.Mapping{
		Source:      sourceMapping.Source,
		Destination: destinationMapping.Destination,
		Triples:     make([]common.MappingTriple, 0),
	}
	// iterate over source mappings
	for _, triple := range sourceMapping.Triples {
		// check range the source mapping maps to (get all matching mappings)

		relevantDestMappings := findRelevantTriples(
			triple.DestinationStart(),
			triple.DestinationStart()+triple.RangeLength(),
			&destinationMapping.Triples,
		)

		// -- if no mapping exists in destination mapping, copy mapping to output & continue
		if len(relevantDestMappings) == 0 {
			output.Triples = append(output.Triples, triple)
			continue
		}

		slices.SortFunc(relevantDestMappings, common.MappingTripleSortBySourceStart)
		// -- if destination mapping maps to different values, create new mappings for all sub ranges
		//    and copy to output
		idx := 0
		offset := 0
		for offset < triple.RangeLength() {
			// case the mapped range goes beyond the last relevant mapping from the right side
			if idx >= len(relevantDestMappings) {
				count := triple.RangeLength() - offset
				output.Triples = append(output.Triples, common.MappingTriple{
					triple.DestinationStart() + offset,
					triple.SourceStart() + offset,
					count,
				})
				offset += count
				continue
			}
			// case the mapped range starts before the relevant mapping from the right side
			if triple.DestinationStart()+offset < relevantDestMappings[idx].SourceStart() {
				count := relevantDestMappings[idx].SourceStart() - triple.DestinationStart() + offset
				output.Triples = append(output.Triples, common.MappingTriple{
					triple.DestinationStart() + offset,
					triple.SourceStart() + offset,
					count,
				})
				offset += count
				continue
			}

			count := min(
				relevantDestMappings[idx].SourceStart()+relevantDestMappings[idx].RangeLength()-triple.DestinationStart()+offset,
				triple.RangeLength()-offset,
			)
			output.Triples = append(output.Triples, common.MappingTriple{
				relevantDestMappings[idx].DestinationStart(),
				triple.SourceStart() + offset,
				count,
			})

			offset += count
			idx++
		}
	}

	// finally, if the last mapped source value (start + offset) is less than the last mapped destination value,
	//    create a new mapping and copy to output

	return output
}

func findRelevantTriples(start, end int, triples *[]common.MappingTriple) []common.MappingTriple {
	relevantDestTriples := make([]common.MappingTriple, 0, len(*triples))
	for _, triple := range *triples {
		if !(triple.SourceStart() > end || triple.SourceStart()+triple.RangeLength()-1 < start) {
			relevantDestTriples = append(relevantDestTriples, triple)
		}
	}
	return relevantDestTriples
}

func convertSeedsToTuples(seedsSlice []int) [][2]int {
	if len(seedsSlice)%2 != 0 {
		panic("uneven number of values for seeds found!")
	}

	output := make([][2]int, len(seedsSlice)/2)
	for idx := 0; idx < len(seedsSlice); idx += 2 {
		output[idx/2] = [2]int{seedsSlice[idx], seedsSlice[idx+1]}
	}

	return output
}

func findSmallestLocation(seeds [][2]int, seed2LocationMapping common.Mapping) int {
	slices.SortFunc(seeds, func(a, b [2]int) int {
		return a[0] - b[0]
	})
	slices.SortFunc(seed2LocationMapping.Triples, common.MappingTripleSortByDestinationStart)

	smallestLocation := math.MaxInt
	for _, triple := range seed2LocationMapping.Triples {
		if triple.DestinationStart() > smallestLocation {
			break
		}

		for _, tuple := range seeds {
			// seeds are below source
			if triple.SourceStart() > tuple[0]+tuple[1] {
				continue
			}

			// seeds start after the targeted range
			if triple.SourceStart()+triple.RangeLength() < tuple[0] {
				// since we sorted the slice, there is no chance to find anything for this range
				break
			}

			// seeds may start before the sourceStart
			if tuple[0] < triple.SourceStart() {
				smallestLocation = triple.DestinationStart()
			} else {
				// seeds start inside possible range
				offset := tuple[0] - triple.SourceStart()
				smallestLocation = triple.DestinationStart() + offset
			}
		}
	}
	return smallestLocation
}
