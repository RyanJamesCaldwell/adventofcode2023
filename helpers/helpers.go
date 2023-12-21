package helpers

import (
	"strconv"
	"strings"
)

func GetIntsFromLine(line string, delimeter string) []int {
	var ints []int

	splitLine := strings.Split(line, delimeter)
	for i := 0; i < len(splitLine); i++ {
		intVal, _ := strconv.Atoi(string(splitLine[i]))
		ints = append(ints, intVal)
	}
	return ints
}
