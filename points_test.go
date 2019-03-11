package betterpath

import (
	"testing"
)

func TestRemoveOutliers(t *testing.T) {

	points := make(Points, 6)

	// TEST1: Wrong points at the end
	points[0] = Point{1, 1, 1}
	points[1] = Point{2, 2, 2}
	points[2] = Point{3, 3, 3}
	points[3] = Point{4, 4, 4}
	points[4] = Point{25, 25, 5}
	points[5] = Point{25, 25, 6}

	points.RemoveOutliers()
	if len(points) != 4 {
		t.Error("Test1: All erroneous points were not removed")
	} else {
		want := [4]int64{1, 2, 3, 4}
		for i, got := range points {
			if want[i] != got.t {
				t.Error("Test1: Wrong points were removed")
			}
		}
	}

	points = make(Points, 6)

	// TEST2: Wrong point at the middle
	points[0] = Point{1, 1, 1}
	points[1] = Point{2, 2, 2}
	points[2] = Point{25, 25, 3}
	points[3] = Point{4, 4, 4}
	points[4] = Point{5, 5, 5}
	points[5] = Point{6, 6, 6}

	points.RemoveOutliers()
	if len(points) != 5 {
		t.Error("Test2: All erroneous points were not removed")
	} else {
		want := [5]int64{1, 2, 4, 5, 6}
		for i, got := range points {
			if want[i] != got.t {
				t.Error("Test2: Wrong points were removed")
			}
		}
	}

	points = make(Points, 6)

	// TEST3: Wrong point at the beginning
	points[0] = Point{25, 25, 1}
	points[1] = Point{2, 2, 2}
	points[2] = Point{3, 3, 3}
	points[3] = Point{4, 4, 4}
	points[4] = Point{5, 5, 5}
	points[5] = Point{6, 6, 6}

	points.RemoveOutliers()
	if len(points) != 5 {
		t.Error("Test3: All erroneous points were not removed")
	} else {
		want := [5]int64{2, 3, 4, 5, 6}
		for i, got := range points {
			if want[i] != got.t {
				t.Error("Test3: Wrong points were removed")
			}
		}
	}

	points = make(Points, 6)

	// TEST4: Nothing to remove
	points[0] = Point{1, 1, 1}
	points[1] = Point{2, 2, 2}
	points[2] = Point{3, 3, 3}
	points[3] = Point{4, 4, 4}
	points[4] = Point{5, 5, 5}
	points[5] = Point{6, 6, 6}

	points.RemoveOutliers()
	if len(points) != 6 {
		t.Error("Test4: All erroneous points were not removed")
	} else {
		want := [6]int64{1, 2, 3, 4, 5, 6}
		for i, got := range points {
			if want[i] != got.t {
				t.Error("Test4: Wrong points were removed")
			}
		}
	}
}
