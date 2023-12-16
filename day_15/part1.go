package day15

import (
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func loadInput(path string) [][]rune {
	var input [][]rune
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		for _, value := range strings.Split(line, ",") {
			input = append(input, []rune(value))
		}
	}

	return input
}

func calculateStepHash(step []rune) int {
	result := 0
	for _, character := range step {
		result += int(character)
		result *= 17
		result %= 256
	}
	return result
}

func Part1(path string) (answer int) {
	input := loadInput(path)

	for _, input := range input {
		answer += calculateStepHash(input)
	}

	return answer
}
