package day21

import "testing"

// nolint
var testcases = []struct {
	lectures    []Lecture
	roomsNeeded int
}{
	{[]Lecture{{30, 75}, {0, 50}, {60, 150}}, 2},
	{[]Lecture{{0, 10}, {10, 20}, {20, 30}, {1, 11}, {11, 21}, {21, 31}}, 2},
}

func TestMinimumRooms(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := MinimumRooms(tc.lectures); result != tc.roomsNeeded {
			t.Errorf("Expected %v but got %v", tc.roomsNeeded, result)
		}
	}
}

func BenchmarkMinimumRooms(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MinimumRooms(tc.lectures)
		}
	}
}
