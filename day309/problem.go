package day309

// MinRedistributeNoGapsBrute returns the minimum cost
// to move all together.
// Runs in O(N^2) time and O(1) space.
func MinRedistributeNoGapsBrute(seating []int) int {
	min := int(^uint(0) >> 1)
	for i := range seating {
		var curr int
		if seating[i] == 1 {
			// left
			leftZero := 0
			leftZeroFound := false
			for j := i - 1; j >= 0; j-- {
				if !leftZeroFound && seating[j] == 0 {
					leftZero = j
					leftZeroFound = true
				} else if seating[j] == 1 && j < leftZero {
					curr += abs(j - leftZero)
					leftZero--
				}
			}
			// right
			rightZero := 0
			rightZeroFound := false
			for j := i + 1; j < len(seating); j++ {
				if !rightZeroFound && seating[j] == 0 {
					rightZero = j
					rightZeroFound = true
				} else if seating[j] == 1 && j > rightZero {
					curr += abs(j - rightZero)
					rightZero--
				}
			}
			if curr < min {
				min = curr
			}
		}
	}
	return min
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
