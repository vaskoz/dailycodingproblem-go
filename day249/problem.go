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

// MaxXorPairSort find the maximum XOR of any two elements.
// Runs in O(N log N) time and O(N) space.
func MaxXorPairSort(nums []int) int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	max := -int(^uint(0)>>1) - 1
	for i := 1; i < len(nums); i++ {
		if result := nums[i] ^ nums[i-1]; result > max {
			max = result
		}
	}
	return max
}
