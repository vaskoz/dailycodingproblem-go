package day236

// Point is a cartesian 2D point.
type Point struct {
	X, Y float64
}

// LiesInsidePolygon answers if the given point lies inside the polygon.
func LiesInsidePolygon(polygon []Point, p Point) bool {
	polygon = append(polygon, polygon[0]) // to link back to first point
	intersections := make(map[Point]struct{})
	for i := 1; i < len(polygon); i++ {
		low, high := polygon[i-1], polygon[i]
		if low.Y > high.Y {
			low, high = high, low
		}
		if low.Y == high.Y || p.Y < low.Y || p.Y > high.Y {
			continue // skip horizontal lines and too high, too low
		}
		if low.X == high.X { // vertical line
			if p.X == low.X && p.Y >= low.Y && p.Y <= high.Y {
				return false // on boundary
			} else if p.X < low.X && p.Y >= low.Y && p.Y <= high.Y {
				intersections[Point{low.X, p.Y}] = struct{}{}
			}
			continue
		}
		slope := (high.Y - low.Y) / (high.X - low.X)
		b := low.Y - slope*low.X
		xAtY := (p.Y - b) / slope
		if p.X == xAtY {
			return false // on boundary
		} else if p.X < xAtY {
			intersections[Point{xAtY, p.Y}] = struct{}{}
		}
	}
	return len(intersections) == 1
}
