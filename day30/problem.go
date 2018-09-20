package day30

// TrappedWater returns the amount of water that can be trapped.
// Runs in O(N) time and O(1) space.
func TrappedWater(heights []int) int {
	if len(heights) < 3 {
		return 0
	}
	total := 0
	hl := heights[0]
	hri, hr := highestToRight(heights, 2)
	for i := 1; i < len(heights)-1; i++ {
		h := min(hl, hr)
		if cavity := h - heights[i]; cavity > 0 {
			total += cavity
		}
		if heights[i] > hl {
			hl = heights[i]
		}
		if i+1 == hri {
			hri, hr = highestToRight(heights, hri)
		}
	}
	return total
}

func highestToRight(heights []int, start int) (int, int) {
	var highest, highestI int
	for i := start; i < len(heights); i++ {
		if heights[i] > highest {
			highest = heights[i]
			highestI = i
		}
	}
	return highestI, highest
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
