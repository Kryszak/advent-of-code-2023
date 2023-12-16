package day01

import (
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

var wordToDigitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func Part2(path string) (answer int) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		var digits []string
		for index, value := range line {
			character := string(value)
			if _, err := strconv.Atoi(character); err == nil {
				digits = append(digits, character)
			}
			for key, characterDigit := range wordToDigitMap {
				if strings.HasPrefix(line[index:], key) {
					digits = append(digits, characterDigit)
				}
			}
		}
		lineValue := common.Atoi(digits[0] + digits[len(digits)-1])
		answer += lineValue
	}

	return answer
}
