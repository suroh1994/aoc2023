package common

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
