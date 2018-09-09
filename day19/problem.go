package day19

import "math"

// PaintingCosts represents the cost to paint in each of these colors.
type PaintingCosts []int

// Houses represents a list of houses that can be painting a different color for a price.
type Houses []PaintingCosts

// MinPaintingCost uses dynamic programming to calculate the cheapest way to paint houses
// so that no adjacent houses have the same color.
func MinPaintingCost(houses Houses) int {
	for i := 1; i < len(houses); i++ {
		for j := range houses[i] {
			houses[i][j] += minimum(houses[i-1], j)
		}
	}
	return minimum(houses[len(houses)-1], -1)
}

// minimum returns the minimum of all the values excluding the value at index 'exclude'.
func minimum(costs PaintingCosts, exclude int) int {
	var minimum int64 = math.MaxInt64
	for i, cost := range costs {
		if i == exclude {
			continue
		}
		if cost64 := int64(cost); cost64 < minimum {
			minimum = cost64
		}
	}
	return int(minimum)
}
