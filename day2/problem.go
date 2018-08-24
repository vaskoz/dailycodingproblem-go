package day2

// ProductsExceptI returns an array with the products of all numbers
// except the value at the i-th position.
// O(N) time complexity.
// O(1) space complexity.
func ProductsExceptI(arr []int) []int {
	result := make([]int, len(arr))
	total := 1
	for _, v := range arr {
		total *= v
	}
	for i, v := range arr {
		result[i] = total / v
	}
	return result
}

// ProductsExceptINoDivision returns an array with the products of all numbers
// except the value at the i-th position. Without using division.
// O(N^2) time complexity.
// O(1) space complexity.
func ProductsExceptINoDivision(arr []int) []int {
	result := make([]int, len(arr))
	for i := range arr {
		total := 1
		for j, w := range arr {
			if i != j {
				total *= w
			}
		}
		result[i] = total
	}
	return result
}
