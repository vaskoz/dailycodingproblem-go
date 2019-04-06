package day209

// LengthLongestCommonSubseqOf3Iterative is the iterative version.
func LengthLongestCommonSubseqOf3Iterative(one, two, three string) int {
	X := []rune(one)
	Y := []rune(two)
	Z := []rune(three)
	dp := make([][][]int, len(X)+1)
	for i := range dp {
		dp[i] = make([][]int, len(Y)+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, len(Z)+1)
			for k := range dp[i][j] {
				switch {
				case i == 0 || j == 0 || k == 0:
					dp[i][j][k] = 0
				case X[i-1] == Y[j-1] && X[i-1] == Z[k-1]:
					dp[i][j][k] = 1 + dp[i-1][j-1][k-1]
				default:
					dp[i][j][k] = max(
						max(
							dp[i-1][j][k],
							dp[i][j-1][k],
						),
						dp[i][j][k-1],
					)
				}
			}
		}
	}
	return dp[len(X)][len(Y)][len(Z)]
}

// LengthLongestCommonSubseqOf3Recursive is the recursive version.
func LengthLongestCommonSubseqOf3Recursive(one, two, three string) int {
	dp := make([][][]int, len(one)+1)
	for i := range dp {
		dp[i] = make([][]int, len(two)+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, len(three)+1)
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
