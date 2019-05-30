package day281

// FewestBricksCut return the fewest number of
// bricks that must be cut to create a vertical line.
func FewestBricksCut(bricks [][]int) int {
	cumm := make(map[int]int)
	for _, row := range bricks {
		last := 0
		for _, brick := range row {
			last += brick
			cumm[last]++
		}
		delete(cumm, last)
	}
	var most int
	for _, count := range cumm {
		if count > most {
			most = count
		}
	}
	return len(bricks) - most
}
