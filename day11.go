package main

import (
	"bufio"
	"os"
)

func parseInput11(input string) ([]int, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		// Parse lines
	}

	panic("Not implemented")
}

func Day11A(input string) int {
	_, err := parseInput11(input)
	CheckError(err)
	panic("Not implemented")
}

func Day11B(input string) int {
	panic("Not implemented")
}
