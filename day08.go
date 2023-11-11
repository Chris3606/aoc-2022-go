package main

import (
	"os"
)

func parseInput8(input string) (Grid[byte], error) {
	f, err := os.Open(input)
	if err != nil {
		return Grid[byte]{}, err
	}
	defer f.Close()

	return ReadDigitGrid(f)
}

func setVisibility(x, y int, grid *Grid[byte], visible *map[Point]bool, maxHeight *int) bool {
	point := Point{x, y}
	val := grid.GetCopy(point)
	if int(val) > *maxHeight {
		(*visible)[point] = true
		*maxHeight = int(val)
	}

	// This is the max height of the digit grid, so we know we're done if we see something
	// at height 9
	return val == 9
}

func calculateScenicScore(point Point, grid *Grid[byte]) int {
	score := 1
	val := grid.GetCopy(point)

	directionScore := 0
	for y := point.Y + 1; y < grid.Height(); y++ {
		directionScore++
		if grid.GetCopy(Point{point.X, y}) >= val {
			break
		}
	}
	score *= directionScore

	directionScore = 0
	for y := point.Y - 1; y >= 0; y-- {
		directionScore++
		if grid.GetCopy(Point{point.X, y}) >= val {
			break
		}
	}
	score *= directionScore

	directionScore = 0
	for x := point.X + 1; x < grid.Width(); x++ {
		directionScore++
		if grid.GetCopy(Point{x, point.Y}) >= val {
			break
		}
	}
	score *= directionScore

	directionScore = 0
	for x := point.X - 1; x >= 0; x-- {
		directionScore++
		if grid.GetCopy(Point{x, point.Y}) >= val {
			break
		}
	}
	score *= directionScore

	return score
}

func Day08A(input string) int {
	grid, err := parseInput8(input)
	CheckError(err)

	// We'll record visible things in a map as we go
	visible := map[Point]bool{}

	// Check all columns going down and up to set visibility
	for x := 0; x < grid.Width(); x++ {
		maxHeight := -1
		for y := 0; y < grid.Height(); y++ {
			setVisibility(x, y, &grid, &visible, &maxHeight)
		}

		maxHeight = -1
		for y := grid.Height() - 1; y > 0; y-- {
			setVisibility(x, y, &grid, &visible, &maxHeight)
		}
	}

	// Same but for rows
	for y := 0; y < grid.Height(); y++ {
		maxHeight := -1
		for x := 0; x < grid.Width(); x++ {
			setVisibility(x, y, &grid, &visible, &maxHeight)
		}

		maxHeight = -1
		for x := grid.Width() - 1; x > 0; x-- {
			setVisibility(x, y, &grid, &visible, &maxHeight)
		}
	}

	return len(visible)
}

func Day08B(input string) int {
	grid, err := parseInput8(input)
	CheckError(err)

	positions := grid.Positions()

	score := -1
	for positions.Next() {
		pos := positions.Current()
		curScore := calculateScenicScore(pos, &grid)
		score = max(score, curScore)
		if curScore > score {
			score = curScore
		}
	}

	return score
}
