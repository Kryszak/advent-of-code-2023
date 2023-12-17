package day16

import (
	"fmt"

	"github.com/Kryszak/aoc2023/common"
)

type tile struct {
	character  rune
	energized  bool
	visitCount int
}

func loadInput(path string) [][]tile {
	var cavern [][]tile

	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		var line []tile
		for _, value := range []rune(fileScanner.Text()) {
			line = append(line, tile{value, false, 0})
		}
		cavern = append(cavern, line)
	}

	return cavern
}

func printCavern(cavern [][]tile) {
	for x := 0; x < len(cavern); x++ {
		for y := 0; y < len(cavern[x]); y++ {
			if cavern[x][y].energized {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func markBeams(cavern [][]tile, visitMap map[string]bool, x, y int, dir common.Direction) {
	if x < 0 || x >= len(cavern) || y < 0 || y >= len(cavern[x]) {
		return
	}
	tile := &cavern[x][y]

	key := fmt.Sprintf("(%d,%d) -> %d", x, y, dir)

	if _, ok := visitMap[key]; ok {
		return
	} else {
		visitMap[key] = true
	}

	currentField := tile.character
	tile.energized = true
	tile.visitCount++
	switch currentField {
	case '.':
		{
			switch dir {
			case common.North:
				markBeams(cavern, visitMap, x-1, y, common.North)
			case common.South:
				markBeams(cavern, visitMap, x+1, y, common.South)
			case common.East:
				markBeams(cavern, visitMap, x, y+1, common.East)
			case common.West:
				markBeams(cavern, visitMap, x, y-1, common.West)
			}
		}
	case '|':
		{
			switch dir {
			case common.North:
				markBeams(cavern, visitMap, x-1, y, common.North)
			case common.South:
				markBeams(cavern, visitMap, x+1, y, common.South)
			case common.East, common.West:
				markBeams(cavern, visitMap, x-1, y, common.North)
				markBeams(cavern, visitMap, x+1, y, common.South)
			}
		}
	case '-':
		{
			switch dir {
			case common.North, common.South:
				markBeams(cavern, visitMap, x, y-1, common.West)
				markBeams(cavern, visitMap, x, y+1, common.East)
			case common.East:
				markBeams(cavern, visitMap, x, y+1, common.East)
			case common.West:
				markBeams(cavern, visitMap, x, y-1, common.West)
			}
		}
	case '/':
		{
			switch dir {
			case common.North:
				markBeams(cavern, visitMap, x, y+1, common.East)
			case common.South:
				markBeams(cavern, visitMap, x, y-1, common.West)
			case common.East:
				markBeams(cavern, visitMap, x-1, y, common.North)
			case common.West:
				markBeams(cavern, visitMap, x+1, y, common.South)
			}
		}
	case '\\':
		{
			switch dir {
			case common.North:
				markBeams(cavern, visitMap, x, y-1, common.West)
			case common.South:
				markBeams(cavern, visitMap, x, y+1, common.East)
			case common.East:
				markBeams(cavern, visitMap, x+1, y, common.South)
			case common.West:
				markBeams(cavern, visitMap, x-1, y, common.North)
			}
		}
	}
}

func sumEnergizedTiles(cavern [][]tile) (result int) {
	for x := 0; x < len(cavern); x++ {
		for y := 0; y < len(cavern[x]); y++ {
			if cavern[x][y].energized {
				result++
			}
		}
	}
	return
}

func Part1(path string) (answer int) {
	cavern := loadInput(path)
	visitMap := make(map[string]bool)

	markBeams(cavern, visitMap, 0, 0, common.East)
	answer = sumEnergizedTiles(cavern)
	return answer
}
