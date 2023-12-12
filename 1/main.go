package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()
	results := make([]int, len(lines))

	for idx, line := range lines {
		line := replaceSpelledOutNumbers(line)

		res := getTwoDigitNumber(line)
		fmt.Println("two digit num: ", res)
		fmt.Println("===========\n")
		results[idx] = res
	}

	fmt.Println("Result: ", getSum(results))
}

func replaceSpelledOutNumbers(line string) string {
	r := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")
	fmt.Println("line: ", line)
	fmt.Println("replaced: ", r.Replace(line))

	return r.Replace(r.Replace(line))
}

func getTwoDigitNumber(line string) int {
	reg := regexp.MustCompile(`[0-9]`)
	ints := reg.FindAllString(line, -1)

	first, err := strconv.Atoi(ints[0])
	if len(ints) == 1 {
		return first*10 + first
	}
	second, err2 := strconv.Atoi(ints[len(ints)-1])
	if err != nil || err2 != nil {
		panic("Error converting string to int")
	}

	return first*10 + second
}

func getSum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func getLines() []string {
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
