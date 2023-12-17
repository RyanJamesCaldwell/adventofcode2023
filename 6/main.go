package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

type Race struct {
	Milliseconds int
	Distance     int
}

func (r *Race) String() string {
	return fmt.Sprintf("Race: %dms, %dmm", r.Milliseconds, r.Distance)
}

// meh, this is short input, just read it in really inefficiently :)
func getRaces(lines []string) []Race {
	races := []Race{}
	timesInMsStr := strings.Split(strings.Split(lines[0], ":")[1], " ")
	distancesStr := strings.Split(strings.Split(lines[1], ":")[1], " ")

	times := []int{}
	for i := 0; i < len(timesInMsStr); i++ {
		if timesInMsStr[i] == " " || timesInMsStr[i] == "" {
			continue
		}

		ms, err := strconv.Atoi(timesInMsStr[i])
		if err != nil {
			panic("Could not convert to Race")
		}

		times = append(times, ms)
	}

	distances := []int{}
	for i := 0; i < len(distancesStr); i++ {
		if distancesStr[i] == " " || distancesStr[i] == "" {
			continue
		}
		distance, err := strconv.Atoi(distancesStr[i])
		if err != nil {
			panic("Could not convert to Race")
		}
		distances = append(distances, distance)
	}

	for i := 0; i < len(times); i++ {
		newRace := Race{}
		newRace.Milliseconds = times[i]
		newRace.Distance = distances[i]
		races = append(races, newRace)
	}

	return races
}

// There's actually only one race, silly us
func getPart2Race(lines []string) Race {
	newRace := Race{}

	timesInMsStr := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	distancesStr := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")

	ms, err := strconv.Atoi(timesInMsStr)
	if err != nil {
		panic("Couldn't convert str to int")
	}
	newRace.Milliseconds = ms

	distance, err := strconv.Atoi(distancesStr)
	if err != nil {
		panic("Couldn't convert str to int")
	}
	newRace.Distance = distance

	return newRace
}

func product(nums []int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}

func getWinnableRaceStrategies(races []Race) map[Race]int {
	winnableRaceStrats := make(map[Race]int, len(races))

	for _, race := range races {
		count := 0

		for timeButtonPressed := 1; timeButtonPressed < race.Milliseconds; timeButtonPressed++ {
			newDistance := timeButtonPressed * (race.Milliseconds - timeButtonPressed)
			if newDistance > race.Distance {
				count++
			}
		}

		winnableRaceStrats[race] = count
	}

	return winnableRaceStrats
}

func main() {
	lines := fileReader.GetLines()
	races := getRaces(lines)

	// Part 1
	strats := getWinnableRaceStrategies(races)
	winningCountsPerRace := []int{}
	for _, count := range strats {
		winningCountsPerRace = append(winningCountsPerRace, count)
	}

	fmt.Println("Part 1: ", product(winningCountsPerRace))

	// Part 2
	race := getPart2Race(lines)
	races = []Race{race}
	strats = getWinnableRaceStrategies(races)
	winningCountsPerRace = []int{}
	for _, count := range strats {
		winningCountsPerRace = append(winningCountsPerRace, count)
	}

	fmt.Println("Part 2: ", product(winningCountsPerRace))
}
