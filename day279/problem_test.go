package day279

import "testing"

// nolint
var testcases = []struct {
	friends        FriendGraph
	distinctGroups int
}{
	{
		FriendGraph{
			0: {1: struct{}{}, 2: struct{}{}},
			1: {0: struct{}{}, 5: struct{}{}},
			2: {0: struct{}{}},
			3: {6: struct{}{}},
			4: {},
			5: {1: struct{}{}},
			6: {3: struct{}{}},
		},
		3,
	},
}

func TestNumberOfFriendGroups(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := NumberOfFriendGroups(tc.friends); result != tc.distinctGroups {
			t.Errorf("Expected %v, got %v", tc.distinctGroups, result)
		}
	}
}

func BenchmarkNumberOfFriendGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			NumberOfFriendGroups(tc.friends)
		}
	}
}
