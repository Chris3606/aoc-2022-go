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

var CARDINAL_DIRS_CLOCKWISE = []Point{UP, RIGHT, DOWN, LEFT}

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

func ChebyshevDistance(p1, p2 Point) int {
	return max(Abs(p2.X-p1.X), Abs(p2.Y-p1.Y))
}

func ManhattanDistance(p1, p2 Point) int {
	return Abs(p2.X-p1.X) + Abs(p2.Y+p1.Y)
}

// Greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
