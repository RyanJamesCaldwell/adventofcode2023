package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ryanjamescaldwell/adventofcode2023/fileReader"
)

type Game struct {
	ID    int
	Blue  int
	Green int
	Red   int
}

func (g *Game) Total() int {
	return g.Blue + g.Green + g.Red
}

func lineToGame(line string) Game {
	var game Game
	game.ID = getGameIdFromLine(line)
	game.Red = getMaxColorCountFromLine(line, "red")
	game.Blue = getMaxColorCountFromLine(line, "blue")
	game.Green = getMaxColorCountFromLine(line, "green")

	return game
}

func getMaxColorCountFromLine(line string, color string) int {
	colorRegex := regexp.MustCompile(`\d+ ` + color)
	stringMatches := colorRegex.FindAllString(line, -1) // e.g. [4 red, 3 red]

	max := -1

	for _, val := range stringMatches {
		intRegex := regexp.MustCompile(`\d+`)
		countStr := intRegex.FindString(val)
		countInt, _ := strconv.Atoi(countStr)
		if countInt > max {
			max = countInt
		}
	}

	return max
}

func possibleGames(games []Game, theoreticalGame Game) []Game {
	possibles := []Game{}

	for _, game := range games {
		if game.Blue <= theoreticalGame.Blue && game.Red <= theoreticalGame.Red && game.Green <= theoreticalGame.Green {
			possibles = append(possibles, game)
		}
	}

	return possibles
}

func getGameIdFromLine(line string) int {
	regex := regexp.MustCompile(`\d+: `)
	regexResult := regex.FindString(line)
	strId := regexResult[0 : len(regexResult)-2]
	intId, _ := strconv.Atoi(strId)

	return intId
}

func (g *Game) String() string {
	return fmt.Sprintf("Game %d: %d blue, %d green, %d red", g.ID, g.Blue, g.Green, g.Red)
}

func buildGames(lines []string) []Game {
	games := make([]Game, len(lines))

	for idx, line := range lines {
		games[idx] = lineToGame(line)
	}

	return games
}

func main() {
	lines := fileReader.GetLines()
	games := buildGames(lines)
	theoreticalGame := Game{Red: 12, Green: 13, Blue: 14}
	possibleGames := possibleGames(games, theoreticalGame)

	idSum := 0
	for _, game := range possibleGames {
		idSum += game.ID
	}
	fmt.Println("Sum: ", idSum)
}
