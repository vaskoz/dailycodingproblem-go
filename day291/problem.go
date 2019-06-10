package day291

import "sort"

// BoatsNeededForEvacuation returns the minimum number of boats needed to
// evacuate Codeville before the hurricane.
// peopleWeights is a list of the people needing evacuation and k is the
// max weight allowed on each boat.
func BoatsNeededForEvacuation(peopleWeights []int, k int) int {
	weights := make([]int, len(peopleWeights))
	copy(weights, peopleWeights)
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))
	var boats int
	for i, weight := range weights {
		if weight != -1 {
			if weight > k {
				panic("individual is over boat's weight. can't evacuate.")
			}
			remaining := k - weight
			weights[i] = -1
			for j := i + 1; j < len(weights); j++ {
				if weights[j] != -1 && weights[j] <= remaining {
					weights[j] = -1
					break
				}
			}
			boats++
		}
	}
	return boats
}
