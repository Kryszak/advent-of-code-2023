package day10

import (
	"fmt"
	"unicode/utf8"

	"github.com/Kryszak/aoc2023/common"
)

const (
	start = 'S'
	north = 'N'
	south = 'S'
	east  = 'E'
	west  = 'W'
)

type point struct {
	x, y      int
	pipe      rune
	direction rune
}

type pointCandidate struct {
	x, y      int
	direction rune
}

func loadInput() (lines []string) {
	fileScanner := common.FileScanner("day_10/input.txt")

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func convertCharacter(raw byte) rune {
	character, _ := utf8.DecodeRune([]byte{raw})
	return character
}

func getPointAt(maze []string, x int, y int, direction rune) (point, error) {
	if x < 0 || x == len(maze) || y < 0 || y == len(maze[x]) {
		return point{}, fmt.Errorf("Out of bounds")
	}
	return point{x, y, convertCharacter(maze[x][y]), direction}, nil
}

func locateStartPointCoordinates(maze []string) (p point) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == start {
				return point{i, j, convertCharacter(maze[i][j]), north}
			}
		}
	}
	return point{-1, -1, 'q', 'N'}
}

func getPossiblePathStarts(startingPoint point, maze []string) (points []point) {
	var pointsAround []point

	pointCandidates := []pointCandidate{
		{startingPoint.x - 1, startingPoint.y, north},
		{startingPoint.x + 1, startingPoint.y, south},
		{startingPoint.x, startingPoint.y - 1, west},
		{startingPoint.x, startingPoint.y + 1, east},
	}

	for _, candidate := range pointCandidates {
		p, err := getPointAt(maze, candidate.x, candidate.y, candidate.direction)
		if err == nil && p.pipe != '.' {
			pointsAround = append(pointsAround, p)
		}
	}

	return pointsAround
}

func getNextPoint(maze []string, currentPoint point) point {
	var nextPoint point
	var x, y int
	var nextDirection rune
	direction := currentPoint.direction
	switch currentPoint.pipe {
	case '|':
		if direction == north {
			x = currentPoint.x - 1
		} else {
			x = currentPoint.x + 1
		}
		y = currentPoint.y
		nextDirection = direction
	case '-':
		if direction == east {
			y = currentPoint.y + 1
		} else {
			y = currentPoint.y - 1
		}
		x = currentPoint.x
		nextDirection = direction
	case 'L':
		if direction == west {
			x = currentPoint.x - 1
			y = currentPoint.y
			nextDirection = north
		} else {
			x = currentPoint.x
			y = currentPoint.y + 1
			nextDirection = east
		}
	case 'J':
		if direction == east {
			x = currentPoint.x - 1
			y = currentPoint.y
			nextDirection = north
		} else {
			x = currentPoint.x
			y = currentPoint.y - 1
			nextDirection = west
		}
	case '7':
		if direction == east {
			x = currentPoint.x + 1
			y = currentPoint.y
			nextDirection = south
		} else {
			x = currentPoint.x
			y = currentPoint.y - 1
			nextDirection = west
		}
	case 'F':
		if direction == north {
			x = currentPoint.x
			y = currentPoint.y + 1
			nextDirection = east
		} else {
			x = currentPoint.x + 1
			y = currentPoint.y
			nextDirection = south
		}
	}

	nextPoint, _ = getPointAt(maze, x, y, nextDirection)
	return nextPoint
}

func findLoop(maze []string) (loopPoints []point) {
	startingPoint := locateStartPointCoordinates(maze)
	pointsAround := getPossiblePathStarts(startingPoint, maze)

	for _, p := range pointsAround {
		loop := []point{startingPoint, p}
		current := p

		for current.pipe != start {
			current = getNextPoint(maze, current)
			if current.pipe == '.' {
				break
			}
			loop = append(loop, current)
		}

		if current.pipe == start {
			loopPoints = loop
		}
	}

	return loopPoints
}

func Part1() int {
	maze := loadInput()
	loop := findLoop(maze)

	answer := len(loop) / 2

	return answer
}
