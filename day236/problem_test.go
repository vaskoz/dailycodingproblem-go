package day236

import "testing"

var testcases = []struct {
	polygon  []Point
	p        Point
	isInside bool
}{
	{[]Point{{0, 0}, {0, 5}, {5, 5}, {5, 0}}, Point{2, 2}, true},
	{[]Point{{0, 0}, {0, 5}, {5, 5}, {5, 0}}, Point{0, 0}, false},
	{[]Point{{0, 0}, {0, 5}, {5, 5}, {5, 0}}, Point{-2, 2}, false},
	{[]Point{{0, 0}, {0, 5}, {5, 5}, {5, 0}}, Point{5, 2}, false},
	{[]Point{{3, 2}, {2, 3}, {3, 4}, {4, 3}}, Point{3, 3}, true},
	{[]Point{{3, 2}, {2, 3}, {3, 4}, {4, 3}}, Point{2, 3}, false},
	{[]Point{{3, 2}, {2, 3}, {3, 4}, {4, 3}}, Point{0, 0}, false},
	{[]Point{{3, 2}, {2, 3}, {3, 4}, {4, 3}}, Point{3, 5}, false},
	{[]Point{{0, 0}, {5, 5}, {10, 0}}, Point{3, 3}, false},
	{[]Point{{0, 0}, {5, 5}, {10, 0}}, Point{6, 4}, false},
	{[]Point{{0, 0}, {5, 5}, {10, 0}}, Point{6, 3}, true},
}

func TestLiesInsidePolygon(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := LiesInsidePolygon(tc.polygon, tc.p); result != tc.isInside {
			t.Errorf("Given %v and p=%v, expected %v, got %v", tc.polygon, tc.p, tc.isInside, result)
		}
	}
}

func BenchmarkLiesInsidePolygon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LiesInsidePolygon(tc.polygon, tc.p)
		}
	}
}
