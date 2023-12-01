package day11

func Part2() int {
	answer := 0

	universe, galaxies := loadUniverse()
	expandedRows, expandedColumns := getExpandedColumnsAndRows(&universe)

	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pathLength := calculateGalaxiesDistance(galaxies[i], galaxies[j], expandedRows, expandedColumns, 1000000)
			answer += pathLength
		}
	}

	return answer
}
