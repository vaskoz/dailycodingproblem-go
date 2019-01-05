package day136

// MaxRectangleOfOnesAreaBrute returns the area of the largest rectangle
// that contains only 1's.
// The input is a 2-D matrix of 0's and 1's.
// Runtime is O(Row*Col*MaxArea) because it checks every possible rectangle of 1's.
func MaxRectangleOfOnesAreaBrute(matrix [][]int) int {
	var maxArea int
	for row := range matrix {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 1 {
				area := areaOfOnes(matrix, row, col)
				maxArea = max(area, maxArea)
			}
		}
	}
	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func areaOfOnes(m [][]int, r, c int) int {
	minWidth := len(m)
	var colWidth int
	for i := c; i < len(m[r]) && m[r][i] == 1; i++ {
		colWidth++
		var width int
		for j := r; j < len(m) && m[j][i] == 1; j++ {
			width++
		}
		minWidth = min(width, minWidth)
	}
	return colWidth * minWidth
}
