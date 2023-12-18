package day18

import (
	"fmt"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

type step struct {
	dir   common.Direction
	count int
	color string
}

type point struct {
	x, y int
}

func loadData(path string) (steps []step) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())
		var direction common.Direction
		count := common.Atoi(line[1])
		switch line[0] {
		case "R":
			direction = common.East
		case "D":
			direction = common.South
		case "L":
			direction = common.West
		case "U":
			direction = common.North
		}
		color := strings.ReplaceAll(strings.ReplaceAll(line[2], ")", ""), "(", "")
		steps = append(steps, step{direction, count, color})
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

func Part1(path string) (answer int) {
	steps := loadData(path)

	// trench := make([]point, 0)

	trench := make(map[point]bool)

	x, y := 0, 0
	trench[point{x, y}] = true

	for _, step := range steps {
		switch step.dir {
		case common.North:
			for i := x; i >= x-step.count; i-- {
				trench[point{i, y}] = true
			}
			x = x - step.count
		case common.South:
			for i := x; i <= x+step.count; i++ {
				trench[point{i, y}] = true
			}
			x = x + step.count
		case common.East:
			for j := y; j <= y+step.count; j++ {
				trench[point{x, j}] = true
			}
			y = y + step.count
		case common.West:
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

	var lagoon [][]rune

	for x := 0; x < height; x++ {
		lagoon = append(lagoon, make([]rune, width))
		for y := 0; y < len(lagoon[x]); y++ {
			if _, ok := trench[point{x - heightOffset, y - widthOffset}]; ok {
				lagoon[x][y] = '#'
			} else {
				lagoon[x][y] = '.'
			}
		}
	}

	for _, row := range lagoon {
		var leftBound, rightBound int
		for y := 0; y < len(row); y++ {
			if row[y] == '#' {
				leftBound = y
				break
			}
		}
		for y := len(row) - 1; y >= 0; y-- {
			if row[y] == '#' {
				rightBound = y
				break
			}
		}
		for y := leftBound; y < rightBound; y++ {
			row[y] = '#'
		}
	}

	printLagoon(lagoon)

	for _, row := range lagoon {
		for _, column := range row {
			if column == '#' {
				answer++
			}
		}
	}

	return answer
}
