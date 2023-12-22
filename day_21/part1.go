package day21

import (
	"fmt"

	"github.com/Kryszak/aoc2023/common"
)

type point struct {
	x, y int
}

func loadGarden(path string) [][]rune {
	var garden [][]rune
	scanner := common.FileScanner(path)

	for scanner.Scan() {
		row := []rune(scanner.Text())
		garden = append(garden, row)
	}

	return garden
}

func printGarden(garden [][]rune) {
	for _, row := range garden {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func getStartingPoint(garden [][]rune) (x, y int) {
	for x, row := range garden {
		for y, col := range row {
			if col == 'S' {
				return x, y
			}
		}
	}
	panic("Starting point not found in garden!")
}

func isInside(garden [][]rune, p point) bool {
	return p.x >= 0 && p.x < len(garden) && p.y >= 0 && p.y < len(garden[p.x])
}

func step(garden [][]rune, visited map[point]bool) map[point]bool {
	visitedPoints := make(map[point]bool)

	for p := range visited {
		for _, candidate := range [4]point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}} {
			if isInside(garden, candidate) && garden[candidate.x][candidate.y] != '#' {
				visitedPoints[candidate] = true
			}
		}
	}

	return visitedPoints
}

func Part1(path string) (answer int) {
	garden := loadGarden(path)
	x, y := getStartingPoint(garden)

	visited := make(map[point]bool)
	visited[point{x, y}] = true

	// change for 64 for real-data solution
	for i := 0; i < 6; i++ {
		visited = step(garden, visited)
	}

	answer = len(visited)

	return answer
}
