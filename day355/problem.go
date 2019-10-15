package day355

import (
	"sort"
)

type index struct {
	pos int
	val float64
}

// FindY returns the value of Y that meets the requirements
// of the problem.
// Runs in O(N log N) time and O(N) space due to sorting
// the data by fractional value.
func FindY(x []float64) []int {
	var sumX float64

	var sumY int

	data := make([]index, len(x))

	for pos, val := range x {
		data[pos] = index{pos, val}
		sumX += val
		sumY += int(val)
	}

	// sort by smallest fractional value.
	sort.Slice(data, func(i, j int) bool {
		di := data[i].val - float64(int(data[i].val))
		dj := data[j].val - float64(int(data[j].val))
		return di < dj
	})

	yLow := make([]int, len(x))
	sumXlow := int(sumX)
	yLowOnes := sumXlow - sumY
	yHighOnes := yLowOnes + 1

	var yLowDiff float64

	for i := 0; i < len(data)-yLowOnes; i++ {
		pos := data[i].pos
		yLow[pos] = int(data[i].val)
		yLowDiff += data[i].val - float64(yLow[pos])
	}

	for i := len(data) - yLowOnes; i < len(data); i++ {
		pos := data[i].pos
		yLow[pos] = int(data[i].val) + 1
		yLowDiff += float64(yLow[pos]) - data[i].val
	}

	var yHighDiff float64

	yHigh := make([]int, len(x))

	for i := 0; i < len(data)-yHighOnes; i++ {
		pos := data[i].pos
		yHigh[pos] = int(data[i].val)
		yHighDiff += (data[i].val - float64(yHigh[pos]))
	}

	for i := len(data) - yHighOnes; i < len(data); i++ {
		pos := data[i].pos
		yHigh[pos] = int(data[i].val) + 1
		yHighDiff += (float64(yHigh[pos]) - data[i].val)
	}

	if yLowDiff < yHighDiff {
		return yLow
	}

	return yHigh
}
