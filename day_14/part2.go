package day14

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func tiltSouth(dish [][]rune) [][]rune {
	tilted := copyDish(dish)
	for y := 0; y < len(tilted[0]); y++ {
		for x := len(tilted) - 1; x >= 0; x-- {
			if tilted[x][y] == roundedRock {
				i := 0
				for x+i+1 < len(tilted) && tilted[x+i+1][y] == empty {
					tilted[x+i+1][y], tilted[x+i][y] = tilted[x+i][y], tilted[x+i+1][y]
					i++
				}
			}
		}
	}
	return tilted
}

func tiltWest(dish [][]rune) [][]rune {
	tilted := copyDish(dish)
	for x := 0; x < len(tilted); x++ {
		for y := 1; y < len(tilted[x]); y++ {
			if tilted[x][y] == roundedRock {
				i := 0
				for y-i-1 >= 0 && tilted[x][y-i-1] == empty {
					tilted[x][y-i-1], tilted[x][y-i] = tilted[x][y-i], tilted[x][y-i-1]
					i++
				}
			}
		}
	}
	return tilted
}

func tiltEast(dish [][]rune) [][]rune {
	tilted := copyDish(dish)
	for x := 0; x < len(tilted); x++ {
		for y := len(tilted[x]) - 2; y >= 0; y-- {
			if tilted[x][y] == roundedRock {
				i := 0
				for y+i+1 < len(tilted[x]) && tilted[x][y+i+1] == empty {
					tilted[x][y+i+1], tilted[x][y+i] = tilted[x][y+i], tilted[x][y+i+1]
					i++
				}
			}
		}
	}
	return tilted
}

func hashDish(dish [][]rune) string {
	var stringBuilder strings.Builder
	for _, line := range dish {
		stringBuilder.WriteString(string(line))
	}

	hasher := sha256.New()
	hasher.Write([]byte(stringBuilder.String()))

	return hex.EncodeToString(hasher.Sum(nil))
}

func Part2(path string) (answer int) {
	hashedTilts := make(map[string]int)
	tiltedCombinations := make([][][]rune, 0)

	dish := loadDish(path)
	iterations := 1000000000

	iterationsBeforeRepeat := 0
	repeatPeriod := 0

	for i := 0; i < iterations; i++ {
		dish = tiltNorth(dish)
		dish = tiltWest(dish)
		dish = tiltSouth(dish)
		dish = tiltEast(dish)

		hash := hashDish(dish)

		if value, exists := hashedTilts[hash]; exists {
			iterationsBeforeRepeat = value
			repeatPeriod = i - iterationsBeforeRepeat
			break
		}
		hashedTilts[hash] = i
		tiltedCombinations = append(tiltedCombinations, copyDish(dish))
	}

	finalDishIndex := (iterations-iterationsBeforeRepeat)%repeatPeriod + iterationsBeforeRepeat - 1

	answer = calculateLoad(tiltedCombinations[finalDishIndex])

	return answer
}
