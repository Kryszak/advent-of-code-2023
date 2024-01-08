package day22

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

type dimension struct {
	start, end int
}

type brick struct {
	x, y, z dimension
}

func loadData(path string) (bricks []brick) {
	scanner := common.FileScanner(path)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "~")
		startPoints := strings.Split(split[0], ",")
		endPoints := strings.Split(split[1], ",")
		xDim := dimension{common.Atoi(startPoints[0]), common.Atoi(endPoints[0])}
		yDim := dimension{common.Atoi(startPoints[1]), common.Atoi(endPoints[1])}
		zDim := dimension{common.Atoi(startPoints[2]), common.Atoi(endPoints[2])}

		b := brick{xDim, yDim, zDim}
		bricks = append(bricks, b)
	}

	return bricks
}

func sortBricksByDistanceFromGround(bricks []brick) {
	slices.SortFunc(bricks, func(a, b brick) int {
		return a.z.start - b.z.start
	})
}

func Part1(path string) (answer int) {
	bricks := loadData(path)

    sortBricksByDistanceFromGround(bricks)

	for _, b := range bricks {
		fmt.Printf("%+v\n", b)
	}

	return answer
}
