package day339

// ThreeSum returns the first possible answer to the 3-SUM
// problem.
// Runs in O(N^2) time and O(N) space due to a map.
// If no answer exists it returns nil.
func ThreeSum(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			needed := k - nums[i] - nums[j]
			if m[needed] >= 0 {
				m[nums[i]]--
				m[nums[j]]--
				m[needed]--
				if m[nums[i]] >= 0 && m[nums[j]] >= 0 && m[needed] >= 0 {
					return []int{nums[i], nums[j], needed}
				}
				m[nums[i]]++
				m[nums[j]]++
				m[needed]++
			}
		}
	}
	return nil
}
