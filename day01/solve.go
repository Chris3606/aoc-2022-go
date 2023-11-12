package day01

import (
	"aoc/utils"
	"os"
	"slices"
	"sort"
)

func parseInput(path string) ([]int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	backpacks, err := utils.ReadIntGroups(f)
	if err != nil {
		return nil, err
	}

	var sums []int
	for _, backpack := range backpacks {
		sums = append(sums, utils.SumInts(backpack))
	}

	return sums, nil
}

func PartA(path string) int {
	backpackSums, err := parseInput(path)
	utils.CheckError(err)

	return slices.Max(backpackSums)
}

func PartB(path string) int {
	backpackSums, err := parseInput(path)
	utils.CheckError(err)

	sort.Slice(backpackSums, func(i, j int) bool {
		return backpackSums[i] > backpackSums[j]
	})

	return utils.SumInts(backpackSums[0:3])
}
