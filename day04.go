package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func RangeFromStr(str string) (Range, error) {
	elems := strings.Split(str, "-")

	start, err := strconv.Atoi(elems[0])
	if err != nil {
		return Range{}, err
	}

	end, err := strconv.Atoi(elems[1])
	if err != nil {
		return Range{}, err
	}

	return Range{start, end}, nil
}

func (r1 Range) Contains(r2 Range) bool {
	return r1.Start <= r2.Start && r1.End >= r2.End
}

func (r1 Range) Overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r2.Start <= r1.End
}

type ElfPair struct {
	range1 Range
	range2 Range
}

func parseInput4(input string) ([]ElfPair, error) {
	f, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var pairs []ElfPair
	for scanner.Scan() {
		rangePair := strings.Split(scanner.Text(), ",")
		r1, err := RangeFromStr(rangePair[0])
		if err != nil {
			return nil, err
		}
		r2, err := RangeFromStr(rangePair[1])
		if err != nil {
			return nil, err
		}

		pairs = append(pairs, ElfPair{r1, r2})
	}

	return pairs, nil
}

func Day04A(input string) int {
	pairs, err := parseInput4(input)
	CheckError(err)

	var pairsWithSubset int
	for _, pair := range pairs {
		if pair.range1.Contains(pair.range2) || pair.range2.Contains(pair.range1) {
			pairsWithSubset++
		}
	}

	return pairsWithSubset
}

func Day04B(input string) int {
	pairs, err := parseInput4(input)
	CheckError(err)

	var overlappingPairs int
	for _, pair := range pairs {
		if pair.range1.Overlaps(pair.range2) {
			overlappingPairs++
		}
	}

	return overlappingPairs
}
