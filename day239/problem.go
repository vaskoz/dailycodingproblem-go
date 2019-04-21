package day239

// ValidUnlockKeypadNumber returns the number of possible
// valid combinations.
func ValidUnlockKeypadNumber() int {
	var jumps [10][10]int
	jumps[1][3], jumps[3][1] = 2, 2
	jumps[7][9], jumps[9][7] = 8, 8
	jumps[1][7], jumps[7][1] = 4, 4
	jumps[3][9], jumps[9][3] = 6, 6
	jumps[1][9], jumps[9][1], jumps[2][8], jumps[8][2],
		jumps[3][7], jumps[7][3], jumps[4][6], jumps[6][4] = 5, 5, 5, 5, 5, 5, 5, 5
	var visited [10]bool
	ways := 0
	for i := 1; i < 10; i++ {
		ways += 4 * validUnlockKeypadNumber(visited, jumps, 1, i-1)
	}
	return ways
}

func validUnlockKeypadNumber(visited [10]bool, jumps [10][10]int, cur, to int) int {
	if to == 0 {
		return 1
	}
	ways := 0
	visited[cur] = true
	for i := 1; i < 10; i++ {
		if !visited[i] && (jumps[i][cur] == 0 || visited[jumps[i][cur]]) {
			ways += validUnlockKeypadNumber(visited, jumps, i, to-1)
		}
	}
	visited[cur] = false
	return ways
}
