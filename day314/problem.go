package day314

import "sort"

// MinimumTowerRange returns the minimum range of a tower to
// cover all listeners.
// Assumes listeners and towers are sorted in ascending order.
// NOTE: This sorts these slices, if they're not already.
// Runs in O(N) time and O(1) space.
func MinimumTowerRange(listeners, towers []int) int {
	if !sort.IntsAreSorted(listeners) {
		sort.Ints(listeners)
	}
	if !sort.IntsAreSorted(towers) {
		sort.Ints(towers)
	}
	var minRange int
	towerIndex := 0
	for _, listener := range listeners {
		dist := abs(listener - towers[towerIndex])
		for i := towerIndex; i < len(towers); i++ {
			if d := abs(listener - towers[i]); d > dist {
				break
			} else {
				towerIndex = i
				dist = d
			}
		}
		if dist > minRange {
			minRange = dist
		}
	}
	return minRange
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
