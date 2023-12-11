package day03

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

var numbersRegexp = regexp.MustCompile(`[0-9]+`)

func loadEngineSchematic(path string) (engineSchematic []string) {
	fileScanner := common.FileScanner(path)

	for i := 0; fileScanner.Scan(); i++ {
		line := fileScanner.Text()
		engineSchematic = append(engineSchematic, line)
	}

	return engineSchematic
}

func Part1(path string) int {
	answer := 0

	engineSchematic := loadEngineSchematic(path)
	lineLenght := len(engineSchematic[0])

	for lineIndex, line := range engineSchematic {
		matches := numbersRegexp.FindAllStringSubmatchIndex(line, -1)
		for _, numberIndexRange := range matches {
			var surroundingCharacters string
			numberValue, _ := strconv.Atoi(line[numberIndexRange[0]:numberIndexRange[1]])
			leftBound := int(math.Max(float64(numberIndexRange[0]-1), 0))
			rightBound := int(math.Min(float64(numberIndexRange[1]+1), float64(lineLenght)))
			if lineIndex > 0 {
				surroundingCharacters = engineSchematic[lineIndex-1][leftBound:rightBound]
			}
			if numberIndexRange[0] > 0 {
				surroundingCharacters += string(engineSchematic[lineIndex][leftBound])
			}
			if numberIndexRange[1] < len(engineSchematic[lineIndex]) {
				surroundingCharacters += string(engineSchematic[lineIndex][numberIndexRange[1]])
			}
			if lineIndex < len(engineSchematic)-1 {
				surroundingCharacters += string(engineSchematic[lineIndex+1][leftBound:rightBound])
			}
			if strings.ReplaceAll(surroundingCharacters, ".", "") != "" {
				answer += numberValue
			}
		}
	}
	return answer
}
