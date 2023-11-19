package day15

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"os"
)

type Sensor struct {
	Position      utils.Point
	NearestBeacon utils.Point
}

type RangeSet struct {
	ranges []utils.Range
}

func (rs *RangeSet) insert(r utils.Range) {
	overlappingRangeIdx := -1

	for idx, curRange := range rs.ranges {
		if curRange.Overlaps(r) {
			overlappingRangeIdx = idx
			break
		}
	}

	if overlappingRangeIdx >= 0 {
		item := rs.ranges[overlappingRangeIdx]
		rs.ranges = utils.RemoveUnordered(rs.ranges, overlappingRangeIdx)

		item, _ = item.JoinWith(r)
		rs.insert(item)
	} else {
		rs.ranges = append(rs.ranges, r)
	}
}

func parseInput(path string) ([]Sensor, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var sensors []Sensor
	for scanner.Scan() {
		var sensor Sensor
		_, err = fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.Position.X, &sensor.Position.Y, &sensor.NearestBeacon.X, &sensor.NearestBeacon.Y)
		if err != nil {
			return nil, err
		}

		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func PartA(path string) int {
	sensors, err := parseInput(path)
	utils.CheckError(err)

	// Unique hash set of beacons (some are duplicates)
	beacons := map[utils.Point]bool{}
	for _, sensor := range sensors {
		beacons[sensor.NearestBeacon] = true
	}

	//const targetY = 10 // Sample
	const targetY = 2000000 // Production

	var rangeSet RangeSet
	for _, sensor := range sensors {
		// Calculate how far center is from its nearest beacon
		distanceToBeacon := utils.ManhattanDistance(sensor.Position, sensor.NearestBeacon)

		// Calculate how far it is to get to the row we're searching from the sensor (at minimum).
		// If it's farther than the distance from the sensor, then the "radius" around the beacon can't
		// extend to the target row so we can skip this sensor
		verticalDistance := utils.Abs(sensor.Position.Y - targetY)
		if verticalDistance > distanceToBeacon {
			continue
		}

		// Otherwise, sine we're manhattan distance; we can calculate the amount of the target row
		// taken up by the given sensor's search radius by figuring out how much of the radius's
		// "distance" we have to use up to get to the row; then the rest can be used fanning out
		// in either direction from that point
		horizontalDistance := distanceToBeacon - verticalDistance

		rangeInRadius := utils.Range{
			Start: sensor.Position.X - horizontalDistance,
			End:   sensor.Position.X + horizontalDistance,
		}
		rangeSet.insert(rangeInRadius)
	}

	// Figure out how many positions are in the range set,
	sum := 0
	for _, r := range rangeSet.ranges {
		sum += (r.End - r.Start + 1)
	}

	//Number of places beacons cannot be does NOT include any actual beacons on that line
	for beacon := range beacons {
		if beacon.Y == targetY {
			sum--
		}
	}

	return sum
}

func PartB(path string) string {
	_, err := parseInput(path)
	utils.CheckError(err)

	return "Not implemented"
}
