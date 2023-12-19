package day18

import (
	"slices"
	"strconv"
	"strings"

	"github.com/Kryszak/aoc2023/common"
	"github.com/Kryszak/aoc2023/common/direction"
)

func loadDataPart2(path string) (steps []step) {
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := strings.Fields(fileScanner.Text())[2]
		line = line[2 : len(line)-1]
		count, _ := strconv.ParseInt(line[0:len(line)-1], 16, 64)
		var dir direction.Direction
		switch line[len(line)-1] {
		case '0':
			dir = direction.East
		case '1':
			dir = direction.South
		case '2':
			dir = direction.West
		case '3':
			dir = direction.North
		}
		steps = append(steps, step{dir, int(count)})
	}

	return steps
}

func getOutlinePoints(steps []step) (points []point) {
	x, y := 0, 0
	points = append(points, point{x, y})

	for _, step := range steps {
		switch step.dir {
		case direction.North:
			for i := x; i >= x-step.count; i-- {
				points = append(points, point{i, y})
			}
			x = x - step.count
		case direction.South:
			for i := x; i <= x+step.count; i++ {
				points = append(points, point{i, y})
			}
			x = x + step.count
		case direction.East:
			for j := y; j <= y+step.count; j++ {
				points = append(points, point{x, j})
			}
			y = y + step.count
		case direction.West:
			for j := y; j >= y-step.count; j-- {
				points = append(points, point{x, j})
			}
			y = y - step.count
		}
	}

	return points
}

func shoelace(points []point) int {
	sum := 0

	p0 := points[len(points)-1]
	for _, p1 := range points {
		sum += p0.y*p1.x - p0.x*p1.y
		p0 = p1
	}

	return sum/2 + len(points)/2 + 1
}

func Part2(path string) (answer int) {
	steps := loadDataPart2(path)
	points := getOutlinePoints(steps)
	points = slices.Compact(points)

	answer = shoelace(points)

	return answer
}
