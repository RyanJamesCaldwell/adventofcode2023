package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

func getSeeds(lines []string) []int {
	seeds := []int{}
	seedsStr := strings.Split(lines[0], ": ")[1]
	for _, seedStr := range strings.Split(seedsStr, " ") {
		seedInt, ok := strconv.Atoi(seedStr)
		if ok != nil {
			panic("couldn't parse seed to int")
		}
		seeds = append(seeds, seedInt)
	}
	return seeds
}

type Mapper struct {
	Name     string
	FromType string
	ToType   string
	Ranges   []Range
}

func (m *Mapper) String() string {
	return fmt.Sprintf("Mapper: %s\nFromType: %s\nToType: %s\nRanges: %v\n", m.Name, m.FromType, m.ToType, m.Ranges)
}

type Range struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func getMappers(lines []string) []Mapper {
	mappers := []Mapper{}

	lineIdx := 2
	for lineIdx != len(lines) {
		line := lines[lineIdx]
		// if we hit a blank line, we've reached the end of a mapper definition
		if line == "" {
			continue
		}

		mapper := Mapper{}
		nameAndTypes := strings.Split(line, " ")[0]
		splitTypes := strings.Split(nameAndTypes, "-to-")

		mapper.Name = nameAndTypes
		mapper.FromType = splitTypes[0]
		mapper.ToType = splitTypes[1]

		lineIdx++
		line = lines[lineIdx]
		for line != "" {
			newRange := Range{}
			splitRange := strings.Split(line, " ")
			newRange.DestinationRangeStart, _ = strconv.Atoi(splitRange[0])
			newRange.SourceRangeStart, _ = strconv.Atoi(splitRange[1])
			newRange.RangeLength, _ = strconv.Atoi(splitRange[2])
			mapper.Ranges = append(mapper.Ranges, newRange)
			lineIdx++
			if lineIdx == len(lines) {
				break
			}
			line = lines[lineIdx]
			if line == "" {
				lineIdx++
			}
		}
		mappers = append(mappers, mapper)
	}
	return mappers
}

func main() {
	lines := fileReader.GetLines()
	seeds := getSeeds(lines)
	mappers := getMappers(lines)

	fmt.Println("Seeds: ", seeds)
	for _, m := range mappers {
		fmt.Println(m.String())
	}
}
