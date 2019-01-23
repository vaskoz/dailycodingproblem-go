package day150

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	pts      []Point
	center   Point
	k        int
	expected []Point
}{
	{[]Point{{0, 0}, {5, 4}, {3, 1}}, Point{1, 2}, 2, []Point{{0, 0}, {3, 1}}},
	{[]Point{{5, 4}, {3, 1}, {0, 0}}, Point{1, 2}, 2, []Point{{3, 1}, {0, 0}}},
	{[]Point{{5, 4}, {3, 1}, {0, 0}, {10, 10}, {20, 20}, {30, 30}}, Point{1, 2}, 2, []Point{{3, 1}, {0, 0}}},
	{[]Point{{5, 4}, {0, 0}, {10, 10}, {20, 20}, {3, 1}, {30, 30}}, Point{1, 2}, 2, []Point{{0, 0}, {3, 1}}},
}

func TestKNearestPoints(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := KNearestPoints(tc.pts, tc.center, tc.k); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkKNearestPoints(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			KNearestPoints(tc.pts, tc.center, tc.k)
		}
	}
}
