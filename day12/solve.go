package day12

import (
	"aoc/utils"
	"bufio"
	"math"
	"os"

	"github.com/oleiade/lane/v2"
)

type AStarItem struct {
	Position utils.Point
	F        int
	G        int
}

func (item *AStarItem) Priority() int {
	return item.F
}

func shortestPath(grid utils.Grid[byte], start, end utils.Point) (int, bool) {
	dist := utils.GridFromDimensions[int](grid.Width(), grid.Height())
	dist.Fill(math.MaxInt)

	heap := lane.NewMinPriorityQueue[AStarItem, int]()

	dist.Set(start, 0)
	itm := AStarItem{Position: start, G: 0, F: utils.ManhattanDistance(start, end)}
	heap.Push(itm, itm.Priority())

	for !heap.Empty() {
		// Pop item
		item, _, ok := heap.Pop()
		if !ok {
			panic("Remove from priority queue failed.")
		}

		// We found the shortest path
		if item.Position == end {
			return item.G, true
		}

		// If we've already found a better way, we won't visit this node on the current path;
		// this can happen if multiple states with the same value were pushed into the queue
		if item.G > dist.GetCopy(item.Position) {
			continue
		}

		// Test all neighbors to see if there is a better path to them by going though the current
		// position
		for _, dir := range utils.CARDINAL_DIRS_CLOCKWISE {
			neighbor := item.Position.Add(dir)
			// Off edge of grid
			if !grid.Contains(neighbor) {
				continue
			}

			// Location is too much elevation change
			if grid.GetCopy(neighbor) > (grid.GetCopy(item.Position) + 1) {
				continue
			}

			node := AStarItem{Position: neighbor, G: item.G + 1, F: item.G + 1 + utils.ManhattanDistance(neighbor, end)}

			// If cost is lower, add it to the list of nodes to visit and update cost
			if node.G < dist.GetCopy(neighbor) {
				dist.Set(node.Position, node.G)
				heap.Push(node, node.Priority())
			}
		}
	}

	return -1, false
}

func parseInput(path string) (utils.Point, utils.Point, utils.Grid[byte], error) {
	f, err := os.Open(path)
	if err != nil {
		return utils.Point{}, utils.Point{}, utils.Grid[byte]{}, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var start, end utils.Point
	grid, err := utils.ReadGridFromBytes(f, func(b byte, p utils.Point) (byte, error) {
		if b == 'S' {
			b = 'a'
			start = p
		} else if b == 'E' {
			b = 'z'
			end = p
		}

		return b - 'a', nil
	})
	if err != nil {
		if err != nil {
			return utils.Point{}, utils.Point{}, utils.Grid[byte]{}, err
		}
	}

	return start, end, grid, nil
}

func PartA(path string) int {
	start, end, grid, err := parseInput(path)
	utils.CheckError(err)

	pathLen, ok := shortestPath(grid, start, end)
	if !ok {
		panic("No path found from start to end")
	}
	return pathLen
}

func PartB(path string) int {
	_, end, grid, err := parseInput(path)
	utils.CheckError(err)

	// Depending on map semantics, dijkstra's may be more efficient here (calculate from end to
	// all other points); but, we've already implemented AStar, so might as well use it,
	// it's fast enough
	var startLocations []utils.Point
	posIt := grid.Positions()
	for posIt.Next() {
		pos := posIt.Current()
		if grid.GetCopy(pos) == 0 {
			startLocations = append(startLocations, pos)
		}
	}

	minLen := math.MaxInt
	for _, startLocation := range startLocations {
		curLen, ok := shortestPath(grid, startLocation, end)
		if !ok {
			continue
		}
		minLen = min(minLen, curLen)
	}

	return minLen
}
