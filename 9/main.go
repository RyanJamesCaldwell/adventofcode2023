package main

import (
	"fmt"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
	"github.com/ryanjamescaldwell/adventofcode2023/helpers"
)

func getInputFromLine(lines []string) [][]int {
	input := make([][]int, len(lines))
	for i := 0; i < len(lines); i++ {
		input[i] = helpers.GetIntsFromLine(lines[i], " ")
	}

	return input
}

func main() {
	input := getInputFromLine(fileReader.GetLines())

	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
	}

	fmt.Println("Part 1: ", 0)
}
