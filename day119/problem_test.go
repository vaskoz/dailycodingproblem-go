package day119

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	intervals, expected []Interval
}{
	{nil, nil},
	{[]Interval{{5, 10}}, []Interval{{5, 5}}},
	{[]Interval{{0, 3}, {2, 6}, {3, 4}, {6, 9}}, []Interval{{3, 6}}},
	{[]Interval{{0, 1}, {2, 4}, {3, 4}, {6, 9}}, []Interval{{0, 0}, {4, 4}, {6, 6}}},
	{[]Interval{{10, 20}, {2, 4}, {3, 4}, {6, 9}}, []Interval{{4, 4}, {6, 6}, {10, 10}}},
	{[]Interval{{10, 20}, {1, 10}, {3, 4}, {20, 30}}, []Interval{{4, 20}}},
}

func TestSmallestSetCoverage(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SmallestSetCoverage(tc.intervals); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkSmallestSetCoverage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SmallestSetCoverage(tc.intervals)
		}
	}
}
