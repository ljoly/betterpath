package betterpath

import (
	"math"
)

// StdDeviation returns the standard deviation of the distances
func StdDeviation(distances []float64, mean float64, zeroDistances int) float64 {

	std := 0.0
	for _, d := range distances {
		abs := math.Abs(d - mean)
		std += abs * abs
	}
	std /= float64(len(distances)) - float64(zeroDistances)
	std = math.Sqrt(std)
	return std
}

// Mean returns the mean of the non-zero distances
func Mean(distances []float64, zeroDistances int) float64 {

	mean := 0.0
	for _, d := range distances {
		mean += d
	}
	mean /= float64(len(distances)) - float64(zeroDistances)
	return mean
}

// Distance returns the distance between two points
func Distance(a Point, b Point) float64 {

	return math.Sqrt(math.Pow((a.x-b.x), 2) + math.Pow((a.y-b.y), 2))
}

// Distances returns an array of distances and the number of zero distances
func Distances(points Points) ([]float64, int) {

	distances := make([]float64, len(points)-1)
	iterDistance := 0.0
	zeroDistances := 0
	for i := 0; i < len(points)-1; i++ {

		iterDistance = Distance(points[i+1], points[i])
		if iterDistance == 0 {
			zeroDistances++
		}
		distances[i] = iterDistance
	}
	return distances, zeroDistances
}
