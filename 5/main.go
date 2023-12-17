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

func getPart2Seeds(seeds []int) []int {
	newSeeds := []int{}

	for i := 0; i < len(seeds); i += 2 {
		rangeStart := seeds[i]
		rangeLength := seeds[i+1]

		for j := rangeStart; j < rangeStart+rangeLength; j++ {
			newSeeds = append(newSeeds, j)
		}
	}

	return newSeeds
}

type Mapper struct {
	Name     string
	FromType string
	ToType   string
	Ranges   []Range
}

func (m *Mapper) GetDestinationNumber(sourceNumber int) int {
	for _, r := range m.Ranges {
		if sourceNumber >= r.SourceRangeStart && sourceNumber <= r.SourceRangeStart+(r.RangeLength-1) { // e.g. 90, 2 is 90 and 91, not 90, 91, 92
			return r.DestinationRangeStart + (sourceNumber - r.SourceRangeStart)
		}
	}
	return sourceNumber
}

func (m *Mapper) String() string {
	return fmt.Sprintf("Mapper: %s\nFromType: %s\nToType: %s\nRanges: %v\n", m.Name, m.FromType, m.ToType, m.Ranges)
}

type Range struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

func getOrderedMappers(lines []string) map[int]*Mapper {
	orderedMappers := map[int]*Mapper{}

	lineIdx := 2
	mapperCount := 0
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
		orderedMappers[mapperCount] = &mapper
		mapperCount++
	}
	return orderedMappers
}

func traverseMappersForSeedLocations(seeds []int, mappers map[int]*Mapper) []int {
	locationNumbers := []int{}

	for _, seed := range seeds {
		currentValue := seed
		// iterate through each mapper in order and update currentValue
		for i := 0; i < len(mappers); i++ {
			currentValue = mappers[i].GetDestinationNumber(currentValue)
		}
		locationNumbers = append(locationNumbers, currentValue)
	}

	return locationNumbers
}

func getMinValue(numbers []int) int {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func main() {
	lines := fileReader.GetLines()
	seeds := getSeeds(lines)
	mappers := getOrderedMappers(lines)

	// Part 1
	results := traverseMappersForSeedLocations(seeds, mappers)
	fmt.Println("Part 1: ", getMinValue(results))

	// Part 2
	// D'oh! Seeds are actually defined in pairs, e.g. start_value + (range_length - 1) => 90 5 = 90, 91, 92, 93, 94
	// Brute force because absolutely fucking not. Got the result in ~7 minutes, which is faster than I would've spent
	// coming up with some highly optimized reverse-search / range splitting solution.
	p2Seeds := getPart2Seeds(seeds)
	results = traverseMappersForSeedLocations(p2Seeds, mappers)
	fmt.Println("Part 2: ", getMinValue(results))
}
