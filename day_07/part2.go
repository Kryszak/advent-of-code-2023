package day07

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

var cardOrderWithJokerRule = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

func evaluateCardHandTypeWithJokers(hand string) string {
	characterCountMap := make(map[string]int)

	for _, character := range hand {
		characterCountMap[string(character)]++
	}

	jokerCount := characterCountMap["J"]
	delete(characterCountMap, "J")

	var counts []int

	for _, value := range characterCountMap {
		counts = append(counts, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	var highestFigureCount int
	var secondFigureCount int

	if len(counts) == 0 {
		highestFigureCount = jokerCount
	} else {
		highestFigureCount = counts[0] + jokerCount
	}

	if len(counts) > 1 {
		secondFigureCount = counts[1]
	} else {
		secondFigureCount = 0
	}

	if highestFigureCount == 5 {
		return five_of_kind
	}

	if highestFigureCount == 4 {
		return four_of_kind
	}

	if highestFigureCount == 3 && secondFigureCount == 2 {
		return full_house
	}

	if highestFigureCount == 3 && secondFigureCount != 2 {
		return three_of_kind
	}

	if highestFigureCount == 2 && secondFigureCount == 2 {
		return two_pair
	}

	if highestFigureCount == 2 && secondFigureCount != 2 {
		return pair
	}

	return high_card
}

func loadCardsWithJokers(fileScanner *bufio.Scanner) []card {
	var cards []card

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		hand := line[0]
		cardBid, _ := strconv.Atoi(line[1])
		handType := evaluateCardHandTypeWithJokers(hand)
		card := card{hand, cardBid, handType}
		cards = append(cards, card)
	}

	return cards
}

func Part2() int {
	answer := 0

	fileScanner := common.FileScanner("day_07/input.txt")
	cards := loadCardsWithJokers(fileScanner)

	sort.Slice(cards, func(i, j int) bool {
		firstIndex := getIndexOfElement(typeOrder, cards[i].handType)
		secondIndex := getIndexOfElement(typeOrder, cards[j].handType)
		if firstIndex == secondIndex {
			return compareSameFigures(cardOrderWithJokerRule, cards[i], cards[j])
		}
		return firstIndex > secondIndex
	})

	for index, card := range cards {
		answer += card.bid * (index + 1)
	}

	return answer
}
