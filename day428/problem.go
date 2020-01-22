package day428

// LowestCostPyramid returns the lowest cost option.
// O(N) time and O(2*N) space.
func LowestCostPyramid(stones []int) ([]int, int) {
	left := make([]int, len(stones))
	right := make([]int, len(stones))

	left[0] = min(stones[0], 1)
	for i := 1; i < len(stones); i++ {
		left[i] = min(stones[i],
			min(left[i-1]+1, i+1))
	}

	right[len(right)-1] = min(stones[len(stones)-1], 1)

	for i := len(right) - 2; i >= 0; i-- {
		right[i] = min(stones[i],
			min(right[i+1]+1, len(right)-i))
	}

	maxI := 0
	totalMaxI := min(left[0], right[0])

	for i := range stones {
		totalI := min(left[i], right[i])
		if totalI > totalMaxI {
			maxI = i
			totalMaxI = totalI
		}
	}

	cost := 0
	height := totalMaxI

	pyramid := make([]int, len(stones))

	for x := maxI; x >= 0; x-- {
		cost += stones[x] - height
		pyramid[x] = height

		if height > 0 {
			height--
		}
	}

	height = totalMaxI - 1

	for x := maxI + 1; x < len(stones); x++ {
		cost += stones[x] - height
		pyramid[x] = height

		if height > 0 {
			height--
		}
	}

	return pyramid, cost
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
