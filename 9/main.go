package main

import (
	"fmt"
	"os"

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

	if os.Args[1] == "sample.txt" {
		fmt.Println("Sample answer: 114")
	}

	// TODO
	// For each history (line), we need to figure out the next number and add it to an array
	// We shouldn't need to wait for 0, just need all numbers in the line to be the same
	// Once each history's next number is identified, sum them.
	fmt.Println("Part 1: ", 0)
}
