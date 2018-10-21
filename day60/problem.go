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

// PartitionSet returns if partitioning into 2-sets with equal sums is possible.
// Runs in O(K*N) time using O(K*N) additional space.
// K is the sum of the elements in the input and N is the size of the input.
func PartitionSet(multiset []int) bool {
	sum := 0
	for _, num := range multiset {
		sum += num
	}
	p := make([][]bool, (sum/2)+1)
	for i := range p {
		p[i] = make([]bool, len(multiset)+1)
	}
	for i := 0; i < len(multiset)+1; i++ {
		p[0][i] = true
	}
	for i := 1; i <= sum/2; i++ {
		for j := 1; j <= len(multiset); j++ {
			if i-multiset[j-1] >= 0 {
				p[i][j] = p[i][j-1] || p[i-multiset[j-1]][j-1]
			} else {
				p[i][j] = p[i][j-1]
			}
		}
	}
	return p[sum/2][len(multiset)]
}
