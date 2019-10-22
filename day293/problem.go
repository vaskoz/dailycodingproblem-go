package day293

// LowestCostPyramid returns the lowest cost option.
// O(N) time and O(N) space.
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

	total := make([]int, len(stones))
	for i := range total {
		total[i] = min(left[i], right[i])
	}

	maxI := 0

	for i := range total {
		if total[i] > total[maxI] {
			maxI = i
		}
	}

	cost := 0
	height := total[maxI]

	pyramid := make([]int, len(stones))

	for x := maxI; x >= 0; x-- {
		cost += stones[x] - height
		pyramid[x] = height

		if height > 0 {
			height--
		}
	}

	height = total[maxI] - 1

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
