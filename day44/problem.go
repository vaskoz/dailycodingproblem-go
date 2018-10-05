package day44

// InversionCountBrute returns the inversion count using brute-force.
// Runtime is O(N^2) and O(1) space.
func InversionCountBrute(array []int) int {
	var inversions int
	for i := range array {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				inversions++
			}
		}
	}
	return inversions
}

// InversionCount returns the inversion count.
// Runtime is O(N log N).
func InversionCount(array []int) int {
	copied := make([]int, len(array))
	copy(copied, array)
	return mergeSort(copied, make([]int, len(array)), 0, len(array)-1)
}

func mergeSort(input, temp []int, left, right int) int {
	var count int
	if left < right {
		mid := (left + right) / 2
		count = mergeSort(input, temp, left, mid)
		count += mergeSort(input, temp, mid+1, right)
		count += merge(input, temp, left, mid+1, right)
	}
	return count
}

func merge(input, temp []int, left, mid, right int) int {
	li, ri, resi, count := left, mid, left, 0
	for (li <= mid-1) && (ri <= right) {
		if input[li] <= input[ri] {
			temp[resi] = input[li]
			resi++
			li++
		} else {
			temp[resi] = input[ri]
			resi++
			ri++
			count += (mid - li)
		}
	}
	for li <= mid-1 {
		temp[resi] = input[li]
		resi++
		li++
	}

	for ri <= right {
		temp[resi] = input[ri]
		resi++
		ri++
	}

	for i := left; i <= right; i++ {
		input[i] = temp[i]
	}

	return count
}
