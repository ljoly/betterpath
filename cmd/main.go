package main

import (
	"fmt"
	"os"

	"github.com/ljoly/betterpath"
)

func main() {

	points := make(betterpath.Points, 0)

	err := points.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	points.SortByTimestamp()
	points.RemoveOutliers()
	points.Print()
}
