package day302

// DistinctRegions returns the number of separate regions divided
// by the lines.
// Uses Depth First Search to visit everything.
func DistinctRegions(grid []string) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
		for j := range visited[i] {
			if grid[i][j] != ' ' {
				visited[i][j] = true
			}
		}
	}
	var regions int
	for i := range visited {
		for j := range visited[i] {
			if !visited[i][j] {
				markRegion(visited, i, j)
				regions++
			}
		}
	}
	return regions
}

func markRegion(visited [][]bool, r, c int) {
	if r < 0 || r >= len(visited) || c < 0 || c >= len(visited[r]) || visited[r][c] {
		return
	}
	visited[r][c] = true
	markRegion(visited, r-1, c) // up
	markRegion(visited, r+1, c) // down
	markRegion(visited, r, c-1) // left
	markRegion(visited, r, c+1) // right
}
