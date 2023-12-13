package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

var symbols = []string{"*", "#", "$", "+", "%", "-", "=", "@", "&", "/"}

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
			line = strings.Replace(line, val, strings.Repeat(".", len(val)), 1)
		}
	}

	return numbers
}

func findSymbols(lines []string) []Symbol {
	foundSymbols := []Symbol{}

	for rowIdx, line := range lines {
		for colIdx, char := range line {
			if isNaN(string(char)) && string(char) != "." {
				sym := Symbol{RowIdx: rowIdx, ColIdx: colIdx, Value: string(char)}
				foundSymbols = append(foundSymbols, sym)
			}
		}
	}

	return foundSymbols
}

func isNaN(char string) bool {
	_, err := strconv.Atoi(char)
	return err != nil
}

func getPartNumbers(nums []Number, syms []Symbol) []Number {
	partNumbers := []Number{}

	for _, num := range nums {
		if num.DoesBorderSymbol(syms) {
			partNumbers = append(partNumbers, num)
		}
	}

	return partNumbers
}

func (n *Number) DoesBorderSymbol(syms []Symbol) bool {
	for _, sym := range syms {
		for i := n.StartIdx; i <= n.EndIdx; i++ {
			if abs(n.RowIdx-sym.RowIdx) <= 1 && abs(i-sym.ColIdx) <= 1 {
				return true
			}
		}
		// Check if the symbol is directly above or below a singular number
		if n.StartIdx == n.EndIdx && abs(sym.RowIdx-n.RowIdx) == 1 && sym.ColIdx == n.StartIdx {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	sum := 0
	for _, num := range partNumbers {
		sum += num.Value
	}
	fmt.Println("Part 1: ", sum)
}
