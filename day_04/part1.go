package day04

import (
	"slices"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func parseNumbers(rawLine string) []string {
	return strings.Fields(rawLine)
}

func splitWinningAndCardNumbers(cardLine string) (winningNumbers []string, cardNumbers []string) {
	splitCard := strings.Split(cardLine, "|")

	winningNumbers = parseNumbers(splitCard[0])
	cardNumbers = parseNumbers(splitCard[1])
	return winningNumbers, cardNumbers
}

func calculateCardValue(winningNumbers []string, cardNumbers []string) (cardScore int) {
	for _, winningNumber := range winningNumbers {
		if slices.Contains(cardNumbers, winningNumber) {
			if cardScore == 0 {
				cardScore += 1
			} else {
				cardScore = cardScore * 2
			}
		}
	}
	return cardScore
}

func Part1(path string) (answer int) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ":")

		winningNumbers, cardNumbers := splitWinningAndCardNumbers(split[1])

		cardScore := calculateCardValue(winningNumbers, cardNumbers)
		answer += cardScore
	}

	return answer
}
