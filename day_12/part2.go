package day12

func dp(i, j int, record string, group []int, cache [][]int) int {
	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	result := 0
	if record[i] == operational {
		result = dp(i+1, j, record, group, cache)
	} else {
		if record[i] == unknown {
			result += dp(i+1, j, record, group, cache)
		}
		if j < len(group) {
			count := 0
			for k := i; k < len(record); k++ {
				if count > group[j] || record[k] == operational || count == group[j] && record[k] == unknown {
					break
				}
				count += 1
			}

			if count == group[j] {
				if i+count < len(record) && record[i+count] != damaged {
					result += dp(i+count+1, j+1, record, group, cache)
				} else {
					result += dp(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = result
	return result
}

func calculatePossibleArrangements(record spring) int {
	var cache [][]int
	for i := 0; i < len(record.sequence); i++ {
		cache = append(cache, make([]int, len(record.brokenGroups)+1))
		for j := 0; j < len(record.brokenGroups)+1; j++ {
			cache[i][j] = -1
		}
	}

	return dp(0, 0, record.sequence, record.brokenGroups, cache)
}

func Part2(path string) int {
	answer := 0
	records := loadEntries(path, 5)

	for _, spring := range records {
		answer += calculatePossibleArrangements(spring)
	}

	return answer
}
