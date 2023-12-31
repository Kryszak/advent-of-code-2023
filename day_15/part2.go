package day15

import (
	"slices"

	"github.com/Kryszak/aoc2023/common"
)

type step struct {
	label       []rune
	operation   rune
	focalLength int
}

type lens struct {
	label string
	power int
}

func (a lens) hasLabel(label string) bool {
	return a.label == label
}

func parseDigit(character rune) int {
	return common.Atoi(string(character))
}

func parseSteps(input [][]rune) []step {
	var steps []step

	for _, raw := range input {
		var label []rune
		var operation rune
		var focalLength int
		if slices.Contains(raw, '=') {
			i := 0
			for raw[i] != '=' {
				label = append(label, raw[i])
				i++
			}
			operation = '='
			focalLength = parseDigit(raw[len(raw)-1])
		}
		if slices.Contains(raw, '-') {
			label = raw[0 : len(raw)-1]
			operation = '-'
		}
		steps = append(steps, step{label, operation, focalLength})
	}

	return steps
}

func installLenses(steps []step) map[int][]lens {
	boxes := make(map[int][]lens, 256)

	for _, step := range steps {
		boxNumber := calculateStepHash(step.label)
		lensLabel := string(step.label)

		if step.operation == '=' {
			lensIndex := slices.IndexFunc(boxes[boxNumber], func(l lens) bool {
				return l.hasLabel(lensLabel)
			})
			if lensIndex != -1 {
				boxes[boxNumber][lensIndex].power = step.focalLength
			} else {
				boxes[boxNumber] = append(boxes[boxNumber], lens{lensLabel, step.focalLength})
			}
		} else {
			boxes[boxNumber] = slices.DeleteFunc(boxes[boxNumber], func(l lens) bool {
				return l.hasLabel(lensLabel)
			})
		}
	}

	for i := range boxes {
		if len(boxes[i]) == 0 {
			delete(boxes, i)
		}
	}

	return boxes
}

func calculateBoxFocusingPower(boxIndex int, lenses []lens) int {
	result := 0

	for lensIndex, lens := range lenses {
		result += (boxIndex + 1) * (lensIndex + 1) * lens.power
	}

	return result
}

func Part2(path string) (answer int) {
	input := loadInput(path)
	steps := parseSteps(input)
	boxes := installLenses(steps)

	for i, box := range boxes {
		answer += calculateBoxFocusingPower(i, box)
	}

	return answer
}
