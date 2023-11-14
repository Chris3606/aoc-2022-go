package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
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
//
// The values can be any arbitrary byte, whether those are characters, actual bytes, etc.
// The parsing function must translate the values to the appropriate result type.
func ReadGridFromBytes[T any](r io.Reader, parser func(byte, Point) (T, error)) (Grid[T], error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	var slice []T
	var width int

	var y int
	for scanner.Scan() {
		text := scanner.Text()
		if width == 0 {
			width = len(text)
		}

		for x := range text {
			val, err := parser(text[x], Point{x, y})
			if err != nil {
				return Grid[T]{}, err
			}
			slice = append(slice, val)
		}

		y++
	}

	return GridFromSlice[T](slice, width), nil
}

// Reads points in form 1,2
func ReadPoint(r io.Reader) (Point, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanDelimiterFunc(","))

	b := scanner.Scan()
	if !b {
		return Point{}, errors.New("invalid point format: X coordinate not found")
	}
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return Point{}, err
	}

	b = scanner.Scan()
	if !b {
		return Point{}, errors.New("invalid point format: Y coordinate not found")
	}
	y, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return Point{}, err
	}

	return Point{x, y}, nil
}

func ScanDelimiterFunc(separator string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte(separator)
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}

// Reads in a list of items separated by the given separator
func ReadItems[T any](r io.Reader, separator string, parser func(string) (T, error)) ([]T, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(ScanDelimiterFunc(separator))

	var results []T
	for scanner.Scan() {
		val, err := parser(scanner.Text())
		if err != nil {
			return nil, err
		}

		results = append(results, val)
	}

	return results, nil
}
