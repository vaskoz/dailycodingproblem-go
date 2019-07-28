package day340

import "math"

// Point is a 2D cartesian point.
type Point struct {
	X, Y int
}

// TwoClosestPointsBrute returns the two closest points.
// It runs in O(N^2) time and O(1) extra space.
// Just compares each point against the others.
func TwoClosestPointsBrute(pts []Point) []Point {
	if len(pts) < 2 {
		return nil
	} else if len(pts) == 2 {
		return pts
	}
	res := make([]Point, 2)
	dist := math.MaxFloat64
	for i := 0; i < len(pts)-1; i++ {
		for j := i + 1; j < len(pts); j++ {
			d := math.Sqrt(
				math.Pow(float64(pts[i].X-pts[j].X), 2.0) +
					math.Pow(float64(pts[i].Y-pts[j].Y), 2.0))
			if d < dist {
				dist = d
				res[0] = pts[i]
				res[1] = pts[j]
			}
		}
	}
	return res
}
