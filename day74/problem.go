package day74

// MultiplicationTableBrute counts the number of X occurrences
// in a N by N multiplication table.
// Runs in O(N^2) time.
func MultiplicationTableBrute(n, x int) int {
	count := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i*j == x {
				count++
			}
		}
	}
	return count
}

// MultiplicationTableLinear counts the number of X occurrences
// in a N by N multiplication table.
// Runs in O(N) time.
func MultiplicationTableLinear(n, x int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if i > x {
			break
		} else if x%i == 0 && x/i <= n {
			count++
		}
	}
	return count
}
