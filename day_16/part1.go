package day16

import (
	"fmt"

	"github.com/Kryszak/aoc2023/common"
)

type direction int

const (
	north direction = iota
	south
	east
	west
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

func markBeams(cavern [][]tile, visitMap map[string]bool, x, y int, dir direction) {
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
			case north:
				markBeams(cavern, visitMap, x-1, y, north)
			case south:
				markBeams(cavern, visitMap, x+1, y, south)
			case east:
				markBeams(cavern, visitMap, x, y+1, east)
			case west:
				markBeams(cavern, visitMap, x, y-1, west)
			}
		}
	case '|':
		{
			switch dir {
			case north:
				markBeams(cavern, visitMap, x-1, y, north)
			case south:
				markBeams(cavern, visitMap, x+1, y, south)
			case east, west:
				markBeams(cavern, visitMap, x-1, y, north)
				markBeams(cavern, visitMap, x+1, y, south)
			}
		}
	case '-':
		{
			switch dir {
			case north, south:
				markBeams(cavern, visitMap, x, y-1, west)
				markBeams(cavern, visitMap, x, y+1, east)
			case east:
				markBeams(cavern, visitMap, x, y+1, east)
			case west:
				markBeams(cavern, visitMap, x, y-1, west)
			}
		}
	case '/':
		{
			switch dir {
			case north:
				markBeams(cavern, visitMap, x, y+1, east)
			case south:
				markBeams(cavern, visitMap, x, y-1, west)
			case east:
				markBeams(cavern, visitMap, x-1, y, north)
			case west:
				markBeams(cavern, visitMap, x+1, y, south)
			}
		}
	case '\\':
		{
			switch dir {
			case north:
				markBeams(cavern, visitMap, x, y-1, west)
			case south:
				markBeams(cavern, visitMap, x, y+1, east)
			case east:
				markBeams(cavern, visitMap, x+1, y, south)
			case west:
				markBeams(cavern, visitMap, x-1, y, north)
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

	markBeams(cavern, visitMap, 0, 0, east)
	answer = sumEnergizedTiles(cavern)
	return answer
}
