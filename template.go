package main

import (
	"aoc/utils"
	"bufio"
	"os"
)

func parseInputX(input string) ([]int, error) {
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

func Day0XA(input string) int {
	_, err := parseInputX(input)
	utils.CheckError(err)
	panic("Not implemented")
}

func Day0XB(input string) int {
	panic("Not implemented")
}
