package common

import "math"

func Min(first int, second int) int {
	return int(math.Min(float64(first), float64(second)))
}

func Max(first int, second int) int {
	return int(math.Max(float64(first), float64(second)))
}
