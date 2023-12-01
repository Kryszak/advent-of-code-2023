package day02

import (
	"strconv"
	"strings"
)

func parseIdAndRoundValues(line string) (id int, roundValues []string) {
	idValuesSplit := strings.Split(line, ":")
	idString := strings.Replace(idValuesSplit[0], "Game ", "", -1)

	id, _ = strconv.Atoi(idString)
	roundValues = strings.Split(idValuesSplit[1], ";")
	return id, roundValues
}

func parseColorCountAndName(entry string) (count int, colorName string) {
	entry = strings.Trim(entry, " ")
	colorCountSplit := strings.Split(entry, " ")
	count, _ = strconv.Atoi(strings.Trim(colorCountSplit[0], " "))
	colorName = strings.ToLower(colorCountSplit[1])
	return count, colorName
}
