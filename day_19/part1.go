package day19

import (
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

type rule struct {
	operand   string
	operator  rune
	threshold int
	moveTo    string
}

type workflow struct {
	name  string
	rules []rule
}

func loadInput(path string) (workflows map[string][]rule, partRatings []map[string]int) {
	workflows = make(map[string][]rule)
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}
		nameRulesSplit := strings.Split(line, "{")
		name := nameRulesSplit[0]
		var rules []rule
		for _, rawRule := range strings.Split(nameRulesSplit[1][:len(nameRulesSplit[1])-1], ",") {
			split := strings.Split(rawRule, ":")
			if len(split) > 1 {
				rules = append(rules, rule{string(split[0][0]), rune(split[0][1]), common.Atoi(split[0][2:]), split[1]})
			} else {
				rules = append(rules, rule{"", ' ', 0, split[0]})
			}
		}
		workflows[name] = rules
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		line = line[1 : len(line)-1]
		ratings := make(map[string]int)
		for _, rating := range strings.Split(line, ",") {
			split := strings.Split(rating, "=")
			ratings[split[0]] = common.Atoi(split[1])
		}
		partRatings = append(partRatings, ratings)
	}

	return workflows, partRatings
}

func evaluatePart(rating map[string]int) int {
	result := 0
	for _, value := range rating {
		result += value
	}
	return result
}

func getRuleResult(rating map[string]int, r rule) string {
	if r.operator == ' ' {
		return r.moveTo
	}

	if r.operator == '>' {
		if rating[r.operand] > r.threshold {
			return r.moveTo
		}
	}

	if r.operator == '<' {
		if rating[r.operand] < r.threshold {
			return r.moveTo
		}
	}

	return ""
}

func processPart(rating map[string]int, workflows map[string][]rule, workflowName string) int {
	evaluated := false
	result := ""
	nextRuleSet := workflowName
	for !evaluated {
		currentRuleSet := workflows[nextRuleSet]
		for _, r := range currentRuleSet {
			result = getRuleResult(rating, r)
			if result == "" {
				continue
			}
			if result == "A" || result == "R" {
				evaluated = true
			}
			nextRuleSet = result
			break
		}
	}
	if result == "A" {
		return evaluatePart(rating)
	}
	return 0
}

func Part1(path string) (answer int) {
	workflows, partRatings := loadInput(path)

	for _, rating := range partRatings {
		value := processPart(rating, workflows, "in")
		answer += value
	}
	return answer
}
