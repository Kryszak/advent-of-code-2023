package day19

func cloneMap(original map[string][2]int) map[string][2]int {
	c := make(map[string][2]int, 4)
	for k, v := range original {
		c[k] = v
	}
	return c
}

func checkWorkflowsForRanges(workflows map[string][]rule, workflowName string, values map[string][2]int) int {
	if workflowName == "R" {
		return 0
	}
	if workflowName == "A" {
		result := 1
		for _, v := range values {
			result *= (v[1] - v[0] + 1)
		}
		return result
	}

	total := 0

	for _, r := range workflows[workflowName] {
		v := values[r.operand]

		var trueSplit, falseSplit [2]int

		if r.operator == '<' {
			trueSplit = [2]int{v[0], r.threshold - 1}
			falseSplit = [2]int{r.threshold, v[1]}
		}
		if r.operator == '>' {
			trueSplit = [2]int{r.threshold + 1, v[1]}
			falseSplit = [2]int{v[0], r.threshold}
		}
		if r.operator == ' ' {
			total += checkWorkflowsForRanges(workflows, r.moveTo, values)
			continue
		}

		if trueSplit[0] <= trueSplit[1] {
			v2 := cloneMap(values)
			v2[r.operand] = trueSplit
			total += checkWorkflowsForRanges(workflows, r.moveTo, v2)
		}

		if falseSplit[0] > falseSplit[1] {
			break
		}

		values[r.operand] = falseSplit
	}

	return total
}

func Part2(path string) (answer int) {
	workflows, _ := loadInput(path)

	answer = checkWorkflowsForRanges(workflows, "in", map[string][2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	})

	return answer
}
