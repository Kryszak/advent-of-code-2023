package day10

import "fmt"

func printVisualization(visualization [][]string) {
	for x := 0; x < len(visualization); x++ {
		for y := 0; y < len(visualization[x]); y++ {
			fmt.Print(visualization[x][y])
		}
		fmt.Println()
	}
}

func prepareMazeVisualization(maze []string, loopPoints []point) [][]string {
	visualization := make([][]string, len(maze))

	for i := 0; i < len(maze); i++ {
		visualization[i] = make([]string, len(maze[i]))
		for j := 0; j < len(maze[i]); j++ {
			visualization[i][j] = "."
		}
	}

	for _, p := range loopPoints {
		visualization[p.x][p.y] = string(p.pipe)
	}

	return visualization
}

func findTilesInsideCount(visualization [][]string) (tilesInsideCount int) {
	for x := 0; x < len(visualization); x++ {
		isInside := false
		previousCharacter := ""
		for y := 0; y < len(visualization[x]); y++ {
			nodeCharacter := visualization[x][y]
			if nodeCharacter != "." {
				if previousCharacter == "" {
					previousCharacter = nodeCharacter
				}
				if nodeCharacter == "|" {
					isInside = !isInside
					previousCharacter = ""
				}
				if previousCharacter == "F" && nodeCharacter == "J" {
					isInside = !isInside
					previousCharacter = ""
				}
				if previousCharacter == "L" && nodeCharacter == "J" {
					previousCharacter = ""
				}
				if previousCharacter == "L" && nodeCharacter == "7" {
					isInside = !isInside
					previousCharacter = ""
				}
				if previousCharacter == "F" && nodeCharacter == "7" {
					previousCharacter = ""
				}
			}

			if nodeCharacter == "." && isInside {
				visualization[x][y] = "I"
				tilesInsideCount++
			}
		}
	}

	// printVisualization(visualization)
	return tilesInsideCount
}

func Part2(path string) int {
	maze := loadInput(path)
	loopPoints := findLoop(maze)

	visualization := prepareMazeVisualization(maze, loopPoints)

	answer := findTilesInsideCount(visualization)

	return answer
}
