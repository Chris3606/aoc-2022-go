package main

import (
	"bufio"
	"os"
	"strconv"
)

type RopeMove struct {
	Direction Point
	Amount    int
}

func parseInput9(input string) ([]RopeMove, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var moves []RopeMove
	for scanner.Scan() {
		text := scanner.Text()

		var dir Point
		switch text[0] {
		case 'U':
			dir = UP
		case 'D':
			dir = DOWN
		case 'L':
			dir = LEFT
		case 'R':
			dir = RIGHT
		default:
			panic("Invalid movement direction")
		}

		amount, err := strconv.Atoi(text[2:])
		if err != nil {
			return nil, err
		}
		moves = append(moves, RopeMove{dir, amount})
	}

	return moves, nil

}

// Simulates the rope, and records number of positions visited.
func simulateRope(moves []RopeMove) int {
	head := Point{0, 0}
	tail := Point{0, 0}

	tailPositions := map[Point]bool{tail: true}
	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			// execute move
			head = head.Add(move.Direction)
			if ChebyshevDistance(head, tail) == 2 {
				delta := head.Sub(tail)
				// Normalize to a maximum of 1 movement
				if delta.X > 1 {
					delta.X = 1
				} else if delta.X < -1 {
					delta.X = -1
				}
				if delta.Y > 1 {
					delta.Y = 1
				} else if delta.Y < -1 {
					delta.Y = -1
				}

				// Move and record new position
				tail = tail.Add(Point{min(delta.X, 1), min(delta.Y, 1)})
				tailPositions[tail] = true
			}
		}
	}

	return len(tailPositions)
}

func Day09A(input string) int {
	moves, err := parseInput9(input)
	CheckError(err)

	return simulateRope(moves)
}

func Day09B(input string) int {
	panic("Not implemented")
}
