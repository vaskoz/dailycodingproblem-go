package day230

// EggDropRecursive answers the egg drop question
// recursively.
func EggDropRecursive(eggs, floors int) int {
	if floors < 0 || eggs < 0 {
		panic("floors or eggs cannot be negative")
	}
	if floors < 2 || eggs == 1 {
		return floors
	}
	min := int(^uint(0) >> 1)
	for floor := 1; floor <= floors; floor++ {
		result := max(EggDropRecursive(eggs-1, floor-1),
			EggDropRecursive(eggs, floors-floor))
		if result < min {
			min = result
		}
	}
	return min + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
