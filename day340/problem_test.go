package day340

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	pts      []Point
	expected []Point
}{
	{[]Point{}, nil},
	{[]Point{{5, 10}, {20, 100}}, []Point{{5, 10}, {20, 100}}},
	{[]Point{{1, 1}, {-1, -1}, {3, 4}, {6, 1}, {-1, -6}, {-4, -3}}, []Point{{1, 1}, {-1, -1}}},
}

func TestTwoClosestPointsBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := TwoClosestPointsBrute(tc.pts); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkTwoClosestPointsBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TwoClosestPointsBrute(tc.pts)
		}
	}
}
