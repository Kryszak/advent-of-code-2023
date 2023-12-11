package day04

import (
	"slices"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func calculateWonCards(winningNumbers []string, cardNumbers []string) (cardScore int) {
	for _, winningNumber := range winningNumbers {
		if slices.Contains(cardNumbers, winningNumber) {
			cardScore += 1
		}
	}
	return cardScore
}

func Part2(path string) int {
	answer := 0

	fileScanner := common.FileScanner(path)
	var scratchCards []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, ":")[1]

		scratchCards = append(scratchCards, split)
	}

	cardWeight := make(map[int]int, len(scratchCards))
	for i := 0; i < len(scratchCards); i++ {
		cardWeight[i] = 1
	}

	for cardIndex, card := range scratchCards {
		numberOfCurrentCards := cardWeight[cardIndex]

		winningNumbers, cardNumbers := splitWinningAndCardNumbers(card)
		cardValue := calculateWonCards(winningNumbers, cardNumbers)

		lastWonCardIndex := cardIndex + cardValue

		processCount := 0
		for processCount < numberOfCurrentCards {
			for i := cardIndex + 1; i < len(scratchCards) && i <= lastWonCardIndex; i++ {
				cardWeight[i] = cardWeight[i] + 1
			}
			processCount++
		}

		answer += cardWeight[cardIndex]
	}

	return answer
}
