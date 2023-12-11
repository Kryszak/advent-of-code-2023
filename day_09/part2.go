package day09

import (
	"github.com/Kryszak/aoc2023/common"
)

func predictPreviousValue(readingLine []int) int {
	diffLines := calculateLineDifferences(readingLine)

	for i := len(diffLines) - 1; i > 0; i-- {
		predictedValue := diffLines[i-1][0] - diffLines[i][0]
		diffLines[i-1] = append([]int{predictedValue}, diffLines[i-1]...)
		diffLines[i-1] = append(diffLines[i-1], predictedValue)
	}

	return diffLines[0][0]
}

func Part2(path string) int {
	answer := 0

	fileScanner := common.FileScanner(path)
	readings := loadReadings(fileScanner)

	for _, readingLine := range readings {
		answer += predictPreviousValue(readingLine)
	}

	return answer
}
