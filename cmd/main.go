package main

import (
	"github.com/ljoly/betterpath"
)

func main() {

	points := make(betterpath.Points, 0)
	points.Parse()
	points.SortByTimestamp()
	points.RemoveOutliers()
	points.Print()
}
