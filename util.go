package main

import (
	"bufio"
	"io"
	"strconv"
)

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadIntGroups(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result [][]int
	var curSlice []int
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return result, err
			}

			curSlice = append(curSlice, val)
		} else {
			result = append(result, curSlice)
			curSlice = nil
		}
	}

	if curSlice != nil {
		result = append(result, curSlice)
	}

	return result, scanner.Err()
}

func SumInts(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum
}

// TODO: Generic function
func BuildHistogram(slice []byte) map[byte]int {
	hist := map[byte]int{}

	for _, v := range slice {
		hist[v] += 1
	}

	return hist
}
