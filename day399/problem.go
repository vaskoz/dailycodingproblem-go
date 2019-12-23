package day399

import (
	"errors"
)

var errNotPossible = errors.New("impossible to get 3 partitions equal")

// PartitionIntoThreeEqualSum tries to partition the input
// into 3 contiguous sets with equal sums.
func PartitionIntoThreeEqualSum(nums []int) ([][]int, error) {
	var sum int

	for _, num := range nums {
		sum += num
	}

	target := sum / 3

	if sum%3 != 0 {
		return nil, errNotPossible
	}

	ans := make([][]int, 3)
	pos := 0
	sum = 0

	for _, num := range nums {
		sum += num

		switch {
		case sum < target:
			ans[pos] = append(ans[pos], num)
		case sum == target:
			ans[pos] = append(ans[pos], num)
			sum = 0
			pos++
		default:
			return nil, errNotPossible
		}
	}

	return ans, nil
}
