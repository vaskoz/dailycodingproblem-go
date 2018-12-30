package day122

// MaxCoinsPath finds the maximum possible coins collected by going from upper-left
// to lower-right. This is a backtracking algorithm.
func MaxCoinsPath(grid [][]int) int {
	return maxCoinsPathRecursive(grid, 0, 0)
}

func maxCoinsPathRecursive(grid [][]int, row, col int) int {
	if row == len(grid)-1 && col == len(grid[0])-1 {
		lastRow := grid[len(grid)-1]
		return lastRow[len(lastRow)-1]
	}
	var right, down int
	if col+1 < len(grid[0]) {
		right = maxCoinsPathRecursive(grid, row, col+1)
	}
	if row+1 < len(grid) {
		down = maxCoinsPathRecursive(grid, row+1, col)
	}
	if right > down {
		return right + grid[row][col]
	}
	return down + grid[row][col]
}
