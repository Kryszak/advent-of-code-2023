package day14

import (
	"fmt"

	"github.com/Kryszak/aoc2023/common"
)

const (
	roundedRock = 'O'
	cubeRock    = '#'
	empty       = '.'
)

func loadDish(path string) [][]rune {
	var dish [][]rune
	fileScanner := common.FileScanner(path)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		dish = append(dish, []rune(line))
	}

	return dish
}

func printDish(dish [][]rune) {
	for x := 0; x < len(dish); x++ {
		for y := 0; y < len(dish[x]); y++ {
			fmt.Print(string(dish[x][y]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func copyDish(dish [][]rune) [][]rune {
	copied := make([][]rune, len(dish))
	for i := range dish {
		copied[i] = make([]rune, len(dish[i]))
		copy(copied[i], dish[i])
	}
	return copied
}

func tiltNorth(dish [][]rune) [][]rune {
	tilted := copyDish(dish)
	for y := 0; y < len(tilted[0]); y++ {
		for x := 1; x < len(tilted); x++ {
			if tilted[x][y] == roundedRock {
				i := 0
				for x-i-1 >= 0 && tilted[x-i-1][y] == empty {
					tilted[x-i-1][y], tilted[x-i][y] = tilted[x-i][y], tilted[x-i-1][y]
					i++
				}
			}
		}
	}
	return tilted
}

func calculateLoad(dish [][]rune) int {
	result := 0
	for y := 0; y < len(dish[0]); y++ {
		for x := 0; x < len(dish); x++ {
			if dish[x][y] == roundedRock {
				result += len(dish) - x
			}
		}
	}
	return result
}

func Part1(path string) (answer int) {
	dish := loadDish(path)
	dish = tiltNorth(dish)

	answer = calculateLoad(dish)

	return answer
}
