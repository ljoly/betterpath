package main

import (
	bp "github.com/ljoly/betterpath"
)

func main() {

	points := make(bp.Points, 0)
	points.Parse()
	points.CheckPoints()
}
