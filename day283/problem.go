package day283

// RegularNumbersFaster returns n regular numbers
// using brute force.
func RegularNumbersFaster(n int) []int {
	if n < 0 {
		panic("negative values are not possible")
	} else if n == 0 {
		return []int{}
	}
	result := make([]int, n)
	result[0] = 1
	var i2, i3, i5 int
	next2, next3, next5 := 2, 3, 5
	for i := 1; i < n; i++ {
		result[i] = min(next2, next3, next5)
		if result[i] == next2 {
			i2++
			next2 = result[i2] * 2
		}
		if result[i] == next3 {
			i3++
			next3 = result[i3] * 3
		}
		if result[i] == next5 {
			i5++
			next5 = result[i5] * 5
		}
	}
	return result
}

func min(vals ...int) int {
	smallest := vals[0]
	for _, v := range vals {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

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
