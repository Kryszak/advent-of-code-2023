package day05

import (
	"bufio"
	"math"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

func parseLineReversed(valueMap *[]sourceToDestination, line string) {
	values := strings.Fields(line)
	source := common.Atoi(values[0])
	destination := common.Atoi(values[1])
	rng := common.Atoi(values[2])
	entry := new(sourceToDestination)
	entry.source = source
	entry.destination = destination
	entry.rng = rng
	*valueMap = append(*valueMap, *entry)
}

func parseDestToSource(valueMap *[]sourceToDestination, fileScanner *bufio.Scanner) {
	line := fileScanner.Text()
	for line != "" {
		fileScanner.Scan()
		line = fileScanner.Text()
		if line == "" {
			break
		}
		parseLineReversed(valueMap, line)
	}
}

func seedExists(seeds []int, seed int) bool {
	for i := 0; i < len(seeds); i += 2 {
		if seed >= seeds[i] && seed < seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}

func Part2(path string) (answer int) {
	fileScanner := common.FileScanner(path)

	var seeds []int
	soilToSeed := make([]sourceToDestination, 0)
	fertilizerToSoil := make([]sourceToDestination, 0)
	waterToFertilizer := make([]sourceToDestination, 0)
	lightToWater := make([]sourceToDestination, 0)
	temperatureToLight := make([]sourceToDestination, 0)
	humidityToTemperature := make([]sourceToDestination, 0)
	locationToHumidity := make([]sourceToDestination, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "seeds:") {
			seedRanges := strings.Fields(strings.TrimSpace(strings.Split(line, ":")[1]))
			for _, seed := range seedRanges {
				value := common.Atoi(seed)
				seeds = append(seeds, value)
			}
		}
		if strings.HasPrefix(line, "seed-to") {
			parseDestToSource(&soilToSeed, fileScanner)
		}
		if strings.HasPrefix(line, "soil-to") {
			parseDestToSource(&fertilizerToSoil, fileScanner)
		}
		if strings.HasPrefix(line, "fertilizer-to") {
			parseDestToSource(&waterToFertilizer, fileScanner)
		}
		if strings.HasPrefix(line, "water-to") {
			parseDestToSource(&lightToWater, fileScanner)
		}
		if strings.HasPrefix(line, "light-to") {
			parseDestToSource(&temperatureToLight, fileScanner)
		}
		if strings.HasPrefix(line, "temperature-to") {
			parseDestToSource(&humidityToTemperature, fileScanner)
		}
		if strings.HasPrefix(line, "humidity-to") {
			parseDestToSource(&locationToHumidity, fileScanner)
		}
	}

	for i := 0; i < math.MaxInt; i++ {
		humidity := getOrDefault(locationToHumidity, i)
		temperature := getOrDefault(humidityToTemperature, humidity)
		light := getOrDefault(temperatureToLight, temperature)
		water := getOrDefault(lightToWater, light)
		fertilizer := getOrDefault(waterToFertilizer, water)
		soil := getOrDefault(fertilizerToSoil, fertilizer)
		seed := getOrDefault(soilToSeed, soil)

		if seedExists(seeds, seed) {
			answer = i
			break
		}
	}

	return answer
}
