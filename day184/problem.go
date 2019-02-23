package day184

// GcdSlice returns the greatest common denominator
// of all the values in the slice.
func GcdSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = gcd(nums[i], result)
	}
	return result
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
