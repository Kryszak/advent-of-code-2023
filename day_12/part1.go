package day12

import (
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

const (
	operational = '.'
	damaged     = '#'
	unknown     = '?'
)

type spring struct {
	sequence     string
	brokenGroups []int
}

func loadEntries(path string, unfoldCount int) []spring {
	var records []spring
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		initialSequence := line[0]
		sequence := line[0]
		var groups []int
		for _, group := range strings.Split(line[1], ",") {
			value, _ := strconv.Atoi(group)
			groups = append(groups, value)
		}
		initialGroups := groups
		for i := 0; i < unfoldCount-1; i++ {
			sequence = sequence + string(unknown) + initialSequence
			groups = append(groups, initialGroups...)
		}
		records = append(records, spring{sequence, groups})
	}
	return records
}

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

func Part1(path string) int {
	answer := 0

	records := loadEntries(path, 1)

	for _, spring := range records {
		answer += calculatePossibleArrangements(spring)
	}

	return answer
}
