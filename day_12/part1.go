package day12

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

const (
	operational = '.'
	damaged     = '#'
	unknown     = '?'
)

type spring struct {
	sequence     string
	brokenGroups []int
}

func loadEntries(path string, unfoldCount int) []spring {
	var records []spring
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " ")
		initialSequence := line[0]
		sequence := line[0]
		var groups []int
		for _, group := range strings.Split(line[1], ",") {
			value := common.Atoi(group)
			groups = append(groups, value)
		}
		initialGroups := groups
		for i := 0; i < unfoldCount-1; i++ {
			sequence = sequence + string(unknown) + initialSequence
			groups = append(groups, initialGroups...)
		}
		records = append(records, spring{sequence, groups})
	}
	return records
}

func generateCombinations(length int, current string, combinations []string) []string {
	if length == 0 {
		combinations = append(combinations, current)
		return combinations
	}

	combinations = generateCombinations(length-1, current+"0", combinations)
	combinations = generateCombinations(length-1, current+"1", combinations)

	return combinations
}

func generatePossibleSequences(entry spring) []string {
	unknownCount := strings.Count(entry.sequence, "?")
	combinations := []string{}
	possiblePermutations := generateCombinations(unknownCount, "", combinations)
	possibleSequences := make([]string, 0)
	for _, possiblePermutation := range possiblePermutations {
		sequence := entry.sequence
		for _, value := range possiblePermutation {
			var replaced string
			if value == '0' {
				replaced = string(operational)
			} else {
				replaced = string(damaged)
			}
			sequence = strings.Replace(sequence, string(unknown), replaced, 1)
		}
		possibleSequences = append(possibleSequences, sequence)
	}
	return possibleSequences
}

func isValidSequence(sequence string, locations []int) bool {
	regex := regexp.MustCompile("[#]+")
	indexes := regex.FindAllStringSubmatchIndex(sequence, -1)
	var matches []int
	for _, indexRange := range indexes {
		matches = append(matches, indexRange[1]-indexRange[0])
	}
	return reflect.DeepEqual(locations, matches)
}

func Part1(path string) int {
	answer := 0

	records := loadEntries(path, 1)

	for _, spring := range records {
		possibleSequences := generatePossibleSequences(spring)
		for _, sequence := range possibleSequences {
			if isValidSequence(sequence, spring.brokenGroups) {
				answer++
			}
		}
	}

	return answer
}
