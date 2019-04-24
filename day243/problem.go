package day243

// MinimizePartitionMaximumSum returns the smallest max sum of k-partitions.
func MinimizePartitionMaximumSum(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	return minimizePartitionMaximumSum(nums, k)
}

func minimizePartitionMaximumSum(nums []int, k int) int {
	if k == 1 {
		var sum int
		for _, v := range nums {
			sum += v
		}
		return sum
	}
	minimumMax := int(^uint(0) >> 1)
	for i := 1; i <= len(nums)-k+1; i++ {
		var partitionSumMax int
		for j := 0; j < i; j++ {
			partitionSumMax += nums[j]
		}
		if minMax := minimizePartitionMaximumSum(nums[i:], k-1); minMax > partitionSumMax {
			partitionSumMax = minMax
		}
		if partitionSumMax < minimumMax {
			minimumMax = partitionSumMax
		}
	}
	return minimumMax
}
