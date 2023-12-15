package day06

import (
	"bufio"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func parseSingleValueLine(fileScanner *bufio.Scanner) int {
	fileScanner.Scan()
	parsed := common.Atoi(strings.Join(strings.Fields(strings.Split(fileScanner.Text(), ":")[1]), ""))

	return parsed
}

func Part2(path string) int {
	answer := 0

	fileScanner := common.FileScanner(path)

	raceTime := parseSingleValueLine(fileScanner)
	raceDistance := parseSingleValueLine(fileScanner)

	for j := 1; j < raceTime; j++ {
		distanceTravelled := calculateDistance(j, raceTime)
		if distanceTravelled > raceDistance {
			answer++
		}
	}

	return answer
}
