package day21

import "sort"

// Lecture represents a class that needs a room with a start time and end time.
type Lecture struct {
	start, end int
}

// lectureEvent exists for sorting purposes.
type lectureEvent struct {
	time     int
	starting bool
}

// MinimumRooms calculates the minimum number of rooms to accommodate
// overlapping schedules.
// Runtime is O(N log N) with O(2N) space
func MinimumRooms(lectures []Lecture) int {
	events := make([]lectureEvent, 0, 2*len(lectures))
	for i := range lectures {
		events = append(events, lectureEvent{lectures[i].start, true})
		events = append(events, lectureEvent{lectures[i].end, false})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time < events[j].time
	})

	var max, active int

	for i := range events {
		if events[i].starting {
			active++
		} else {
			active--
		}

		if max < active {
			max = active
		}
	}

	return max
}
