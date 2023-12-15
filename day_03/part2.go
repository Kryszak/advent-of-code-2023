package day03

import (
	"regexp"

	"github.com/Kryszak/aoc2023/common"
)

func getStarIndexesInLine(line string) (starIndexes []int) {
	starMatches := regexp.MustCompile(`\*`).FindAllStringSubmatchIndex(line, -1)
	for _, match := range starMatches {
		starIndexes = append(starIndexes, match[0])
	}

	return starIndexes
}

func extractNumberNearStarInAnotherLine(line string, numberRanges [][]int, starIndex int) []int {
	var numbers []int
	for _, numberRange := range numberRanges {
		leftBound := common.Max(0, numberRange[0]-1)
		rightBound := common.Min(numberRange[1], len(line)-1)
		if leftBound <= starIndex && rightBound >= starIndex {
			numberValue := common.Atoi(line[numberRange[0]:numberRange[1]])
			numbers = append(numbers, numberValue)
		}
	}
	return numbers
}

func extractNumberNearStarInSameLine(line string, numberRanges [][]int, starIndex int) []int {
	var numbers []int
	for _, numberRange := range numberRanges {
		if numberRange[0] == starIndex+1 || numberRange[1] == starIndex {
			numberValue := common.Atoi(line[numberRange[0]:numberRange[1]])
			numbers = append(numbers, numberValue)
		}
	}
	return numbers
}

func Part2(path string) int {
	answer := 0

	engineSchematic := loadEngineSchematic(path)

	var gearRatioParts [][]int

	for lineIndex, line := range engineSchematic {
		starIndexes := getStarIndexesInLine(line)
		if len(starIndexes) > 0 {
			var numbersInLineAbove, numbersInLine, numbersInLineBelow [][]int
			if lineIndex > 0 {
				numbersInLineAbove = numbersRegexp.FindAllStringSubmatchIndex(engineSchematic[lineIndex-1], -1)
			}
			numbersInLine = numbersRegexp.FindAllStringSubmatchIndex(line, -1)
			if lineIndex < len(engineSchematic) {
				numbersInLineBelow = numbersRegexp.FindAllStringSubmatchIndex(engineSchematic[lineIndex+1], -1)
			}
			for _, index := range starIndexes {
				var starNeighbours []int
				starNeighbours = append(starNeighbours, extractNumberNearStarInAnotherLine(engineSchematic[lineIndex-1], numbersInLineAbove, index)...)
				starNeighbours = append(starNeighbours, extractNumberNearStarInAnotherLine(engineSchematic[lineIndex+1], numbersInLineBelow, index)...)
				starNeighbours = append(starNeighbours, extractNumberNearStarInSameLine(line, numbersInLine, index)...)
				if len(starNeighbours) == 2 {
					gearRatioParts = append(gearRatioParts, starNeighbours)
				}
			}
		}
	}
	for _, ratio := range gearRatioParts {
		answer += ratio[0] * ratio[1]
	}
	return answer
}
