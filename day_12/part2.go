package day12

func Part2(path string) int {
	answer := 0
	records := loadEntries(path, 5)

	for _, spring := range records {
		answer += calculatePossibleArrangements(spring)
	}

	return answer
}
