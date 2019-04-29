package day249

// MaxXorPairBrute find the maximum XOR of any two elements.
// Runs in O(N^2) time and O(1) space.
func MaxXorPairBrute(nums []int) int {
	max := -int(^uint(0)>>1) - 1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if result := nums[i] ^ nums[j]; result > max {
				max = result
			}
		}
	}
	return max
}
