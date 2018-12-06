package day106

// CanHopToEnd can you hop from 0 to the end.
func CanHopToEnd(hops []int) bool {
	start, end := 0, len(hops)-1
	for start < end {
		if hops[start] == 0 {
			break
		}
		start += hops[start]
	}
	return start == end
}
