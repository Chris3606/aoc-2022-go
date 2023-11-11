package main

import (
	"io"
	"os"
)

func parseInput6(input string) ([]byte, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func findMarkerIndex(data []byte, markerSize int) int {
	for i := 0; i <= len(data)-markerSize; i++ {
		hist := BuildHistogram(data[i : i+markerSize])
		if len(hist) == markerSize {
			return i + markerSize
		}
	}

	panic("No marker found.")
}

func Day06A(input string) int {
	data, err := parseInput6(input)
	CheckError(err)

	return findMarkerIndex(data, 4)
}

func Day06B(input string) int {
	data, err := parseInput6(input)
	CheckError(err)

	return findMarkerIndex(data, 14)
}
