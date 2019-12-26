package day336

// PossibleMaxHeaps returns a count of all possible
// max heaps.
func PossibleMaxHeaps(n int) int {
	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = -1
	}

	nck := make([][]int, n+1)

	for i := range nck {
		nck[i] = make([]int, n+1)
		for j := range nck[i] {
			nck[i][j] = -1
		}
	}

	currLog2 := -1
	currPower2 := 1
	log2 := make([]int, n+1)

	for i := 1; i <= n; i++ {
		if currPower2 == i {
			currLog2++

			currPower2 *= 2
		}

		log2[i] = currLog2
	}

	return numberOfHeaps(n, dp, log2, nck)
}

func numberOfHeaps(n int, dp, log2 []int, nck [][]int) int {
	if n <= 1 {
		return 1
	}

	if dp[n] != -1 {
		return dp[n]
	}

	left := getLeft(n, log2)
	ans := choose(n-1, left, nck) *
		numberOfHeaps(left, dp, log2, nck) *
		numberOfHeaps(n-1-left, dp, log2, nck)
	dp[n] = ans

	return ans
}

func getLeft(n int, log2 []int) int {
	if n == 1 {
		return 0
	}

	h := log2[n]
	numh := 1 << h
	last := n - ((1 << h) - 1)

	if last >= numh/2 {
		return (1 << h) - 1
	}

	return (1 << h) - 1 - ((numh / 2) - last)
}

func choose(n, k int, nck [][]int) int {
	if k > n {
		return 0
	}

	if n <= 1 {
		return 1
	}

	if k == 0 {
		return 1
	}

	if nck[n][k] != -1 {
		return nck[n][k]
	}

	answer := choose(n-1, k-1, nck) + choose(n-1, k, nck)
	nck[n][k] = answer

	return answer
}
