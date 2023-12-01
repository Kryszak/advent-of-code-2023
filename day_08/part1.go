package day08

import (
	"bufio"
	"regexp"
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

var currentStepIndex = 0

type node struct {
	left  string
	right string
}

func loadNodes(fileScanner *bufio.Scanner) map[string]node {
	nodeMap := make(map[string]node)

	fileScanner.Scan()

	for fileScanner.Scan() {
		line := strings.Split(fileScanner.Text(), " = ")
		nodeKey := line[0]
		nodeValues := strings.Fields(regexp.MustCompile("[(),]").ReplaceAllString(line[1], ""))

		nodeMap[nodeKey] = node{nodeValues[0], nodeValues[1]}
	}

	return nodeMap
}

func getNextNode(nodeMap map[string]node, currentNode string, direction string) string {
	nextNode, _ := nodeMap[currentNode]
	if direction == "L" {
		return nextNode.left
	} else {
		return nextNode.right
	}
}

func Part1() int {
	answer := 0

	fileScanner := common.FileScanner("day_08/input.txt")

	fileScanner.Scan()
	stepSequence := fileScanner.Text()

	nodeMap := loadNodes(fileScanner)

	startNode := "AAA"
	endNode := "ZZZ"

	currentNode := startNode

	for currentNode != endNode {
		if currentStepIndex == len(stepSequence) {
			currentStepIndex = 0
		}
		direction := string(stepSequence[currentStepIndex])
		currentNode = getNextNode(nodeMap, currentNode, direction)

		answer++
		currentStepIndex++
	}
	return answer
}
