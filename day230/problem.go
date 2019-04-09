package day230

// EggDropDP answers the egg drop question using dynamic programming.
// Consumes O(eggs*floors) additional memory.
func EggDropDP(eggs, floors int) int {
	if floors < 0 || eggs < 0 {
		panic("floors or eggs cannot be negative")
	}
	dp := make([][]int, eggs+1)
	for i := range dp {
		dp[i] = make([]int, floors+1)
		dp[i][0] = 1
	}
	for i := 1; i < floors+1; i++ {
		dp[1][i] = i
	}
	for i := 2; i < eggs+1; i++ {
		for j := 2; j < floors+1; j++ {
			dp[i][j] = int(^uint(0) >> 1)
			for k := 1; k < j+1; k++ {
				res := 1 + max(dp[i-1][k-1], dp[i][j-k])
				if res < dp[i][j] {
					dp[i][j] = res
				}
			}
		}
	}
	return dp[eggs][floors]
}

// EggDropRecursive answers the egg drop question
// recursively.
func EggDropRecursive(eggs, floors int) int {
	if floors < 0 || eggs < 0 {
		panic("floors or eggs cannot be negative")
	}
	if floors < 2 || eggs == 1 {
		return floors
	}
	min := int(^uint(0) >> 1)
	for floor := 1; floor <= floors; floor++ {
		result := max(EggDropRecursive(eggs-1, floor-1),
			EggDropRecursive(eggs, floors-floor))
		if result < min {
			min = result
		}
	}
	return min + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
