package main

import (
	"bufio"
	"os"
)

type Move int

const (
	Rock     Move = 1
	Paper    Move = 2
	Scissors Move = 3
)

type Result int

const (
	Loss Result = 0
	Draw Result = 3
	Win  Result = 6
)

type Round struct {
	PlayerMove Move
	Outcome    Result
}

func (move Move) beats() Move {
	switch move {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	case Scissors:
		return Paper
	default:
		panic("Unsupported move.")
	}
}

func (move Move) beatBy() Move {
	switch move {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	case Scissors:
		return Rock
	default:
		panic("Unsupported move.")
	}
}

func getOutcome(playerMove, opponentMove Move) Result {
	if playerMove == opponentMove {
		return Draw
	} else if opponentMove == playerMove.beats() {
		return Win
	} else {
		return Loss
	}
}

func scoreRounds(rounds []Round) int {
	var totalScore int
	for _, round := range rounds {
		totalScore += (int)(round.PlayerMove) + (int)(round.Outcome)
	}

	return totalScore
}

func parseInputA(input string) ([]Round, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var rounds []Round
	for scanner.Scan() {
		var opponentMove Move
		switch scanner.Text()[0] {
		case 'A':
			opponentMove = Rock
		case 'B':
			opponentMove = Paper
		case 'C':
			opponentMove = Scissors
		}

		var playerMove Move
		switch scanner.Text()[2] {
		case 'X':
			playerMove = Rock
		case 'Y':
			playerMove = Paper
		case 'Z':
			playerMove = Scissors
		}

		rounds = append(rounds, Round{playerMove, getOutcome(playerMove, opponentMove)})
	}

	return rounds, nil
}

func parseInputB(input string) ([]Round, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var rounds []Round
	for scanner.Scan() {
		var opponentMove Move
		switch scanner.Text()[0] {
		case 'A':
			opponentMove = Rock
		case 'B':
			opponentMove = Paper
		case 'C':
			opponentMove = Scissors
		}

		var outcome Result
		switch scanner.Text()[2] {
		case 'X':
			outcome = Loss
		case 'Y':
			outcome = Draw
		case 'Z':
			outcome = Win
		}
		var playerMove Move
		switch outcome {
		case Loss:
			playerMove = opponentMove.beats()
		case Draw:
			playerMove = opponentMove
		case Win:
			playerMove = opponentMove.beatBy()
		}

		rounds = append(rounds, Round{playerMove, outcome})
	}

	return rounds, nil
}

func Day02A(input string) int {
	rounds, err := parseInputA(input)
	CheckError(err)

	return scoreRounds(rounds)
}

func Day02B(input string) int {
	rounds, err := parseInputB(input)
	CheckError(err)

	return scoreRounds(rounds)
}
