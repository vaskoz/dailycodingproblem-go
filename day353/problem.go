package day353

// LargestRectangleHistogram returns the area of the
// largest possible rectangle given the histogram.
// Runs in O(N) time with O(N) space for the stack.
func LargestRectangleHistogram(hist []int) int {
	stack := make([]int, 1, len(hist))
	stack[0] = -1
	var result int
	for i := range hist {
		for stack[len(stack)-1] != -1 && hist[stack[len(stack)-1]] >= hist[i] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = max(result, hist[last]*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	for stack[len(stack)-1] != -1 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = max(result, hist[last]*(len(hist)-stack[len(stack)-1]-1))
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
