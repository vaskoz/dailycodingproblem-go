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
			var leftZero int
			for j := i - 1; j >= 0; j-- {
				if seating[j] == 0 {
					leftZero = j
					break
				}
			}
			for j := i - 1; j >= 0; j-- {
				if seating[j] == 1 && j < leftZero {
					curr += abs(j - leftZero)
					leftZero--
				}
			}
			// right
			var rightZero int
			for j := i + 1; j < len(seating); j++ {
				if seating[j] == 0 {
					rightZero = j
					break
				}
			}
			for j := i + 1; j < len(seating); j++ {
				if seating[j] == 1 && j > rightZero {
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
