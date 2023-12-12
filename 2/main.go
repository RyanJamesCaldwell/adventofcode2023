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

	return game
}

func getGameIdFromLine(line string) int {
	regex := regexp.MustCompile(`\d+: `)
	regexResult := regex.FindString(line)
	strId := regexResult[0 : len(regexResult)-2]
	intId, _ := strconv.Atoi(strId)

	return intId
}

func (g *Game) PopulateFromLine(line string) {
}

func (g *Game) String() string {
	return fmt.Sprintf("Game %d: %d, %d, %d", g.ID, g.Blue, g.Green, g.Red)
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

	for _, game := range games {
		fmt.Println(game.String())
	}
}
