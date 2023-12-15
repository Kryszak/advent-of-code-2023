package day01

import (
	"strconv"

	"github.com/Kryszak/aoc2023/common"
)

func Part1(path string) int {
	answer := 0
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var digits []string
		for _, value := range line {
			character := string(value)
			if _, err := strconv.Atoi(character); err == nil {
				digits = append(digits, character)
			}
		}
		lineValue := common.Atoi(digits[0] + digits[len(digits)-1])
		answer += lineValue
	}

	return answer
}
