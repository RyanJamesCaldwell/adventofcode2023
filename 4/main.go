package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

type Card struct {
	ID             int
	Numbers        []int
	WinningNumbers []int
}

func (c *Card) String() string {
	return fmt.Sprintf("Card ID %d, Numbers: %v, Winning Numbers: %v", c.ID, c.Numbers, c.WinningNumbers)
}

func (c *Card) CardValue(part int) float64 {
	winCount := 0
	for _, num := range c.WinningNumbers {
		if slices.Contains(c.Numbers, num) {
			winCount++
		}
	}

	if part == 1 {
		if winCount <= 2 {
			return float64(winCount)
		} else {
			return math.Pow(2, float64(winCount-1))
		}
	} else {
		return float64(winCount)
	}
}

func getCards(lines []string) []Card {
	cards := []Card{}

	for _, line := range lines {
		newCard := Card{}

		// card id
		regex := regexp.MustCompile(`\s+\d+`)
		cardIDStr := strings.TrimSpace(regex.FindString(line))
		cardID, _ := strconv.Atoi(cardIDStr)
		newCard.ID = cardID

		// winning card numbers
		regex = regexp.MustCompile(`:\s+\d+.*\|`)
		winningCardNumbersStr := regex.FindString(line)
		winningCardNumbers := strings.Split(strings.NewReplacer(":", "", "|", "").Replace(winningCardNumbersStr), " ")
		for _, c := range winningCardNumbers {
			number, err := strconv.Atoi(c)
			if err != nil {
				continue
			}
			newCard.WinningNumbers = append(newCard.WinningNumbers, number)
		}

		// actual numbers for card
		regex = regexp.MustCompile(`\|\s+\d+.*`)
		numbersStr := regex.FindString(line)
		numbers := strings.Split(strings.NewReplacer("|", "").Replace(numbersStr), " ")
		for _, c := range numbers {
			number, err := strconv.Atoi(c)
			if err != nil {
				continue
			}
			newCard.Numbers = append(newCard.Numbers, number)
		}

		cards = append(cards, newCard)
	}

	return cards
}

func main() {
	lines := fileReader.GetLines()
	cards := getCards(lines)

	// part 1
	points := 0.0
	for _, card := range cards {
		points += card.CardValue(1)
	}
	fmt.Println("Part 1: ", points)

	// part 2
	for i := 0; i < len(cards); i++ {
		cardValue := int(cards[i].CardValue(2))
		if cardValue > 0 {
			for j := cards[i].ID; j < cards[i].ID+cardValue+1; j++ {
				cards = slices.Insert(cards, cards[j], i+1)
			}
		}
	}
	fmt.Println("Part 2: ", len(cards))
}
