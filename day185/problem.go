package day185

// Rectangle with a top-left coordinate plus width and height.
type Rectangle struct {
	topLeftX, topLeftY int
	width, height      int
}

// AreaOfIntersection returns the area of intersection of two
// rectangles.
func AreaOfIntersection(r1, r2 Rectangle) int {
	leftMostR := r1
	rightMostR := r2
	if r2.topLeftX < r1.topLeftX {
		leftMostR = r2
		rightMostR = r1
	}
	deltaX := 0
	if leftMostRR := leftMostR.topLeftX + leftMostR.width; leftMostRR > rightMostR.topLeftX {
		deltaX = leftMostRR - rightMostR.topLeftX
	}
	if deltaX == 0 {
		return 0
	}
	topMostR := r1
	bottomMostR := r2
	if r2.topLeftY < r1.topLeftY {
		topMostR = r2
		bottomMostR = r1
	}
	deltaY := 0
	if topMostRR := topMostR.topLeftY + topMostR.height; topMostRR > bottomMostR.topLeftY {
		deltaY = topMostRR - bottomMostR.topLeftY
	}
	if deltaY == 0 {
		return 0
	}
	return deltaX * deltaY
}
