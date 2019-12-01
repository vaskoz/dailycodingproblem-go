package day377

import "sort"

// BruteKMedians calculates the median of every k-sized
// window using brute force.
// It runs in O(N*K log K).
func BruteKMedians(nums []float64, k int) []float64 {
	ans := make([]float64, 0, len(nums)-k)
	window := make([]float64, k)

	for i := 0; i <= len(nums)-k; i++ {
		copy(window, nums[i:k+i])
		sort.Float64s(window)

		if k%2 == 0 {
			ans = append(ans, (window[k/2]+window[(k/2)-1])/2)
		} else {
			ans = append(ans, window[k/2])
		}
	}

	return ans
}
