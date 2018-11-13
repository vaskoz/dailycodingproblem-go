package day84

// CountIslands returns the number of islands.
// 1 represents land, and 0 represents water.
// Runs in O(GridSize) and uses O(GridSize) space.
func CountIslands(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}
	islands := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && !visited[i][j] {
				islands++
				exploreIsland(grid, visited, i, j)
			}
		}
	}
	return islands
}

// exploreIsland is a DFS of the island
func exploreIsland(grid [][]int, visited [][]bool, i, j int) {
	visited[i][j] = true
	if up := i - 1; up >= 0 && isUnexplored(grid[up][j], visited[up][j]) {
		exploreIsland(grid, visited, up, j)
	}
	if down := i + 1; down < len(grid) && isUnexplored(grid[down][j], visited[down][j]) {
		exploreIsland(grid, visited, down, j)
	}
	if left := j - 1; left >= 0 && isUnexplored(grid[i][left], visited[i][left]) {
		exploreIsland(grid, visited, i, left)
	}
	if right := j + 1; right < len(grid[i]) && isUnexplored(grid[i][right], visited[i][right]) {
		exploreIsland(grid, visited, i, right)
	}
}

func isUnexplored(gridValue int, visitedStatus bool) bool {
	return gridValue == 1 && !visitedStatus
}
