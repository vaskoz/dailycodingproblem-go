package day151

// PixelGrid is a 2-D grid of pixel colors.
type PixelGrid [][]Color

// Coordinate represents a pixel location in the grid.
type Coordinate struct {
	Row, Col int
}

// Color represents a unique color that a pixel can have.
type Color rune

// ReplaceAdjacentColorPixel changes all the pixels adjacent to the given coordinate
// of the same color to the new color.
func ReplaceAdjacentColorPixel(grid PixelGrid, c Coordinate, color Color) {
	start := grid[c.Row][c.Col]
	replaceAdjacentColorPixel(grid, c, start, color)
}

func replaceAdjacentColorPixel(grid PixelGrid, c Coordinate, start, target Color) {
	if c.Row < 0 || c.Col < 0 || c.Row >= len(grid) || c.Col >= len(grid[0]) {
		return
	} else if grid[c.Row][c.Col] == start {
		grid[c.Row][c.Col] = target
		replaceAdjacentColorPixel(grid, Coordinate{c.Row - 1, c.Col}, start, target) // up
		replaceAdjacentColorPixel(grid, Coordinate{c.Row + 1, c.Col}, start, target) // down
		replaceAdjacentColorPixel(grid, Coordinate{c.Row, c.Col - 1}, start, target) // left
		replaceAdjacentColorPixel(grid, Coordinate{c.Row, c.Col + 1}, start, target) // right
	}
}
