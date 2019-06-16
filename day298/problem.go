package day298

// LongestAppleRun returns the length of the longest run of two types of apples.
// Runs in O(N) time.
func LongestAppleRun(apples []int) int {
	lastIndex := make(map[int]int, 2)
	var answer int
	start := 0
	for index, appleType := range apples {
		lastIndex[appleType] = index
		if len(lastIndex) > 2 {
			if answer < index-start {
				answer = index - start
			}
			var earliest int
			min := int(^uint(0) >> 1)
			for t, i := range lastIndex {
				if i < min {
					min = i
					earliest = t
				}
			}
			delete(lastIndex, earliest)
			start = min + 1
		}
	}
	if answer < len(apples)-start {
		answer = len(apples) - start
	}
	return answer
}
