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
func simulateRope(moves []RopeMove, numKnots int) int {
	knots := make([]Point, numKnots)

	tailPositions := map[Point]bool{knots[len(knots)-1]: true}
	for _, move := range moves {
		for i := 0; i < move.Amount; i++ {
			// Execute move
			knots[0] = knots[0].Add(move.Direction)
			// Follow the leader
			for knot_idx := 1; knot_idx < len(knots); knot_idx++ {
				if ChebyshevDistance(knots[knot_idx-1], knots[knot_idx]) == 2 {
					delta := knots[knot_idx-1].Sub(knots[knot_idx])
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

					// Execute the appropriate follow action
					knots[knot_idx] = knots[knot_idx].Add(Point{delta.X, delta.Y})
				}
			}

			// Record new tail position as visited
			tailPositions[knots[len(knots)-1]] = true
		}
	}

	return len(tailPositions)
}

func Day09A(input string) int {
	moves, err := parseInput9(input)
	CheckError(err)

	return simulateRope(moves, 2)
}

func Day09B(input string) int {
	moves, err := parseInput9(input)
	CheckError(err)

	return simulateRope(moves, 10)
}
