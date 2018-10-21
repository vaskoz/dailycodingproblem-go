package day60

// BruteForcePartitionSet returns if partitioning into 2-sets with equal sums is possible.
// Runs in O(2^N) time.
func BruteForcePartitionSet(multiset []int) bool {
	sum := 0
	for _, num := range multiset {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	half := sum / 2
	return subsetSum(multiset, half)
}

func subsetSum(multiset []int, target int) bool {
	for i, v := range multiset {
		if remainder := target - v; remainder > 0 {
			if subsetSum(multiset[i+1:], remainder) {
				return true
			}
		} else if remainder == 0 {
			return true
		}
	}
	return false
}
