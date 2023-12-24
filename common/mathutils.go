package common

import "math"

func Min(first int, second int) int {
	return int(math.Min(float64(first), float64(second)))
}

func Max(first int, second int) int {
	return int(math.Max(float64(first), float64(second)))
}

func Abs(value int) int {
	return int(math.Abs(float64(value)))
}

func greatestCommonDivisor(first int, second int) int {
	for second != 0 {
		first, second = second, first%second
	}
	return first
}

func lowestCommonMultiple(first int, second int) int {
	return first * second / greatestCommonDivisor(first, second)
}

func Lcm(numbers []int) (result int) {
	result = 1
	for _, value := range numbers {
		result = lowestCommonMultiple(result, value)
	}
	return result
}
