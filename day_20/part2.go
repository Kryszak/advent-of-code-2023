package day20

import (
	"slices"

	"github.com/Kryszak/aoc2023/common"
)

func allValuesArePositive(values map[string]int) bool {
	for _, value := range values {
		if !(value > 0) {
			return false
		}
	}
	return true
}

func countRequiredIterations(machines map[string]module) int {
	rxSender := ""
	for key, module := range machines {
		for _, target := range module.outputs {
			if target == "rx" {
				rxSender = key
				break
			}
		}
	}

	watched := make([]string, 0)
	for _, target := range machines[rxSender].inputs {
		watched = append(watched, target.name)
	}

	firstHighPulse := make(map[string]int)
	for _, key := range watched {
		firstHighPulse[key] = 0
	}

	i := 0
	for !allValuesArePositive(firstHighPulse) {
		i++
		signals := processPulse(machines, &pulseCounter{0, 0})
		for _, signal := range signals {
			if signal.signal && slices.Contains(watched, signal.sender) && firstHighPulse[signal.sender] == 0 {
				firstHighPulse[signal.sender] = i
			}
		}
	}
	values := make([]int, 0)
	for _, value := range firstHighPulse {
		values = append(values, value)
	}

	return common.Lcm(values)
}

func Part2(path string) (answer int) {
	machines := loadData(path)

	answer = countRequiredIterations(machines)

	return answer
}
