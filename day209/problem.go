package day209

// LengthLongestCommonSubseqOf3 is the recursive version.
func LengthLongestCommonSubseqOf3(one, two, three string) int {
	var max int
	if len(one) > max {
		max = len(one)
	}
	if len(two) > max {
		max = len(two)
	}
	if len(three) > max {
		max = len(three)
	}
	dp := make([][][]int, max+1)
	for i := range dp {
		dp[i] = make([][]int, max+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, max+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	return lengthLongestCommonSubseqOf3(dp, []rune(one), []rune(two), []rune(three))
}

func lengthLongestCommonSubseqOf3(dp [][][]int, one, two, three []rune) int {
	i, j, k := len(one)-1, len(two)-1, len(three)-1
	if i < 0 || j < 0 || k < 0 {
		return 0
	}
	if dp[i][j][k] != -1 {
		return dp[i][j][k]
	}
	if one[i] == two[j] && two[j] == three[k] {
		dp[i][j][k] = 1 + lengthLongestCommonSubseqOf3(dp, one[:i], two[:j], three[:k])
		return dp[i][j][k]
	}
	dp[i][j][k] = max(
		max(
			lengthLongestCommonSubseqOf3(dp, one[:i], two, three),
			lengthLongestCommonSubseqOf3(dp, one, two[:j], three),
		),
		lengthLongestCommonSubseqOf3(dp, one, two, three[:k]),
	)
	return dp[i][j][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
