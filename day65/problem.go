package day65

import (
	"strconv"
	"strings"
)

// ClockwiseSpiral returns a string that is a representation of the matrix
// as a clockwise spiral starting from the top-left [0][0] coordinate.
func ClockwiseSpiral(matrix [][]int) string {
	if len(matrix) == 0 {
		return ""
	}
	spiral := make([]string, 0, len(matrix)*len(matrix[0]))
	var r, c int
	var left, top, right, bottom int = 0, 0, len(matrix[0]), len(matrix)
	for left < right && top < bottom {
		// go right
		for c < right {
			spiral = append(spiral, strconv.Itoa(matrix[r][c]))
			c++
		}
		c--
		r++
		top++
		for r < bottom {
			spiral = append(spiral, strconv.Itoa(matrix[r][c]))
			r++
		}
		r--
		c--
		right--
		for c >= left {
			spiral = append(spiral, strconv.Itoa(matrix[r][c]))
			c--
		}
		c++
		r--
		bottom--
		for r >= top {
			spiral = append(spiral, strconv.Itoa(matrix[r][c]))
			r--
		}
		r++
		c++
		left++
	}
	return strings.Join(spiral, " ")
}
