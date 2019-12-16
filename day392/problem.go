package day392

// IslandPerimeter finds the perimeter of exactly one island
// in the input.
// Connected land is horizontal or vertical only.
func IslandPerimeter(grid [][]int) int {
	perimeter := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				continue
			}
			// check left
			if nr, nc := r, c-1; nc < 0 || grid[nr][nc] == 0 {
				perimeter++
			}

			// check right
			if nr, nc := r, c+1; nc >= len(grid[nr]) || grid[nr][nc] == 0 {
				perimeter++
			}

			// check up
			if nr, nc := r-1, c; nr < 0 || grid[nr][nc] == 0 {
				perimeter++
			}

			// check down
			if nr, nc := r+1, c; nr >= len(grid) || grid[nr][nc] == 0 {
				perimeter++
			}
		}
	}

	return perimeter
}
