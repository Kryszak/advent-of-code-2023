package day13

func getArrayDifference(first []int, second []int) []int {
	m := make(map[int]bool)

	for _, val := range first {
		m[val] = true
	}

	result := make([]int, 0)

	for _, val := range second {
		if _, ok := m[val]; !ok {
			result = append(result, val)
		}
	}

	return result
}

func evaluateReflectionsWithSmudges(pattern []string) int {
	result := 0

	rows := pattern
	columns := getColumnsFrom(pattern)

	initialRowIndexes := findReflectionIndexes(rows)
	initialColumnIndexes := findReflectionIndexes(columns)

	found := false

	for i := 0; i < len(pattern) && !found; i++ {
		for j := 0; j < len(pattern[i]) && !found; j++ {
			withSmudgeRemoved := make([]string, len(pattern))
			copy(withSmudgeRemoved, pattern)
			character := withSmudgeRemoved[i][j]
			var swapped rune
			if character == '.' {
				swapped = '#'
			} else {
				swapped = '.'
			}
			rowToChange := withSmudgeRemoved[i]
			out := []rune(rowToChange)
			out[j] = swapped
			withSmudgeRemoved[i] = string(out)

			newRowIndexes := getArrayDifference(initialRowIndexes, findReflectionIndexes(withSmudgeRemoved))
			newColumnIndexes := getArrayDifference(initialColumnIndexes, findReflectionIndexes(getColumnsFrom(withSmudgeRemoved)))

			if indexes := newRowIndexes; len(indexes) > 0 {
				result += indexes[0] * 100
			}
			if indexes := newColumnIndexes; len(indexes) > 0 {
				result += indexes[0] * 1
			}

			if result > 0 {
				found = true
			}
		}
	}

	return result
}

func Part2(path string) (answer int) {
	patterns := loadPatterns(path)

	for _, pattern := range patterns {
		answer += evaluateReflectionsWithSmudges(pattern)
	}

	return answer
}
