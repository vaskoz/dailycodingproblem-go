package day251

// RadixSortInt32 sorts int32's using radix sort.
// Runs in linear time. O(w*n)
func RadixSortInt32(nums []int32) {
	var max int32
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	exp := int32(1)
	for max/exp > 0 {
		CountingSort(nums, exp)
		exp *= 10
	}
}

// CountingSort is a counting sort for int32's.
func CountingSort(nums []int32, exp int32) {
	counts := make([][]int32, 10)
	for _, v := range nums {
		pos := (v / exp) % 10
		counts[pos] = append(counts[pos], v)
	}
	start := 0
	for _, c := range counts {
		end := start + len(c)
		copy(nums[start:end], c)
		start = end
	}
}
