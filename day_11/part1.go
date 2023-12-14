package day11

import (
	"math"

	"github.com/Kryszak/aoc2023/common"
)

type location struct {
	x, y int
}

func loadUniverse(path string) (universe [][]rune, galaxies []location) {
	fileScanner := common.FileScanner(path)

	for x := 0; fileScanner.Scan(); x++ {
		line := make([]rune, 0)
		for y, character := range fileScanner.Text() {
			line = append(line, character)
			if character == '#' {
				galaxies = append(galaxies, location{x, y})
			}
		}
		universe = append(universe, line)
	}

	return universe, galaxies
}

func isEmptyLine(line *[]rune) bool {
	for _, character := range *line {
		if character != '.' {
			return false
		}
	}
	return true
}

func getExpandedColumnsAndRows(universe *[][]rune) (emptyRows, emptyColumns []int) {
	for y := 0; y < len((*universe)[0]); y++ {
		var line []rune
		for x := 0; x < len(*universe); x++ {
			line = append(line, (*&*universe)[x][y])
		}
		if isEmptyLine(&line) {
			emptyColumns = append(emptyColumns, y)
		}
	}
	for x := 0; x < len(*universe); x++ {
		if isEmptyLine(&(*universe)[x]) {
			emptyRows = append(emptyRows, x)
		}
	}

	return emptyRows, emptyColumns
}

func getEmptyBetween(first, second int, emptyIndexes []int) (count int) {
	minVal := common.Min(first, second)
	maxVal := common.Max(first, second)
	for _, value := range emptyIndexes {
		if value > minVal && value < maxVal {
			count++
		}
	}
	return count
}

func calculateGalaxiesDistance(first, second location, expandedRows, expandedColumns []int, expansionFactor int) int {
	expandTimes := expansionFactor - 1
	xDistance := int(
		math.Abs(float64(first.x-second.x)) +
			float64(expandTimes)*float64(getEmptyBetween(first.x, second.x, expandedRows)))

	yDistance := int(
		math.Abs(float64(first.y-second.y)) +
			float64(expandTimes)*float64(getEmptyBetween(first.y, second.y, expandedColumns)))

	return xDistance + yDistance
}

func Part1(path string) int {
	answer := 0

	universe, galaxies := loadUniverse(path)
	expandedRows, expandedColumns := getExpandedColumnsAndRows(&universe)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pathLength := calculateGalaxiesDistance(galaxies[i], galaxies[j], expandedRows, expandedColumns, 2)
			answer += pathLength
		}
	}

	return answer
}
