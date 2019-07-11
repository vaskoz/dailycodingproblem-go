package day323

import (
	"math/rand"
	"sort"
	"time"
)

// global but unexported for testing.
// nolint
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// ApproximateMedian returns a value approximately close to
// the real median.
// Works by randomly sampling half (N/2) the input.
// Runs in O(N/2) time and O(N/2) extra space.
func ApproximateMedian(unordered []int) int {
	size := len(unordered)
	sampleSize := int(0.5 * float64(size))
	picks := make(map[int]struct{}, sampleSize)
	for len(picks) < sampleSize {
		picks[r.Intn(size)] = struct{}{}
	}
	nums := make([]int, 0, len(picks))
	for pick := range picks {
		nums = append(nums, unordered[pick])
	}
	sort.Ints(nums)
	mid := len(nums) / 2
	return (nums[mid] + nums[mid+1]) / 2
}
