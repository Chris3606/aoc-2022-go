package utils

import (
	"errors"
	"math"

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

// Returns whether or not r1 contains the given number.
func (r1 Range) ContainsNum(num int) bool {
	return num >= r1.Start && num < r1.End
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

type Point3d struct {
	X int
	Y int
	Z int
}

func (p1 Point3d) Add(p2 Point3d) Point3d {
	return Point3d{p1.X + p2.X, p1.Y + p2.Y, p1.Z + p2.Z}
}

func (p1 Point3d) Sub(p2 Point3d) Point3d {
	return Point3d{p1.X - p2.X, p1.Y - p2.Y, p1.Z - p2.Z}
}

// 3d directions
var UP_3D = Point3d{0, -1, 0}
var RIGHT_3D = Point3d{1, 0, 0}
var DOWN_3D = Point3d{0, 1, 0}
var LEFT_3D = Point3d{-1, 0, 0}
var FORWARD_3D = Point3d{0, 0, 1}
var BACK_3D = Point3d{0, 0, -1}

var DIRS_3D = []Point3d{UP_3D, RIGHT_3D, DOWN_3D, LEFT_3D, FORWARD_3D, BACK_3D}

type Cube struct {
	MinExtent Point3d
	MaxExtent Point3d
}

// Returns whether or not the position given is contained within the cube.
func (cube Cube) Contains(point Point3d) bool {
	return point.X >= cube.MinExtent.X && point.X <= cube.MaxExtent.X &&
		point.Y >= cube.MinExtent.Y && point.Y <= cube.MaxExtent.Y &&
		point.Z >= cube.MinExtent.Z && point.Z <= cube.MaxExtent.Z
}

// Returns whether or not the position given is on the outer edge of the cube.
func (cube Cube) IsPerimeterPosition(point Point3d) bool {
	return cube.Contains(point) &&
		point.X == cube.MinExtent.X || point.X == cube.MaxExtent.X ||
		point.Y == cube.MinExtent.Y || point.Y == cube.MaxExtent.Y ||
		point.Z == cube.MinExtent.Z || point.Z == cube.MaxExtent.Z
}

// / Gets the bounding box for a set of 3d points
func GetBoundingBox3d(points []Point3d) Cube {
	bounds := Cube{Point3d{math.MaxInt, math.MaxInt, math.MaxInt}, Point3d{math.MinInt, math.MinInt, math.MinInt}}

	for _, point := range points {
		bounds.MinExtent.X = min(bounds.MinExtent.X, point.X)
		bounds.MinExtent.Y = min(bounds.MinExtent.Y, point.Y)
		bounds.MinExtent.Z = min(bounds.MinExtent.Z, point.Z)

		bounds.MaxExtent.X = max(bounds.MaxExtent.X, point.X)
		bounds.MaxExtent.Y = max(bounds.MaxExtent.Y, point.Y)
		bounds.MaxExtent.Z = max(bounds.MaxExtent.Z, point.Z)
	}

	return bounds
}
