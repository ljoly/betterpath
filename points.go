package betterpath

import (
	"fmt"
	"sort"
)

// Point stores coordinates and timestamp of a point
type Point struct {
	x float64
	y float64
	t int64
}

// Points stores the coordinates sorted by timestamp
type Points []Point

// Print outputs the updated dataset
func (points Points) Print() {

	for _, point := range points {
		fmt.Print(point.x, ",", point.y, ",", point.t, "\n")
	}
}

// RemoveOutliers removes erroneous points by considering their distance
// to the standard deviation of distances
func (points *Points) RemoveOutliers() {

	distances, zeroDistances := Distances(*points)
	mean := Mean(distances, zeroDistances)
	stdDeviation := StdDeviation(distances, mean, zeroDistances)

	if stdDeviation == 0 {
		// nothing to be done
		return
	}
	// check the first point
	for distances[0] > stdDeviation && distances[1] < stdDeviation {
		// pop the first point
		_, *points = (*points)[0], (*points)[1:]
		_, distances = distances[0], distances[1:]
		distances[0] = Distance((*points)[0], (*points)[1])
	}
	// check all the points
	for i := 1; i < len(distances)-1; i++ {
		if distances[i] > stdDeviation {
			// remove the second point of the pair considered
			*points = append((*points)[:i+1], (*points)[i+2:]...)
			// remove the index of the distance considered
			distances = append(distances[:i], distances[i+1:]...)
			// update the array of distances with the distance of the new pair
			distances[i] = Distance((*points)[i], (*points)[i+1])
			i--
		}
	}
	// check the last point
	if distances[len(distances)-1] > stdDeviation {
		*points = (*points)[:len(*points)-1]
	}
}

// SortByTimestamp sorts points by timestsamp
func (points Points) SortByTimestamp() {

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].t < points[j].t
	})
}
