package day06

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func parseLine(fileScanner *bufio.Scanner) []int {
	var lineValues []int

	fileScanner.Scan()
	for _, value := range strings.Fields(strings.Split(fileScanner.Text(), ":")[1]) {
		parsed, _ := strconv.Atoi(value)
		lineValues = append(lineValues, parsed)
	}

	return lineValues
}

func calculateDistance(timeHold int, raceTime int) int {
	return timeHold * (raceTime - timeHold)
}

func Part1() int {
	answer := 1

	fileScanner := common.FileScanner("day_06/input.txt")

	raceTimes := parseLine(fileScanner)
	raceDistances := parseLine(fileScanner)

	for i := 0; i < len(raceTimes); i++ {
		raceTime := raceTimes[i]
		raceDistanceToBeat := raceDistances[i]
		numberOfWaysToBeatRecord := 0
		for j := 1; j < raceTime; j++ {
			distanceTravelled := calculateDistance(j, raceTime)
			if distanceTravelled > raceDistanceToBeat {
				numberOfWaysToBeatRecord++
			}
		}
		answer *= numberOfWaysToBeatRecord
	}

	return answer
}
