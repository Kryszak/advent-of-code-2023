package day16

import (
	"slices"

	"github.com/Kryszak/aoc2023/common"
)

func Part2(path string) (answer int) {
	var energizedTilesCount []int
	cavern := loadInput(path)

	for y := 0; y < len(cavern[0]); y++ {
		cavernCopy := common.Copy(cavern)
		visitMap := make(map[string]bool)
		markBeams(cavernCopy, visitMap, 0, y, common.South)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))

		cavernCopy = common.Copy(cavern)
		visitMap = make(map[string]bool)
		markBeams(cavernCopy, visitMap, len(cavernCopy)-1, y, common.North)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))
	}

	for x := 0; x < len(cavern); x++ {
		cavernCopy := common.Copy(cavern)
		visitMap := make(map[string]bool)
		markBeams(cavernCopy, visitMap, x, 0, common.East)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))

		cavernCopy = common.Copy(cavern)
		visitMap = make(map[string]bool)
		markBeams(cavernCopy, visitMap, x, len(cavern[x])-1, common.West)
		energizedTilesCount = append(energizedTilesCount, sumEnergizedTiles(cavernCopy))
	}

	answer = slices.Max(energizedTilesCount)

	return answer
}
