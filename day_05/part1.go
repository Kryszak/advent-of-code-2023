package day05

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

type sourceToDestination struct {
	source      int
	destination int
	rng         int
}

func parseLine(valueMap *[]sourceToDestination, line string) {
	values := strings.Fields(line)
	source, _ := strconv.Atoi(values[1])
	destination, _ := strconv.Atoi(values[0])
	rng, _ := strconv.Atoi(values[2])
	entry := new(sourceToDestination)
	entry.source = source
	entry.destination = destination
	entry.rng = rng
	*valueMap = append(*valueMap, *entry)
}

func parseSourceToDest(valueMap *[]sourceToDestination, fileScanner *bufio.Scanner) {
	line := fileScanner.Text()
	for line != "" {
		fileScanner.Scan()
		line = fileScanner.Text()
		if line == "" {
			break
		}
		parseLine(valueMap, line)
	}
}

func getOrDefault(sourceToDestinationMappings []sourceToDestination, index int) int {
	for _, mapping := range sourceToDestinationMappings {
		if index >= mapping.source && index < mapping.source+mapping.rng {
			return mapping.destination + (index - mapping.source)
		}
	}
	return index
}

func Part1(path string) int {
	answer := 0

	fileScanner := common.FileScanner(path)

	var seeds []int
	seedToSoil := make([]sourceToDestination, 0)
	soilToFertilizer := make([]sourceToDestination, 0)
	fertilizerToWater := make([]sourceToDestination, 0)
	waterToLight := make([]sourceToDestination, 0)
	lightToTemperature := make([]sourceToDestination, 0)
	temperatureToHumidity := make([]sourceToDestination, 0)
	humidityToLocation := make([]sourceToDestination, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			for _, seed := range strings.Fields(strings.TrimSpace(strings.Split(line, ":")[1])) {
				seedValue, _ := strconv.Atoi(seed)
				seeds = append(seeds, seedValue)
			}
		}
		if strings.HasPrefix(line, "seed-to") {
			parseSourceToDest(&seedToSoil, fileScanner)
		}
		if strings.HasPrefix(line, "soil-to") {
			parseSourceToDest(&soilToFertilizer, fileScanner)
		}
		if strings.HasPrefix(line, "fertilizer-to") {
			parseSourceToDest(&fertilizerToWater, fileScanner)
		}
		if strings.HasPrefix(line, "water-to") {
			parseSourceToDest(&waterToLight, fileScanner)
		}
		if strings.HasPrefix(line, "light-to") {
			parseSourceToDest(&lightToTemperature, fileScanner)
		}
		if strings.HasPrefix(line, "temperature-to") {
			parseSourceToDest(&temperatureToHumidity, fileScanner)
		}
		if strings.HasPrefix(line, "humidity-to") {
			parseSourceToDest(&humidityToLocation, fileScanner)
		}
	}

	for _, seed := range seeds {
		soil := getOrDefault(seedToSoil, seed)
		fertilizer := getOrDefault(soilToFertilizer, soil)
		water := getOrDefault(fertilizerToWater, fertilizer)
		light := getOrDefault(waterToLight, water)
		temperature := getOrDefault(lightToTemperature, light)
		humidity := getOrDefault(temperatureToHumidity, temperature)
		location := getOrDefault(humidityToLocation, humidity)
		if answer == 0 {
			answer = location
		}
		if location < answer {
			answer = location
		}
	}

	return answer
}
