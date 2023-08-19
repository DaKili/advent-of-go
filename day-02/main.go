package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var inputFile = "input.txt"

var actionMap = map[string]string{
	"A": "Rock",
	"B": "Paper",
	"C": "Scissors",
	"X": "Rock",
	"Y": "Paper",
	"Z": "Scissors",
}
var pointMap = map[string]int{
	"Rock":     1,
	"Paper":    2,
	"Scissors": 3,
	"win":      6,
	"draw":     3,
	"loss":     0,
}
var resultMap = map[string]map[string]string{
	"Rock": {
		"Rock":     "draw",
		"Paper":    "win",
		"Scissors": "loss",
	},
	"Paper": {
		"Rock":     "loss",
		"Paper":    "draw",
		"Scissors": "win",
	},
	"Scissors": {
		"Rock":     "win",
		"Paper":    "loss",
		"Scissors": "draw",
	},
}
var requiredMap = map[string]map[string]string{
	"Rock": {
		"X": "Scissors",
		"Y": "Rock",
		"Z": "Paper",
	},
	"Paper": {
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	},
	"Scissors": {
		"X": "Paper",
		"Y": "Scissors",
		"Z": "Rock",
	},
}

type Game struct {
	opponent string
	self     string
	points   int
}

func newGame(oppententChoice string, selfChoice string) *Game {
	var game = new(Game)
	game.opponent = oppententChoice
	game.self = selfChoice
	game.points = gameResult(oppententChoice, selfChoice)
	return game
}

func gameResult(opponent string, self string) int {
	var score = pointMap[self] + pointMap[resultMap[opponent][self]]
	return int(score)
}

func main() {
	games1 := loadGuide1()
	points1 := getTotalPoints(games1)
	fmt.Printf("Total scode part 1: %v\n", points1)
	games2 := loadGuide2()
	points2 := getTotalPoints(games2)
	fmt.Printf("Total scode part 2: %v\n", points2)
}

func loadGuide1() []*Game {
	games := []*Game{}
	lines := getLinesFromFile()
	for _, v := range lines {
		if v == "" {
			continue
		}
		splitGames := strings.Split(v, " ")
		game := newGame(actionMap[splitGames[0]], actionMap[splitGames[1]])
		games = append(games, game)
	}
	return games
}

func loadGuide2() []*Game {
	games := []*Game{}
	lines := getLinesFromFile()
	for _, v := range lines {
		if v == "" {
			continue
		}
		splitGames := strings.Split(v, " ")
		required := requiredMap[actionMap[splitGames[0]]][splitGames[1]]
		game := newGame(actionMap[splitGames[0]], required)
		games = append(games, game)
	}
	return games
}

func getLinesFromFile() []string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func getTotalPoints(games []*Game) int {
	totalScore := 0
	for _, game := range games {
		totalScore += (*game).points
	}
	return totalScore
}
