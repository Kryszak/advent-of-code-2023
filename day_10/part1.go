package day10

import (
	"fmt"
	"unicode/utf8"

	"github.com/Kryszak/aoc2023/common"
)

const (
	start = 'S'
)

type point struct {
	x, y      int
	pipe      rune
	direction common.Direction
}

type pointCandidate struct {
	x, y      int
	direction common.Direction
}

func loadInput(path string) (lines []string) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func convertCharacter(raw byte) rune {
	character, _ := utf8.DecodeRune([]byte{raw})
	return character
}

func getPointAt(maze []string, x int, y int, direction common.Direction) (point, error) {
	if x < 0 || x == len(maze) || y < 0 || y == len(maze[x]) {
		return point{}, fmt.Errorf("Out of bounds")
	}
	return point{x, y, convertCharacter(maze[x][y]), direction}, nil
}

func locateStartPointCoordinates(maze []string) (p point) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == start {
				return point{i, j, convertCharacter(maze[i][j]), common.North}
			}
		}
	}
	return point{-1, -1, 'q', 'N'}
}

func getPossiblePathStarts(startingPoint point, maze []string) (points []point) {
	var pointsAround []point

	pointCandidates := []pointCandidate{
		{startingPoint.x - 1, startingPoint.y, common.North},
		{startingPoint.x + 1, startingPoint.y, common.South},
		{startingPoint.x, startingPoint.y - 1, common.West},
		{startingPoint.x, startingPoint.y + 1, common.East},
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
	var nextDirection common.Direction
	direction := currentPoint.direction
	switch currentPoint.pipe {
	case '|':
		if direction == common.North {
			x = currentPoint.x - 1
		} else {
			x = currentPoint.x + 1
		}
		y = currentPoint.y
		nextDirection = direction
	case '-':
		if direction == common.East {
			y = currentPoint.y + 1
		} else {
			y = currentPoint.y - 1
		}
		x = currentPoint.x
		nextDirection = direction
	case 'L':
		if direction == common.West {
			x = currentPoint.x - 1
			y = currentPoint.y
			nextDirection = common.North
		} else {
			x = currentPoint.x
			y = currentPoint.y + 1
			nextDirection = common.East
		}
	case 'J':
		if direction == common.East {
			x = currentPoint.x - 1
			y = currentPoint.y
			nextDirection = common.North
		} else {
			x = currentPoint.x
			y = currentPoint.y - 1
			nextDirection = common.West
		}
	case '7':
		if direction == common.East {
			x = currentPoint.x + 1
			y = currentPoint.y
			nextDirection = common.South
		} else {
			x = currentPoint.x
			y = currentPoint.y - 1
			nextDirection = common.West
		}
	case 'F':
		if direction == common.North {
			x = currentPoint.x
			y = currentPoint.y + 1
			nextDirection = common.East
		} else {
			x = currentPoint.x + 1
			y = currentPoint.y
			nextDirection = common.South
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

func Part1(path string) (answer int) {
	maze := loadInput(path)
	loop := findLoop(maze)

	answer = len(loop) / 2

	return answer
}
