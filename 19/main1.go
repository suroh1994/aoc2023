package main

import (
	aoc "aoc2023/aoccommon"
	"fmt"
	"strings"
)

type Workflow struct {
	Rules         []Rule
	DefaultTarget string
}

type Rule struct {
	Attribute rune
	Operation rune
	Value     int
	Target    string
}

type MachinePart struct {
	Values map[rune]int
}

type WorkflowManual map[string]Workflow

func (w WorkflowManual) ApplyWorkflows(part MachinePart) bool {
	workflowName := "in"
	workflow := w[workflowName]

	for {
		rules := workflow.Rules
		ruleIdx := 0
		for ; ruleIdx < len(rules); ruleIdx++ {
			if (rules[ruleIdx].Operation == '>' && part.Values[rules[ruleIdx].Attribute] > rules[ruleIdx].Value) ||
				(rules[ruleIdx].Operation == '<' && part.Values[rules[ruleIdx].Attribute] < rules[ruleIdx].Value) {
				//fmt.Printf("rule {%c%c%d} applied to %v\n", rules[ruleIdx].Attribute, rules[ruleIdx].Operation, rules[ruleIdx].Value, part)
				workflowName = rules[ruleIdx].Target
				break
			}
			//fmt.Printf("rule {%c%c%d} did not apply to %v\n", rules[ruleIdx].Attribute, rules[ruleIdx].Operation, rules[ruleIdx].Value, part)
		}

		if ruleIdx == len(rules) {
			//fmt.Printf("appling default rule to %v\n", part)
			workflowName = workflow.DefaultTarget
		}

		if workflowName == "A" {
			return true
		} else if workflowName == "R" {
			return false
		} else {
			workflow = w[workflowName]
		}
	}
}

func main() {
	lines := aoc.ReadInputLineByLine("input")
	idx := 0

	workflowManual := WorkflowManual{
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

		parsedRules := make([]Rule, len(rules)-1)
		for ruleIdx, rule := range rules[:len(rules)-1] {
			sepIdx := strings.IndexRune(rule, ':')
			parsedRules[ruleIdx] = Rule{
				Attribute: rune(rule[0]),
				Operation: rune(rule[1]),
				Value:     aoc.MustParseInt(rule[2:sepIdx]),
				Target:    rule[sepIdx+1:],
			}
		}
		workflowManual[name] = Workflow{
			Rules:         parsedRules,
			DefaultTarget: rules[len(rules)-1][:len(rules[len(rules)-1])-1],
		}
	}

	// skip empty line
	idx++

	// parse machine parts
	machineParts := make([]MachinePart, 0, len(lines[idx:]))
	for ; idx < len(lines); idx++ {
		segments := strings.Split(lines[idx], ",")
		machineParts = append(machineParts, MachinePart{
			Values: map[rune]int{
				'x': aoc.MustParseInt(segments[0][3:]),
				'm': aoc.MustParseInt(segments[1][2:]),
				'a': aoc.MustParseInt(segments[2][2:]),
				's': aoc.MustParseInt(segments[3][2 : len(segments[3])-1]),
			},
		})
	}

	acceptedParts := make([]MachinePart, 0, len(machineParts))
	for _, part := range machineParts {
		if workflowManual.ApplyWorkflows(part) {
			acceptedParts = append(acceptedParts, part)
		}
	}
	//fmt.Println(acceptedParts)

	// calculate solution
	var total int64 = 0
	for _, part := range acceptedParts {
		for _, value := range part.Values {
			total += int64(value)
		}
	}

	fmt.Println(total)
}
