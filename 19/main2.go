package main

import (
	"aoc2023/19/common"
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

type Limit struct {
	LowerBound, UpperBound int
}

func NewLimit() Limit {
	return Limit{
		LowerBound: 0,
		UpperBound: 4001,
	}
}

type MachinePartRange struct {
	Accepted bool
	Limits   map[rune]Limit
}

func NewMachinePartRange() MachinePartRange {
	return MachinePartRange{
		Accepted: false,
		Limits: map[rune]Limit{
			'a': NewLimit(),
			'm': NewLimit(),
			's': NewLimit(),
			'x': NewLimit(),
		},
	}
}

func CloneMachinePartRange(partRange MachinePartRange) MachinePartRange {
	newPartRange := NewMachinePartRange()
	for k, v := range partRange.Limits {
		newPartRange.Limits[k] = v
	}
	return newPartRange
}

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

	// combine rules
	ranges := GenerateRanges(workflowManual, "in", NewMachinePartRange())
	for rangeIdx := range ranges {
		fmt.Println(ranges[rangeIdx])
	}

	// calculate solution
	var total int64 = 0
	for _, partRange := range ranges {
		if !partRange.Accepted {
			continue
		}

		var partCount int64 = 1
		for _, limit := range partRange.Limits {
			partCount *= int64(limit.UpperBound - limit.LowerBound - 1)
		}
		total += partCount
	}

	fmt.Println(total)
}

// TODO Ranges are not being calculated correctly
func GenerateRanges(manual common.WorkflowManual, workflowName string, initialRange MachinePartRange) []MachinePartRange {
	if workflowName == "R" {
		initialRange.Accepted = false
		return []MachinePartRange{initialRange}
	} else if workflowName == "A" {
		initialRange.Accepted = true
		return []MachinePartRange{initialRange}
	}

	workflow := manual[workflowName]
	response := make([]MachinePartRange, 0)
	for _, rule := range workflow.Rules {
		newRange := CloneMachinePartRange(initialRange)
		limit := newRange.Limits[rule.Attribute]
		initialLimit := initialRange.Limits[rule.Attribute]
		if rule.Operation == '>' {
			limit.LowerBound = max(limit.LowerBound, rule.Value)
			initialLimit.UpperBound = min(initialLimit.UpperBound, rule.Value+1)
		} else {
			limit.UpperBound = min(limit.UpperBound, rule.Value)
			initialLimit.LowerBound = max(initialLimit.LowerBound, rule.Value-1)
		}
		newRange.Limits[rule.Attribute] = limit
		initialRange.Limits[rule.Attribute] = initialLimit

		response = append(response, GenerateRanges(manual, rule.Target, newRange)...)
	}

	// default rule
	response = append(response, GenerateRanges(manual, workflow.DefaultTarget, initialRange)...)

	return response
}
