package day21

func mod(a, n int) int {
	return ((a % n) + n) % n
}

func stepInfinite(garden [][]rune, visited map[point]bool) map[point]bool {
	visitedPoints := make(map[point]bool)

	for p := range visited {
		for _, candidate := range [4]point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}} {
			if garden[mod(candidate.y, len(garden))][mod(candidate.x, len(garden[0]))] != '#' {
				visitedPoints[candidate] = true
			}
		}
	}

	return visitedPoints
}

func Part2(path string) (answer int) {
	garden := loadGarden(path)
	x, y := getStartingPoint(garden)

	visited := make(map[point]bool)
	visited[point{x, y}] = true

	var values [3]int

	steps := 26501365

	for i := 1; i <= (len(garden)/2)+2*len(garden); i++ {
		visited = stepInfinite(garden, visited)
		if i == (len(garden) / 2) {
			values[0] = len(visited)
		} else if i == (len(garden)/2)+len(garden) {
			values[1] = len(visited)
		} else if i == (len(garden)/2)+2*len(garden) {
			values[2] = len(visited)
		}
	}

	a := (values[2] + values[0] - 2*values[1]) / 2
	b := values[1] - values[0] - a
	c := values[0]
	n := steps / len(garden)

	answer = a*n*n + b*n + c

	return answer
}
