package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

type Hand struct {
	Cards     []string
	Bid       int
	TypeCache Type
	PairCount int
}

type Type struct {
	Name  string
	Value int // relative value of hand type 7 = five of a kind, 1 = high card
}

var cardLabelValues = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
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
	h.PairCount = 0
	for _, count := range uniqueTypeCounts {
		if count == 2 {
			h.PairCount++
		}
	}

	uniqueTypeCountsValues := []int{}
	for _, value := range uniqueTypeCounts {
		uniqueTypeCountsValues = append(uniqueTypeCountsValues, value)
	}

	if uniqueTypeCountsLength == 1 {
		h.TypeCache = Type{Name: "five of a kind", Value: 7}
	} else if uniqueTypeCountsLength == 2 && slices.Contains(uniqueTypeCountsValues, 4) {
		h.TypeCache = Type{Name: "four of a kind", Value: 6}
	} else if uniqueTypeCountsLength == 3 && slices.Contains(uniqueTypeCountsValues, 3) {
		h.TypeCache = Type{Name: "three of a kind", Value: 4}
	} else if uniqueTypeCountsLength == 2 && slices.Contains(uniqueTypeCountsValues, 3) && slices.Contains(uniqueTypeCountsValues, 2) {
		h.TypeCache = Type{Name: "full house", Value: 5}
	} else if uniqueTypeCountsLength == 3 && h.PairCount == 2 {
		h.TypeCache = Type{Name: "two pair", Value: 3}
	} else if uniqueTypeCountsLength == 4 && h.PairCount == 1 {
		h.TypeCache = Type{Name: "one pair", Value: 2}
	} else if uniqueTypeCountsLength == 5 {
		h.TypeCache = Type{Name: "high card", Value: 1}
	} else {
		panic("wasn't able to parse hand correctly: " + h.String())
	}

	fmt.Println("Cards ", h.Cards, "are of type", h.TypeCache)

	return h.TypeCache
}

func (h *Hand) String() string {
	return fmt.Sprintf("Hand %v %v", h.Cards, h.Bid)
}

func (h *Hand) Beats(otherHand Hand) bool {
	if h.Type().Value == otherHand.Type().Value {
		for i := 0; i < len(h.Cards); i++ {
			leftLabel := h.Cards[i]
			rightLabel := otherHand.Cards[i]
			if cardLabelValues[leftLabel] == cardLabelValues[rightLabel] {
				continue
			} else {
				return cardLabelValues[leftLabel] > cardLabelValues[rightLabel]
			}
		}
		return true
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

func totalWinnings(sortedHands []Hand) int {
	total := 0

	for i := 0; i < len(sortedHands); i++ {
		rank := len(sortedHands) - i
		handWinAmount := rank * sortedHands[i].Bid
		total += handWinAmount
	}
	return total
}

func main() {
	lines := fileReader.GetLines()
	hands := getHands(lines)

	// Sort hands in descending order
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Beats(hands[j])
	})

	fmt.Println("Part 1: ", totalWinnings(hands))
}
