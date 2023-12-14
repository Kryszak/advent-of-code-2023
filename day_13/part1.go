package day13

import (
	"github.com/Kryszak/aoc2023/common"
)

func loadPatterns(path string) (patterns [][]string) {
	fileScanner := common.FileScanner(path)

	var currentPattern []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			patterns = append(patterns, currentPattern)
			currentPattern = make([]string, 0)
		} else {
			currentPattern = append(currentPattern, line)
		}
	}
	patterns = append(patterns, currentPattern)

	return
}

func findReflectionIndexes(pattern []string) []int {
	var indexes []int
	for i := 1; i < len(pattern); i++ {
		limit := common.Min(i, len(pattern)-i)
		found := true
		for j := 0; j < limit && found; j++ {
			found = pattern[i-1-j] == pattern[i+j]
		}
		if found {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func getColumnsFrom(pattern []string) []string {
	var columns []string

	for i := 0; i < len(pattern[0]); i++ {
		var column string
		for j := 0; j < len(pattern); j++ {
			column += string(pattern[j][i])
		}
		columns = append(columns, column)
	}

	return columns
}

func evaluateReflections(pattern []string) int {
	result := 0

	rows := pattern
	columns := getColumnsFrom(pattern)

	if indexes := findReflectionIndexes(rows); len(indexes) > 0 {
		result += indexes[0] * 100
	}

	if indexes := findReflectionIndexes(columns); len(indexes) > 0 {
		result += indexes[0] * 1
	}

	return result
}

func Part1(path string) int {
	answer := 0

	patterns := loadPatterns(path)

	for _, pattern := range patterns {
		answer += evaluateReflections(pattern)
	}

	return answer
}
