package common

import "strconv"

func Atoi(value string) int {
	parsed, err := strconv.Atoi(value)
	check(err)
	return parsed
}
