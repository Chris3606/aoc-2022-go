package utils

import (
	"bufio"
	"io"
)

func ReadGroups[T any](r io.Reader, parser func(string) (T, error)) ([][]T, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var result [][]T
	var curSlice []T
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, err := parser(scanner.Text())
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

// Reads grid in the following format
// 012345
// 654568
// 923598
func ReadDigitGrid(r io.Reader) (Grid[byte], error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var slice []byte
	var width int
	for scanner.Scan() {
		text := scanner.Text()
		if width == 0 {
			width = len(text)
		}

		for i := range text {
			slice = append(slice, text[i]-'0')
		}
	}

	return GridFromSlice[byte](slice, width), nil
}
