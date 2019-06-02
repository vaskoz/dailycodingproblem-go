package day283

// RegularNumbersBrute returns n regular numbers
// using brute force.
func RegularNumbersBrute(n int) []int {
	if n < 0 {
		panic("negative values are not possible")
	}
	result := make([]int, 0, n)
	count := 0
	for x := 1; count < n; x++ {
		remain := repeatDivide(x, 2)
		remain = repeatDivide(remain, 3)
		remain = repeatDivide(remain, 5)
		if remain == 1 {
			result = append(result, x)
			count++
		}
	}
	return result
}

func repeatDivide(val, factor int) int {
	for val%factor == 0 {
		val /= factor
	}
	return val
}
