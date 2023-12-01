package day09

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func loadReadings(fileScanner *bufio.Scanner) [][]int {
	var readings [][]int

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		var readingLine []int
		for _, value := range line {
			converted, _ := strconv.Atoi(value)
			readingLine = append(readingLine, converted)
		}
		readings = append(readings, readingLine)
	}

	return readings
}

func isZerosArray(array []int) bool {
	for _, value := range array {
		if value != 0 {
			return false
		}
	}
	return true
}

func calculateLineDifferences(line []int) [][]int {
	var diffLines [][]int
	diffLines = append(diffLines, line)
	reachedEnd := false

	for !reachedEnd {
		currentLine := diffLines[len(diffLines)-1]
		var diffs []int
		for i := 0; i < len(currentLine)-1; i++ {
			diffs = append(diffs, currentLine[i+1]-currentLine[i])
		}
		if isZerosArray(diffs) {
			reachedEnd = true
		} else {
			diffLines = append(diffLines, diffs)
		}
	}

	return diffLines
}

func predictNextValue(readingLine []int) int {
	diffLines := calculateLineDifferences(readingLine)

	for i := len(diffLines) - 1; i > 0; i-- {
		predictedValue := diffLines[i-1][len(diffLines[i-1])-1] + diffLines[i][len(diffLines[i])-1]
		diffLines[i-1] = append(diffLines[i-1], predictedValue)
	}

	return diffLines[0][len(diffLines[0])-1]
}

func Part1() int {
	answer := 0

	fileScanner := common.FileScanner("day_09/input.txt")
	readings := loadReadings(fileScanner)

	for _, readingLine := range readings {
		answer += predictNextValue(readingLine)
	}

	return answer
}
