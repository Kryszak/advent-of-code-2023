package day08

import (
	"strings"

	"github.com/Kryszak/aoc2023/common"
)

var multiNodeStepIndex = 0

func getStartingNodes(nodeMap map[string]node) []string {
	var nodes []string

	for key := range nodeMap {
		if strings.HasSuffix(key, "A") {
			nodes = append(nodes, key)
		}
	}

	return nodes
}

func areAllNodesFinal(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

func getRequiredStepsForNode(nodeMap map[string]node, stepSequence string, startNode string) (requiredSteps int) {
	sequenceCounter := 0
	currentNode := startNode
	for !strings.HasSuffix(currentNode, "Z") {
		if sequenceCounter == len(stepSequence) {
			sequenceCounter = 0
		}
		direction := string(stepSequence[sequenceCounter])
		currentNode = getNextNode(nodeMap, currentNode, direction)

		requiredSteps++
		sequenceCounter++
	}
	return requiredSteps
}

func Part2(path string) (answer int) {
	answer = 1

	fileScanner := common.FileScanner(path)

	fileScanner.Scan()
	stepSequence := fileScanner.Text()

	nodeMap := loadNodes(fileScanner)

	currentNodes := getStartingNodes(nodeMap)
	var requiredStepsPerNode []int

	for _, node := range currentNodes {
		requiredStepsPerNode = append(requiredStepsPerNode, getRequiredStepsForNode(nodeMap, stepSequence, node))
	}

	answer = common.Lcm(requiredStepsPerNode)
	return answer
}
