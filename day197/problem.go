package day197

// RotateRightK rotates the nums slice in-place to the right.
// Runs in O(N) time and O(1) space.
func RotateRightK(nums []int, k int) {
	k %= len(nums)
	for i := 0; i < k/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}
	for i := 0; i < (len(nums)-k)/2; i++ {
		nums[k+i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[k+i]
	}
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
