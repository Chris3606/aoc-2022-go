package utils

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

type Point struct {
	X int
	Y int
}

// Directions
var UP = Point{0, -1}
var RIGHT = Point{1, 0}
var DOWN = Point{0, 1}
var LEFT = Point{-1, 0}

func FromIndex(index, width int) Point {
	return Point{index % width, index / width}
}

func (point Point) ToIndex(width int) int {
	return point.Y*width + point.X
}

func (p1 Point) Add(p2 Point) Point {
	return Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 Point) Sub(p2 Point) Point {
	return Point{p1.X - p2.X, p1.Y - p2.Y}
}

type Grid[T any] struct {
	Slice  []T
	width  int
	height int
}

func GridFromDimensions[T any](width, height int) Grid[T] {
	slice := make([]T, width*height)
	return GridFromSlice(slice, width)
}

func GridFromSlice[T any](slice []T, width int) Grid[T] {
	return Grid[T]{slice, width, len(slice) / width}
}

func (grid *Grid[T]) Width() int {
	return grid.width
}

func (grid *Grid[T]) Height() int {
	return grid.height
}

func (grid *Grid[T]) GetCopy(position Point) T {
	return grid.Slice[position.ToIndex(grid.width)]
}

func (grid *Grid[T]) Get(position Point) *T {
	return &grid.Slice[position.ToIndex(grid.width)]
}

func (grid *Grid[T]) Set(position Point, value T) {
	grid.Slice[position.ToIndex(grid.width)] = value
}

func (grid *Grid[T]) Positions() GridPositionsIterator[T] {
	return GridPositionsIterator[T]{-1, grid}
}

func (grid *Grid[T]) Fill(value T) {
	for i := range grid.Slice {
		grid.Slice[i] = value
	}
}

type GridPositionsIterator[T any] struct {
	index int
	grid  *Grid[T]
}

func (it *GridPositionsIterator[T]) Next() bool {
	it.index++
	return it.index < len(it.grid.Slice)
}

func (it *GridPositionsIterator[T]) Current() Point {
	return FromIndex(it.index, it.grid.width)
}

func Abs(v1 int) int {
	if v1 < 0 {
		return -v1
	}

	return v1
}

func ChebyshevDistance(p1 Point, p2 Point) int {
	return max(Abs(p2.X-p1.X), Abs(p2.Y-p1.Y))
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
