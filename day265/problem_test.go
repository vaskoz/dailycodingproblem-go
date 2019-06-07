package day265

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	locs    []int
	bonuses []int
}{
	{[]int{10, 40, 200, 1000, 60, 30}, []int{1, 2, 3, 4, 2, 1}},
	{[]int{}, nil},
	{nil, nil},
	{[]int{1000}, []int{1}},
	{[]int{10, 40, 200, 1000, 900, 800, 30}, []int{1, 2, 3, 4, 3, 2, 1}},
}

func TestBonuses(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if bonuses := Bonuses(tc.locs); !reflect.DeepEqual(bonuses, tc.bonuses) {
			t.Errorf("Expected %v, got %v", tc.bonuses, bonuses)
		}
	}
}

func BenchmarkBonuses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			Bonuses(tc.locs)
		}
	}
}
