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

	for i := 1; i < len(*points)-1; i++ {

		if distances[i-1] > stdDeviation {
			// remove the second point of the pair considered
			*points = append((*points)[:i], (*points)[i+1:]...)
			// remove the index of the distance considered
			distances = append(distances[:i-1], distances[i:]...)
			// update the array of distances with the distance of the new pair
			distances[i] = Distance((*points)[i], (*points)[i+1])
		}
	}
}

// SortByTimestamp sorts points by timestsamp
func (points Points) SortByTimestamp() {

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].t < points[j].t
	})
}
