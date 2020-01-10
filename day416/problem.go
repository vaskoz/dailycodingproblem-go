package day416

// Point is a X and Y coordinate.
type Point struct {
	X, Y int
}

// MinimumStepsPath returns the minimum number of steps required
// to walk the path. The path must be walked in order.
// 8 directions possible: 4 cardinal and 4 intercardinal.
func MinimumStepsPath(path []Point) int {
	if len(path) < 2 {
		return 0
	}

	steps := 0
	current := path[0]

	for i := 1; i < len(path); i++ {
		next := path[i]
		if dx, dy := delta(current, next); dx < dy {
			steps += dy
		} else {
			steps += dx
		}

		current = next
	}

	return steps
}

func delta(from, to Point) (int, int) {
	return abs(to.X - from.X), abs(to.Y - from.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
