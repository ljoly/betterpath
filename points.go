package betterpath

import (
	"fmt"
	"math"
	"sort"
)

// Point stores coordinates and timestamp of a point
type Point struct {
	x float64
	y float64
	t int64
}

// Points stores the path
type Points []Point

func (points Points) averageDistance() float64 {

	d := 0.0
	xPrev := 0.0
	yPrev := 0.0
	for _, point := range points {

		d += math.Sqrt(math.Pow((point.x-xPrev), 2) + math.Pow((point.y-yPrev), 2))
		xPrev = point.x
		yPrev = point.y
	}
	d /= float64(len(points))
	return d
}

// CheckPoints removes erroneous points
func (points Points) CheckPoints() {

	// average distance between two points
	avDist := points.averageDistance()

	fmt.Println(avDist)
}

// SortByTimestamp sorts points by timestsamp
func (points Points) SortByTimestamp() {

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].t < points[j].t
	})
}
