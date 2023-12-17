package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

type Hand struct {
	Cards     []string
	Bid       int
	TypeCache Type
}

type Type struct {
	Name  string
	Value int // relative value of hand type 7 = five of a kind, 1 = high card
}

func (h *Hand) Type() Type {
	if h.TypeCache.Name != "" && h.TypeCache.Value != 0 {
		return h.TypeCache
	}

	uniqueTypeCounts := map[string]int{}

	for _, card := range h.Cards {
		uniqueTypeCounts[card]++
	}

	uniqueTypeCountsLength := len(uniqueTypeCounts)
	if uniqueTypeCountsLength == 1 {
		h.TypeCache = Type{Name: "five of a kind", Value: 7}
	} else if uniqueTypeCountsLength == 2 && uniqueTypeCounts[h.Cards[0]] == 4 || uniqueTypeCounts[h.Cards[1]] == 4 {
		h.TypeCache = Type{Name: "four of a kind", Value: 6}
	} else if uniqueTypeCountsLength == 3 && uniqueTypeCounts[h.Cards[0]] == 3 || uniqueTypeCounts[h.Cards[1]] == 3 || uniqueTypeCounts[h.Cards[2]] == 3 {
		h.TypeCache = Type{Name: "three of a kind", Value: 4}
	} else if uniqueTypeCountsLength == 2 && uniqueTypeCounts[h.Cards[0]] == 3 || uniqueTypeCounts[h.Cards[1]] == 3 {
		h.TypeCache = Type{Name: "full house", Value: 5}
	} else if uniqueTypeCountsLength == 3 && uniqueTypeCounts[h.Cards[0]] == 2 || uniqueTypeCounts[h.Cards[1]] == 2 || uniqueTypeCounts[h.Cards[2]] == 2 {
		h.TypeCache = Type{Name: "two pair", Value: 3}
	} else if uniqueTypeCountsLength == 4 {
		h.TypeCache = Type{Name: "one pair", Value: 2}
	} else if uniqueTypeCountsLength == 5 {
		h.TypeCache = Type{Name: "high card", Value: 1}
	} else {
		panic("wasn't able to parse hand correctly: " + h.String())
	}

	return h.TypeCache
}

func (h *Hand) String() string {
	return fmt.Sprintf("Hand %v %v", h.Cards, h.Bid)
}

func (h *Hand) Beats(otherHand Hand) bool {
	if h.Type().Value == otherHand.Type().Value {
		// handle more complicated cases here
		return false
	}

	return h.Type().Value > otherHand.Type().Value
}

func getHands(lines []string) []Hand {
	hands := []Hand{}
	for _, line := range lines {
		hand := Hand{}
		splitHand := strings.Split(line, " ")
		hand.Cards = strings.Split(splitHand[0], "")
		bidInt, err := strconv.Atoi(splitHand[1])
		if err != nil {
			panic("could not convert bidStr to int")
		}
		hand.Bid = bidInt

		hands = append(hands, hand)
	}
	return hands
}

func main() {
	lines := fileReader.GetLines()
	hands := getHands(lines)

	for _, hand := range hands {
		fmt.Println(hand.String())
		fmt.Println(hand.Type())
		fmt.Println()
	}
}
