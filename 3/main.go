package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

var symbols = []string{"*", "#", "$", "+"}

type Number struct {
	StartIdx int
	EndIdx   int
	RowIdx   int
	Value    int
}

type Symbol struct {
	RowIdx int
	ColIdx int
	Value  string
}

func getNumbers(lines []string) []Number {
	numbers := []Number{}
	numRegex := regexp.MustCompile(`\d+`)

	for rowIdx, line := range lines {
		strNumbers := numRegex.FindAllString(line, -1)

		for _, val := range strNumbers {
			startIdx := strings.Index(line, val)
			intNumber, _ := strconv.Atoi(val)

			number := Number{StartIdx: startIdx, EndIdx: startIdx + len(val) - 1, RowIdx: rowIdx, Value: intNumber}
			numbers = append(numbers, number)
		}
	}

	return numbers
}

func findSymbols(lines []string) []Symbol {
	foundSymbols := []Symbol{}

	for rowIdx, line := range lines {
		for colIdx, char := range line {
			if slices.Contains(symbols, string(char)) {
				sym := Symbol{RowIdx: rowIdx, ColIdx: colIdx, Value: string(char)}
				foundSymbols = append(foundSymbols, sym)
			}
		}
	}

	return foundSymbols
}

func getPartNumbers(nums []Number, syms []Symbol) []Number {
	partNumbers := []Number{}

	return partNumbers
}

func (n *Number) String() string {
	return fmt.Sprintf("Number StartIdx %d, EndIdx %d, RowIdx %d, Value %d", n.StartIdx, n.EndIdx, n.RowIdx, n.Value)
}

func (s *Symbol) String() string {
	return fmt.Sprintf("Symbol RowIdx %d, ColIdx %d, Value %s", s.RowIdx, s.ColIdx, s.Value)
}

func main() {
	lines := fileReader.GetLines()
	numbers := getNumbers(lines)
	symbols := findSymbols(lines)

	// numbers adjacent to a symbol
	partNumbers := getPartNumbers(numbers, symbols)

	for _, num := range partNumbers {
		fmt.Println(num.String())
	}
}
