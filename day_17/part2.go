package day17

func Part2(path string) (answer int) {
	input := loadInput(path)
	answer = bfs(input, input[len(input)-1][len(input[0])-1], 4, 10)
	return answer
}
