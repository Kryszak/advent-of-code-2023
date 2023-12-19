package day10

import (
	"fmt"
	"unicode/utf8"

	"github.com/Kryszak/aoc2023/common"
	"github.com/Kryszak/aoc2023/common/direction"
)

const (
	start = 'S'
)

type point struct {
	x, y      int
	pipe      rune
	dir direction.Direction
}

type pointCandidate struct {
	x, y      int
	dir direction.Direction
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

func getPointAt(maze []string, x int, y int, dir direction.Direction) (point, error) {
	if x < 0 || x == len(maze) || y < 0 || y == len(maze[x]) {
		return point{}, fmt.Errorf("Out of bounds")
	}
	return point{x, y, convertCharacter(maze[x][y]), dir}, nil
}

func locateStartPointCoordinates(maze []string) (p point) {
	for i := 0; i < len(maze); i++ {
		for j := 0; j < len(maze[i]); j++ {
			if maze[i][j] == start {
				return point{i, j, convertCharacter(maze[i][j]), direction.North}
			}
		}
	}
	return point{-1, -1, 'q', 'N'}
}

func getPossiblePathStarts(startingPoint point, maze []string) (points []point) {
	var pointsAround []point

	pointCandidates := []pointCandidate{
		{startingPoint.x - 1, startingPoint.y, direction.North},
		{startingPoint.x + 1, startingPoint.y, direction.South},
		{startingPoint.x, startingPoint.y - 1, direction.West},
		{startingPoint.x, startingPoint.y + 1, direction.East},
	}

	for _, candidate := range pointCandidates {
		p, err := getPointAt(maze, candidate.x, candidate.y, candidate.dir)
		if err == nil && p.pipe != '.' {
			pointsAround = append(pointsAround, p)
		}
	}

	return pointsAround
}

func getNextPoint(maze []string, currentPoint point) point {
	var nextPoint point
	var x, y int
	var nextDirection direction.Direction
	dir := currentPoint.dir
	switch currentPoint.pipe {
	case '|':
		if dir == direction.North {
			x = currentPoint.x - 1
		} else {
			x = currentPoint.x + 1
		}
		y = currentPoint.y
		nextDirection = dir
	case '-':
		if dir == direction.East {
			y = currentPoint.y + 1
		} else {
			y = currentPoint.y - 1
		}
		x = currentPoint.x
		nextDirection = dir
	case 'L':
		if dir == direction.West {
			x = currentPoint.x - 1
			y = currentPoint.y
			nextDirection = direction.North
		} else {
			x = currentPoint.x
			y = currentPoint.y + 1
			nextDirection = direction.East
		}
	case 'J':
		if dir == direction.East {
			x = currentPoint.x - 1
			y = currentPoint.y
			nextDirection = direction.North
		} else {
			x = currentPoint.x
			y = currentPoint.y - 1
			nextDirection = direction.West
		}
	case '7':
		if dir == direction.East {
			x = currentPoint.x + 1
			y = currentPoint.y
			nextDirection = direction.South
		} else {
			x = currentPoint.x
			y = currentPoint.y - 1
			nextDirection = direction.West
		}
	case 'F':
		if dir == direction.North {
			x = currentPoint.x
			y = currentPoint.y + 1
			nextDirection = direction.East
		} else {
			x = currentPoint.x + 1
			y = currentPoint.y
			nextDirection = direction.South
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
