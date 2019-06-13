package day295

// PascalTriangleRow returns a particular row in Pascal's triangle.
// Runs in O(row) time and space.
func PascalTriangleRow(row int) []int {
	result := make([]int, row+1)
	for entry := range result {
		result[entry] = combinations(row, entry)
	}
	return result
}

func combinations(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
