package day58

// IndexSortedRotated returns the index of a target in a sorted slice.
// Runs in O(log N) time and O(1) space since the input slice is sorted.
// Returns the index or -1 if it doesn't exist.
func IndexSortedRotated(sorted []int, target int) int {
	var low, mid, high int = 0, len(sorted) / 2, len(sorted)
	for low < high {
		if curr := sorted[mid]; curr == target {
			return mid
		} else if inRange(sorted[mid:high], target) {
			low = mid
			mid = (low + high) / 2
		} else {
			high = mid
			mid = (low + high) / 2
		}
	}
	return -1
}

func inRange(sorted []int, target int) bool {
	if sorted[0] < sorted[len(sorted)-1] &&
		target >= sorted[0] && target <= sorted[len(sorted)-1] {
		return true
	} else if sorted[0] > sorted[len(sorted)-1] &&
		(target <= sorted[len(sorted)-1] || target >= sorted[0]) {
		return true
	}
	return false
}
