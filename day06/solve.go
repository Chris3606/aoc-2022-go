package day06

import (
	"aoc/utils"
	"io"
	"os"
)

func parseInput(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func findMarkerIndex(data []byte, markerSize int) int {
	for i := 0; i <= len(data)-markerSize; i++ {
		hist := utils.BuildHistogram(data[i : i+markerSize])
		if len(hist) == markerSize {
			return i + markerSize
		}
	}

	panic("No marker found.")
}

func PartA(path string) int {
	data, err := parseInput(path)
	utils.CheckError(err)

	return findMarkerIndex(data, 4)
}

func PartB(path string) int {
	data, err := parseInput(path)
	utils.CheckError(err)

	return findMarkerIndex(data, 14)
}
