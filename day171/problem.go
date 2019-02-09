package day171

import "sort"

// EnterOrExit represents either entering the building or exiting.
type EnterOrExit string

const (
	// Enter means people are entering the building.
	Enter EnterOrExit = "enter"
	// Exit means people are leaving the building.
	Exit EnterOrExit = "exit"
)

// Movement is an entry in the building's movement slice.
type Movement struct {
	Timestamp uint64
	Count     int
	Type      EnterOrExit
}

// BusiestBuildingTimes returns the timestamp start and end
// when the building has the most people.
// Runs in O(N log N) time because it sorts the input by timestamp.
// Assumes building starts with 0 people and ends with 0 people.
func BusiestBuildingTimes(moves []Movement) (uint64, uint64) {
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].Timestamp < moves[j].Timestamp
	})
	if len(moves) == 0 {
		return 0, 0
	}
	var start, end uint64
	var count, maxCount int
	for i, move := range moves {
		if move.Type == Enter {
			count += move.Count
			if count > maxCount {
				start = move.Timestamp
				end = moves[i+1].Timestamp
			}
		} else if move.Type == Exit {
			count -= move.Count
		}
	}
	return start, end
}
