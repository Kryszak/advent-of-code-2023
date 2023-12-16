package day16

import (
	"slices"
)

func copyCavern(cavern [][]tile) [][]tile {
	copied := make([][]tile, len(cavern))
	for i := range cavern {
		copied[i] = make([]tile, len(cavern[i]))
		copy(copied[i], cavern[i])
	}
	return copied
}

func Part2(path string) int {
	answer := 0

	var energizedTilesCount []int
	cavern := loadInput(path)

	for y := 0; y < len(cavern[0]); y++ {
		cavernCopy := copyCavern(cavern)
		visitMap := make(map[string]bool)
		markBeams(cavernCopy, visitMap, 0, y, south)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))

		cavernCopy = copyCavern(cavern)
		visitMap = make(map[string]bool)
		markBeams(cavernCopy, visitMap, len(cavernCopy)-1, y, north)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))
	}

	for x := 0; x < len(cavern); x++ {
		cavernCopy := copyCavern(cavern)
		visitMap := make(map[string]bool)
		markBeams(cavernCopy, visitMap, x, 0, east)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))

		cavernCopy = copyCavern(cavern)
		visitMap = make(map[string]bool)
		markBeams(cavernCopy, visitMap, x, len(cavern[x])-1, west)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))
	}

	answer = slices.Max(energizedTilesCount)

	return answer
}
