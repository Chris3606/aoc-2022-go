package main

import (
	"os"
	"slices"
	"sort"
)

func ParseInput(input string) ([]int, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	backpacks, err := ReadIntGroups(f)
	if err != nil {
		return nil, err
	}

	var sums []int
	for _, backpack := range backpacks {
		sums = append(sums, SumInts(backpack))
	}

	return sums, nil
}

func Day01A(input string) int {
	backpackSums, err := ParseInput(input)
	CheckError(err)

	return slices.Max(backpackSums)
}

func Day01B(input string) int {
	backpackSums, err := ParseInput(input)
	CheckError(err)

	sort.Slice(backpackSums, func(i, j int) bool {
		return backpackSums[i] > backpackSums[j]
	})

	return SumInts(backpackSums[0:3])
}
