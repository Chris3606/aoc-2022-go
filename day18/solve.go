package day18

import (
	"aoc/utils"
	"bufio"
	"os"
)

func parseInput(path string) ([]utils.Point3d, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return utils.ReadItemsScan(bufio.NewScanner(f), utils.ReadPoint3d)
}

func countSidesVisible(points map[utils.Point3d]bool) int {
	sidesVisible := 0
	for item := range points {
		for _, dir := range utils.DIRS_3D {
			neighbor := item.Add(dir)

			if !points[neighbor] {
				sidesVisible++
			}
		}
	}

	return sidesVisible
}

func PartA(path string) int {
	points, err := parseInput(path)
	utils.CheckError(err)

	var pointsSet = map[utils.Point3d]bool{}
	for _, item := range points {
		pointsSet[item] = true
	}

	return countSidesVisible(pointsSet)
}

func PartB(path string) int {
	_, err := parseInput(path)
	utils.CheckError(err)

	points, err := parseInput(path)
	utils.CheckError(err)

	// Create hash set of points so we can quickly check if points are in the set or not
	var pointsSet = map[utils.Point3d]bool{}
	for _, item := range points {
		pointsSet[item] = true
	}

	// Find bounds of the points
	bounds := utils.GetBoundingBox3d(points)

	// List of points we know are part of a processed area
	areaMap := map[utils.Point3d]bool{}

	// Go through all points in bounding cube.
	// - If they're not part of an area, flood fill from them to the bounding box limits.
	//     - If the flood-fill encounters points in pointsSet on every side, then
	//       the space we flood filled is an enclosed pocket.
	//     - Otherwise, the flood fill stopped because we hit the edge of the mapped area,
	//       so the space is not enclosed (there could be no blocking points beyond the bounds
	//       of all known points)
	enclosedPockets := 0
	for x := bounds.MinExtent.X; x <= bounds.MaxExtent.X; x++ {
		for y := bounds.MinExtent.Y; y <= bounds.MaxExtent.Y; y++ {
			for z := bounds.MinExtent.Z; z <= bounds.MaxExtent.Z; z++ {
				curPoint := utils.Point3d{X: x, Y: y, Z: z}

				// Not empty space so can't be part of a region, or part of a region we already
				// processed
				if pointsSet[curPoint] || areaMap[curPoint] {
					continue
				}

				// Flood fill from area
				curArea := map[utils.Point3d]bool{}

				hitBound := false
				stack := []utils.Point3d{curPoint}
				for len(stack) > 0 {
					ffPoint := stack[0]
					stack = stack[1:]

					// Already processed
					if curArea[ffPoint] {
						continue
					}

					// Hit a point originally in map; space may be enclosed and this is a boundary
					if pointsSet[ffPoint] {
						continue
					}

					// Hit edge of bounds (and wasn't a point in the map); so the space we're
					// processing is not enclosed, and there is no need to continue from here.
					// We'll continue to process the area just to prevent having to re-do the flood
					// fill on this area when we encounter other points within it in the future.
					if bounds.IsPerimeterPosition(ffPoint) {
						hitBound = true
						continue
					}

					// Point is in area, so add it as appropriate
					areaMap[ffPoint] = true
					curArea[ffPoint] = true

					// Process all relevant neighbors and add to stack
					for _, dir := range utils.DIRS_3D {
						neighbor := ffPoint.Add(dir)
						if curArea[neighbor] {
							continue
						}

						stack = append(stack, neighbor)
					}
				}

				// Area is enclosed by squares in the map; so add its total to the total enclosed
				// spaces
				if !hitBound {
					enclosedPockets += len(curArea)
				}
			}
		}
	}

	// TODO: Have to determine the proper number to subtract here
	return enclosedPockets
}
