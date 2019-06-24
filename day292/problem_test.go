package day292

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	enemies          map[int]map[int]struct{}
	teamOne, teamTwo []int
	err              error
}{
	{
		map[int]map[int]struct{}{
			0: {3: struct{}{}},
			1: {2: struct{}{}},
			2: {
				1: struct{}{},
				4: struct{}{},
			},
			3: {
				0: struct{}{},
				4: struct{}{},
				5: struct{}{},
			},
			4: {
				2: struct{}{},
				3: struct{}{},
			},
			5: {3: struct{}{}},
		},
		[]int{0, 1, 4, 5},
		[]int{2, 3},
		nil,
	},
	{
		map[int]map[int]struct{}{
			0: {3: struct{}{}},
			1: {2: struct{}{}},
			2: {
				1: struct{}{},
				4: struct{}{},
			},
			3: {
				0: struct{}{},
				2: struct{}{},
				4: struct{}{},
				5: struct{}{},
			},
			4: {
				2: struct{}{},
				3: struct{}{},
			},
			5: {3: struct{}{}},
		},
		nil,
		nil,
		ErrTeamNotPossible(),
	},
	{
		map[int]map[int]struct{}{
			1: {0: struct{}{}},
			0: {
				1: struct{}{},
				2: struct{}{},
				3: struct{}{},
			},
			2: {0: struct{}{}},
			3: {0: struct{}{}},
		},
		[]int{0},
		[]int{1, 2, 3},
		nil,
	},
}

func TestAssignTeams(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if t1, t2, err := AssignTeams(tc.enemies); !reflect.DeepEqual(t1, tc.teamOne) ||
			!reflect.DeepEqual(t2, tc.teamTwo) || err != tc.err {
			t.Errorf("Expected (%v,%v,%v), got (%v,%v,%v)", tc.teamOne, tc.teamTwo, tc.err, t1, t2, err)
		}
	}
}

func BenchmarkAssignTeams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AssignTeams(tc.enemies) // nolint
		}
	}
}
