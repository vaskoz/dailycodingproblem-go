package day221

// SevenishNumber returns the n-th seven-ish number.
// a "sevenish" number is one which is either a power of 7,
// or the sum of unique powers of 7
func SevenishNumber(n int) int {
	if n < 1 {
		return 0
	}
	sevens := 1
	index := 1
	nums := []int{1}
	for index < n {
		sevens *= 7
		pos := len(nums)
		nums = append(nums, sevens)
		index++
		for i := 0; i < pos; i++ {
			if index >= n {
				break
			}
			nums = append(nums, nums[i]+sevens)
			index++
		}
	}
	return nums[n-1]
}
