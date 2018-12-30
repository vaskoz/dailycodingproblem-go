package day126

// RotateNew returns a new slice that has been rotated by k.
func RotateNew(nums []int, k int) []int {
	result := make([]int, len(nums))
	for i := k; i < k+len(nums); i++ {
		result[i-k] = nums[i%len(nums)]
	}
	return result
}

// RotateKSwaps rotates the values by k in-place using swaps.
// It uses O(K*N) time.
func RotateKSwaps(nums []int, k int) {
	for i := 0; i < k; i++ {
		for j := 1; j < len(nums); j++ {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

// RotateJuggleSwaps rotates using swaps but faster.
func RotateJuggleSwaps(nums []int, k int) {
	for i := 0; i < gcd(len(nums), k); i++ {
		tmp := nums[i]
		j := i
		for {
			jj := j + k
			if jj >= len(nums) {
				jj = jj % len(nums)
			}
			if jj == i {
				break
			}
			nums[j] = nums[jj]
			j = jj
		}
		nums[j] = tmp
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
