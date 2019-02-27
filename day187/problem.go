package day187

// Rectangle is a representation of a rectangle.
type Rectangle struct {
	TopLeftX, TopLeftY int
	Width, Height      int
}

// DoesRectanglePairOverlap returns true if any rectangles
// overlap.
// Runs in O(N^2) time.
func DoesRectanglePairOverlap(rects []Rectangle) bool {
	for i, left := range rects {
		for j, right := range rects {
			if j > i {
				if left.TopLeftX > right.TopLeftX {
					left, right = right, left
				}
				if right.TopLeftX < left.TopLeftX+left.Width {
					up, dn := left, right
					if up.TopLeftY < dn.TopLeftY {
						up, dn = dn, up
					}
					if dn.TopLeftY > up.TopLeftY-up.Height {
						return true
					}
				}
			}
		}
	}
	return false
}
