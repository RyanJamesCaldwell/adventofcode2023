package main

import (
	"fmt"
	"regexp"
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

func (c *Card) CardValue() int {
	return 0
}

func getCards(lines []string) []Card {
	cards := []Card{}

	for _, line := range lines {
		newCard := Card{}
		regex := regexp.MustCompile(`\s+\d+`)
		cardNumberStr := strings.TrimSpace(regex.FindString(line))
		cardNumber, _ := strconv.Atoi(cardNumberStr)
		newCard.ID = cardNumber

		cards = append(cards, newCard)
	}

	return cards
}

func main() {
	lines := fileReader.GetLines()
	cards := getCards(lines)
	for _, card := range cards {
		fmt.Println(card.String())
	}
}
