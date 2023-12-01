package day01

import (
	"strconv"

	"github.com/Kryszak/aoc2023/common"
)

func Part1() int {
	answer := 0
	fileScanner := common.FileScanner("day_01/input.txt")

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var digits []string
		for _, value := range line {
			character := string(value)
			if _, err := strconv.Atoi(character); err == nil {
				digits = append(digits, character)
			}
		}
		lineValue, _ := strconv.Atoi(digits[0] + digits[len(digits)-1])
		answer += lineValue
	}

	return answer
}
