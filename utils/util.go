package utils

import (
	"errors"

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
var UP_RIGHT = Point{1, -1}
var RIGHT = Point{1, 0}
var DOWN_RIGHT = Point{1, 1}
var DOWN = Point{0, 1}
var DOWN_LEFT = Point{-1, 1}
var LEFT = Point{-1, 0}
var UP_LEFT = Point{-1, -1}

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

// Calculates chebyshev distance (aka distance where the distance between a cell and all 8 of a cells
// neighbors is 1)
func ChebyshevDistance(p1, p2 Point) int {
	return max(Abs(p2.X-p1.X), Abs(p2.Y-p1.Y))
}

// Calculates manhattan distance; aka distance where only the neighbors in 4 cardinal directions are
// adjacent.
func ManhattanDistance(p1, p2 Point) int {
	return Abs(p2.X-p1.X) + Abs(p2.Y-p1.Y)
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

// Represents a range of numbers (both ends are inclusive)
type Range struct {
	Start int
	End   int
}

// Returns whether or not r1 contains r2 entirely.
func (r1 Range) Contains(r2 Range) bool {
	return r1.Start <= r2.Start && r1.End >= r2.End
}

// Returns whether or not r1 and r2 overlap.
func (r1 Range) Overlaps(r2 Range) bool {
	return r1.Start <= r2.End && r2.Start <= r1.End
}

// Joins 2 overlapping ranges
func (r1 Range) JoinWith(r2 Range) (Range, error) {
	if !r1.Overlaps(r2) {
		return Range{}, errors.New("you cannot join non-overlapping ranges")
	}

	return Range{min(r1.Start, r2.Start), max(r1.End, r2.End)}, nil
}

// Removes the item from a slice in O(1) time, without guaranteeing preservation of order
func RemoveUnordered[T any](slice []T, idx int) []T {
	lastIdx := len(slice) - 1
	slice[idx] = slice[lastIdx]

	return slice[:lastIdx]
}
