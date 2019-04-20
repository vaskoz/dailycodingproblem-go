package day240

// MinSwapsAdjacentPairs returns the minimum number of swaps
// required to bring all pairs together.
// A pair consists of the same value.
// Invalid input results in a panic.
func MinSwapsAdjacentPairs(pairs []int) int {
	if len(pairs)%2 != 0 {
		panic("must even number of inputs")
	}
	indicies := make(map[int][]int, len(pairs)/2)
	for i, id := range pairs {
		indicies[id] = append(indicies[id], i)
	}
	if len(indicies) != len(pairs)/2 {
		panic("invalid pair data")
	}
	for _, pos := range indicies {
		if len(pos) != 2 {
			panic("all values must be pairs")
		}
	}
	return minSwapsAdjacentPairs(pairs, indicies, 0)
}

func minSwapsAdjacentPairs(pairs []int, indicies map[int][]int, pos int) int {
	if pos == len(pairs) {
		return 0
	}
	firstIndex := pos
	secondIndex := pos + 1
	first := pairs[firstIndex]
	second := pairs[secondIndex]
	if first == second {
		return minSwapsAdjacentPairs(pairs, indicies, pos+2)
	}
	min := int(^uint(0) >> 1)
	// NOTE: It doesn't matter which one you swap due to symmetry.
	swapWithSecond := indicies[first][1]
	pairs[secondIndex], pairs[swapWithSecond] = pairs[swapWithSecond], pairs[secondIndex]
	if minBelow := 1 + minSwapsAdjacentPairs(pairs, indicies, pos+2); min > minBelow {
		min = minBelow
	}
	pairs[secondIndex], pairs[swapWithSecond] = pairs[swapWithSecond], pairs[secondIndex]
	return min
}
