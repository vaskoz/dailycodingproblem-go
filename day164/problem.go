package day164

// FindDuplicate returns the one duplicate value between
// 1 to N.
// Assumes N+1 values are provided.
func FindDuplicate(nums []int) int {
	sub := len(nums) - 1
	sum := 0
	for _, n := range nums {
		sum += n
		for sum > 0 && sub > 0 {
			sum -= sub
			sub--
		}
	}
	return sum
}
