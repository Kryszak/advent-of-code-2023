package day02

import (
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func setMax(current *int, compared int) {
	if *current < compared {
		*current = compared
	}
}

func Part2() int {
	answer := 0

	fileScanner := common.FileScanner("day_02/input.txt")

	for fileScanner.Scan() {
		_, roundValues := parseIdAndRoundValues(fileScanner.Text())
		var minRedCount, minGreenCount, minBlueCount int

		for _, element := range roundValues {
			for _, colorCountRaw := range strings.Split(element, ",") {
				count, colorName := parseColorCountAndName(colorCountRaw)
				switch colorName {
				case "red":
					setMax(&minRedCount, count)
				case "green":
					setMax(&minGreenCount, count)
				case "blue":
					setMax(&minBlueCount, count)
				}
			}
		}

		setPower := minRedCount * minBlueCount * minGreenCount
		answer += setPower
	}

	return answer
}
