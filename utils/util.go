package utils

import (
	"golang.org/x/exp/constraints"
)

type Summable interface {
	constraints.Ordered | constraints.Complex | string
}

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func Sum[T Summable](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}

	return sum
}

func BuildHistogram[T comparable](slice []T) map[T]int {
	hist := map[T]int{}

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

func Abs(v1 int) int {
	if v1 < 0 {
		return -v1
	}

	return v1
}

func ChebyshevDistance(p1 Point, p2 Point) int {
	return max(Abs(p2.X-p1.X), Abs(p2.Y-p1.Y))
}
