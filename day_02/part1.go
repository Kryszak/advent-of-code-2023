package day02

import (
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

const redCount = 12
const greenCount = 13
const blueCount = 14

func Part1(path string) (answer int) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		isValid := true
		id, roundValues := parseIdAndRoundValues(fileScanner.Text())

		for _, element := range roundValues {
			for _, colorCountRaw := range strings.Split(element, ",") {
				count, colorName := parseColorCountAndName(colorCountRaw)
				switch colorName {
				case "red":
					if count > redCount {
						isValid = false
					}
				case "green":
					if count > greenCount {
						isValid = false
					}
				case "blue":
					if count > blueCount {
						isValid = false
					}
				}
				if !isValid {
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			answer += id
		}
	}

	return answer
}
