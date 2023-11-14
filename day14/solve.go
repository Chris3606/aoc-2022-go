package day14

import (
	"aoc/utils"
	"bufio"
	"math"
	"os"
	"strings"
)

func parseInput(path string) (map[utils.Point]bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var points = map[utils.Point]bool{}

	for scanner.Scan() {
		var linePoints []utils.Point
		pointScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		pointScanner.Split(utils.ScanDelimiterFunc(" -> "))
		for pointScanner.Scan() {
			pointData := pointScanner.Text()

			point, err := utils.ReadPoint(strings.NewReader(pointData))
			if err != nil {
				return nil, err
			}
			linePoints = append(linePoints, point)
		}

		// Process lines into points
		for i := 0; i < len(linePoints)-1; i++ {
			p1 := linePoints[i]
			p2 := linePoints[i+1]

			if p1.X != p2.X {
				for x := min(p1.X, p2.X); x <= max(p1.X, p2.X); x++ {
					points[utils.Point{X: x, Y: p1.Y}] = true
				}
			} else {
				for y := min(p1.Y, p2.Y); y <= max(p1.Y, p2.Y); y++ {
					points[utils.Point{X: p1.X, Y: y}] = true
				}
			}
		}
	}

	return points, nil
}

func FindRestingPointForSandGrain(objects map[utils.Point]bool, floorY int, entryPoint utils.Point) (utils.Point, bool) {

	// Fall until we hit something or know we hit the floor
	cur := entryPoint
	for {
		if cur.Y == floorY-1 { // Hit the floor
			break
		}

		// Check neighbors in proper order
		n := cur.Add(utils.DOWN)
		if !objects[n] {
			cur = n
			continue
		}

		n = cur.Add(utils.DOWN_LEFT)
		if !objects[n] {
			cur = n
			continue
		}

		n = cur.Add(utils.DOWN_RIGHT)
		if !objects[n] {
			cur = n
			continue
		}

		// The grain can't move anywhere, so we found final point and it's not the floor
		return cur, false
	}
	return cur, true // Came to rest on floor
}

// func FindRestingPointForSandGrain(objects map[utils.Point]bool, maxY int, entryPoint utils.Point) (utils.Point, bool) {

// 	// Fall until we hit something or know we're below the floor
// 	cur := entryPoint
// 	for {
// 		// Sand will fall into infinity since there's nothing left below it to block it
// 		if cur.Y > maxY {
// 			break
// 		}

// 		// Check neighbors in proper order
// 		n := cur.Add(utils.DOWN)
// 		if !objects[n] {
// 			cur = n
// 			continue
// 		}

// 		n = cur.Add(utils.DOWN_LEFT)
// 		if !objects[n] {
// 			cur = n
// 			continue
// 		}

// 		n = cur.Add(utils.DOWN_RIGHT)
// 		if !objects[n] {
// 			cur = n
// 			continue
// 		}

// 		// The grain can't move anywhere, so we found final point for sand grain
// 		return cur, true
// 	}
// 	return utils.Point{}, false // No resting point
// }

func PartA(path string) int {
	objects, err := parseInput(path)
	utils.CheckError(err)

	// Find max y value; anything that falls below this will never be stopped, so we'll consider
	// it the floor
	floorY := math.MinInt
	for k := range objects {
		floorY = max(floorY, k.Y)
	}
	floorY += 1

	// Simulate sand until one falls into the abyss
	var sandGrains int
	for {
		point, restedOnFloor := FindRestingPointForSandGrain(objects, floorY, utils.Point{X: 500, Y: 0})
		if restedOnFloor { // Falls infinitely
			break
		}

		sandGrains++
		objects[point] = true
	}

	return sandGrains
}

func PartB(path string) int {
	entryPoint := utils.Point{X: 500, Y: 0}

	objects, err := parseInput(path)
	utils.CheckError(err)

	// Find floor level
	floorY := math.MinInt
	for k := range objects {
		floorY = max(floorY, k.Y)
	}
	floorY += 2

	// Simulate sand until one comes to rest at (500, 0)
	var sandGrains int
	for {
		point, _ := FindRestingPointForSandGrain(objects, floorY, entryPoint)

		sandGrains++
		if point == entryPoint {
			break
		}
		objects[point] = true
	}

	return sandGrains
}
