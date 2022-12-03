package main

import (
	"advent-of-code/models"
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type RPSMatch struct{}
type RPSMatches []*RPSMatch
type RPSGame struct {
	Player1 models.Elf
	Player2 models.Elf
	Matches RPSMatches
}

func main() {
	// filename := "sample.txt"
	filename := "input.txt"

	txtlines := utils.ReadFile(filename)

	points := versus(txtlines)

	fmt.Printf("Final points are %d\n", points)
}

func versus(line []string) int {
	var score = 0
	for i := 0; i < len(line); i++ {
		score += playPart2(line[i])
	}
	return score
}

const (
	loss     = 0
	tie      = 3
	win      = 6
	rock     = 1
	paper    = 2
	scissors = 3
)

var (
	actionsMap = map[string]string{
		"A": "rock",
		"B": "paper",
		"C": "scissors",
		"X": "rock",
		"Y": "paper",
		"Z": "scissors",
	}

	outcomeMap = map[string]string{
		"X": "loses",
		"Y": "ties",
		"Z": "wins",
	}
)

func playPart1(line string) int {
	s := strings.Fields(line)
	result := playRound(actionsMap[s[0]], actionsMap[s[1]])

	if result == 0 {
		fmt.Printf("No match detected for input line %s\n", line)
		panic("no result part 1")
	}

	return result
}

func playPart2(line string) int {
	s := strings.Fields(line)
	myAction := determineAction(actionsMap[s[0]], outcomeMap[s[1]])
	if myAction == "" {
		fmt.Printf("No match detected for input line %s\n", line)
		panic("no result part 2")
	}

	result := playRound(actionsMap[s[0]], myAction)

	if result == 0 {
		fmt.Printf("No match detected for input line %s\n", line)
		panic("no result")
	}

	return result
}

func playRound(opponent string, me string) int {
	round := fmt.Sprintf("%s vs %s", opponent, me)
	switch round {
	case "rock vs rock": // tie
		return tie + rock
	case "rock vs paper": // win
		return win + paper
	case "rock vs scissors": // loss
		return loss + scissors
	case "paper vs rock": // loss
		return loss + rock
	case "paper vs paper": // tie
		return tie + paper
	case "paper vs scissors": // win
		return win + scissors
	case "scissors vs rock": // win
		return win + rock
	case "scissors vs paper": // loss
		return loss + paper
	case "scissors vs scissors": // tie
		return tie + scissors
	default:
		fmt.Printf("No match detected for round of %s\n", round)
	}

	return 0
}

func determineAction(opponent string, outcome string) string {
	round := fmt.Sprintf("%s %s", opponent, outcome)
	switch round {
	case "rock ties": // rock ties against rock
		return "rock"
	case "rock wins": // paper wins against rock
		return "paper"
	case "rock loses": // scissors loses against rock
		return "scissors"
	case "paper loses": // rock loses against paper
		return "rock"
	case "paper ties": // paper ties against paper
		return "paper"
	case "paper wins": // scissors wins against paper
		return "scissors"
	case "scissors wins": // rock wins against scissors
		return "rock"
	case "scissors loses": // paper loses against scissors
		return "paper"
	case "scissors ties": // scissors ties against scissors
		return "scissors"
	default:
		fmt.Printf("No match detected for round of %s\n", round)
	}

	return ""
}
