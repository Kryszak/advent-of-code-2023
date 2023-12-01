package common

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FileScanner(path string) *bufio.Scanner {
	inputFile, err := os.Open(path)
	check(err)

	fileScanner := bufio.NewScanner(inputFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}
