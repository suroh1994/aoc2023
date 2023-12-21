package main

import (
	"aoc2023/19/common"
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

func main() {
	lines := aoc.ReadInputLineByLine("input")
	idx := 0

	workflowManual := common.WorkflowManual{
		"A": {
			nil,
			"",
		},
		"R": {
			nil,
			"",
		},
	}

	// parse workflows
	for ; lines[idx] != ""; idx++ {
		segments := strings.Split(lines[idx], "{")
		name := segments[0]
		rules := strings.Split(segments[1][:len(segments[1])], ",")

		parsedRules := make([]common.Rule, len(rules)-1)
		for ruleIdx, rule := range rules[:len(rules)-1] {
			sepIdx := strings.IndexRune(rule, ':')
			parsedRules[ruleIdx] = common.Rule{
				Attribute: rune(rule[0]),
				Operation: rune(rule[1]),
				Value:     aoc.MustParseInt(rule[2:sepIdx]),
				Target:    rule[sepIdx+1:],
			}
		}
		workflowManual[name] = common.Workflow{
			Rules:         parsedRules,
			DefaultTarget: rules[len(rules)-1][:len(rules[len(rules)-1])-1],
		}
	}

	// skip empty line
	idx++

	// parse machine parts
	machineParts := make([]common.MachinePart, 0, len(lines[idx:]))
	for ; idx < len(lines); idx++ {
		segments := strings.Split(lines[idx], ",")
		machineParts = append(machineParts, common.MachinePart{
			Values: map[rune]int{
				'x': aoc.MustParseInt(segments[0][3:]),
				'm': aoc.MustParseInt(segments[1][2:]),
				'a': aoc.MustParseInt(segments[2][2:]),
				's': aoc.MustParseInt(segments[3][2 : len(segments[3])-1]),
			},
		})
	}

	acceptedParts := make([]common.MachinePart, 0, len(machineParts))
	for _, part := range machineParts {
		if workflowManual.ApplyWorkflows(part) {
			acceptedParts = append(acceptedParts, part)
		}
	}

	// calculate solution
	var total int64 = 0
	for _, part := range acceptedParts {
		for _, value := range part.Values {
			total += int64(value)
		}
	}

	fmt.Println(total)
}
