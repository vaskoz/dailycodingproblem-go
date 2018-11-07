package day77

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input, expected []Interval
}{
	{nil, nil},
	{[]Interval{}, nil},
	{[]Interval{
		{1, 3}, {5, 8}, {4, 10}, {20, 25}},
		[]Interval{{1, 3}, {4, 10}, {20, 25}},
	},
	{[]Interval{
		{1, 3}, {1, 8}, {4, 10}, {20, 25}},
		[]Interval{{1, 10}, {20, 25}},
	},
}

func TestMergeOverlappingIntervals(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := MergeOverlappingIntervals(tc.input); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("TC%d Expected %v got %v", tcid, tc.expected, result)
		}
	}
}

func BenchmarkMergeOverlappingIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MergeOverlappingIntervals(tc.input)
		}
	}
}
