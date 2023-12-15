package day07

import (
	"bufio"
	"sort"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

const (
	five_of_kind  = "five_of_kind"
	four_of_kind  = "four_of_kind"
	full_house    = "full_house"
	three_of_kind = "three_of_kind"
	two_pair      = "two_pair"
	pair          = "pair"
	high_card     = "high_card"
)

var typeOrder = []string{five_of_kind, four_of_kind, full_house, three_of_kind, two_pair, pair, high_card}
var cardOrder = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

type card struct {
	hand     string
	bid      int
	handType string
}

func evaluateCardHandType(hand string) string {
	characterCountMap := make(map[string]int)

	for _, character := range hand {
		characterCountMap[string(character)]++
	}

	var counts []int

	for _, value := range characterCountMap {
		counts = append(counts, value)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	if counts[0] == 5 {
		return five_of_kind
	}

	if counts[0] == 4 {
		return four_of_kind
	}

	if counts[0] == 3 && counts[1] == 2 {
		return full_house
	}

	if counts[0] == 3 && counts[1] != 2 {
		return three_of_kind
	}

	if counts[0] == 2 && counts[1] == 2 {
		return two_pair
	}

	if counts[0] == 2 && counts[1] != 2 {
		return pair
	}

	return high_card
}

func loadCards(fileScanner *bufio.Scanner) []card {
	var cards []card

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		hand := line[0]
		cardBid := common.Atoi(line[1])
		handType := evaluateCardHandType(hand)
		card := card{hand, cardBid, handType}
		cards = append(cards, card)
	}

	return cards
}

func getIndexOfElement(array []string, element string) int {
	for index, value := range array {
		if value == element {
			return index
		}
	}
	return -1
}

func compareSameFigures(cardWeights []string, firstCard card, secondCard card) bool {
	for i := 0; i < len(firstCard.hand); i++ {
		firstType := string(firstCard.hand[i])
		secondType := string(secondCard.hand[i])
		if firstType == secondType {
			continue
		}
		return getIndexOfElement(cardWeights, firstType) > getIndexOfElement(cardWeights, secondType)
	}
	return false
}

func Part1(path string) (answer int) {
	fileScanner := common.FileScanner(path)
	cards := loadCards(fileScanner)

	sort.Slice(cards, func(i, j int) bool {
		firstIndex := getIndexOfElement(typeOrder, cards[i].handType)
		secondIndex := getIndexOfElement(typeOrder, cards[j].handType)
		if firstIndex == secondIndex {
			return compareSameFigures(cardOrder, cards[i], cards[j])
		}
		return firstIndex > secondIndex
	})

	for index, card := range cards {
		answer += card.bid * (index + 1)
	}

	return
}
