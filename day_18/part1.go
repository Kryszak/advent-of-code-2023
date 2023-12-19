package day18

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Kryszak/aoc2023/common"
	"github.com/Kryszak/aoc2023/common/direction"
)

type step struct {
	dir   direction.Direction
	count int
}

type point struct {
	x, y int
}

func loadData(path string) (steps []step) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		var dir direction.Direction
		count := common.Atoi(line[1])
		switch line[0] {
		case "R":
			dir = direction.East
		case "D":
			dir = direction.South
		case "L":
			dir = direction.West
		case "U":
			dir = direction.North
		}
		steps = append(steps, step{dir, count})
	}

	return steps
}

func printLagoon(lagoon [][]rune) {
	for _, rowValue := range lagoon {
		for _, colValue := range rowValue {
			fmt.Print(string(colValue))
		}
		fmt.Println()
	}
}

func getOutline(steps []step) [][]rune {
	trench := make(map[point]bool)
	var lagoon [][]rune

	x, y := 0, 0
	trench[point{x, y}] = true

	for _, step := range steps {
		switch step.dir {
		case direction.North:
			for i := x; i >= x-step.count; i-- {
				trench[point{i, y}] = true
			}
			x = x - step.count
		case direction.South:
			for i := x; i <= x+step.count; i++ {
				trench[point{i, y}] = true
			}
			x = x + step.count
		case direction.East:
			for j := y; j <= y+step.count; j++ {
				trench[point{x, j}] = true
			}
			y = y + step.count
		case direction.West:
			for j := y; j >= y-step.count; j-- {
				trench[point{x, j}] = true
			}
			y = y - step.count
		}
	}

	var height, width int
	var heightOffset, widthOffset int

	for key := range trench {
		if key.x > height {
			height = key.x
		}
		if key.x < heightOffset {
			heightOffset = key.x
		}
		if key.y > width {
			width = key.y
		}
		if key.y < widthOffset {
			widthOffset = key.y
		}
	}

	height = height - heightOffset + 1
	width = width - widthOffset + 1

	for i := 0; i < height; i++ {
		lagoon = append(lagoon, make([]rune, width))
		for j := 0; j < len(lagoon[i]); j++ {
			if _, ok := trench[point{i - common.Abs(heightOffset), j - common.Abs(widthOffset)}]; ok {
				lagoon[i][j] = '#'
			} else {
				lagoon[i][j] = '.'
			}
		}
	}
	return lagoon
}

func getStartPoint(lagoon [][]rune) (startX, startY int) {
	for i, row := range lagoon {
		first := slices.Index(row, '#')
		var last int
		for j := len(row) - 1; j > first; j-- {
			if row[j] == '#' {
				last = j
				break
			}
		}
		for j := first; j < last; j++ {
			if row[j] == '.' {
				startX = i
				startY = j
				return startX, startY
			}
		}
	}
	return 0, 0
}

func floodFill(startX, startY int, lagoon [][]rune) {
	x := startX
	y := startY

	if lagoon[x][y] == '#' {
		return
	}

	if lagoon[x][y] == '.' {
		lagoon[x][y] = '#'
	}

	if x > 0 {
		floodFill(x-1, y, lagoon)
	}
	if x < len(lagoon)-1 {
		floodFill(x+1, y, lagoon)
	}
	if y > 0 {
		floodFill(x, y-1, lagoon)
	}
	if y < len(lagoon[x])-1 {
		floodFill(x, y+1, lagoon)
	}
}

func calculateLagoonScore(lagoon [][]rune) (result int) {
	for _, row := range lagoon {
		for _, column := range row {
			if column == '#' {
				result++
			}
		}
	}
	return result
}

func Part1(path string) (answer int) {
	steps := loadData(path)
	lagoon := getOutline(steps)

	startX, startY := getStartPoint(lagoon)
	floodFill(startX, startY, lagoon)

	answer = calculateLagoonScore(lagoon)

	return answer
}
