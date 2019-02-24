package day185

import "testing"

var testcases = []struct {
	r1, r2 Rectangle
	area   int
}{
	{Rectangle{1, 4, 3, 3}, Rectangle{0, 5, 4, 3}, 6},
	{Rectangle{10, 10, 3, 3}, Rectangle{0, 0, 3, 3}, 0},
	{Rectangle{10, 10, 3, 3}, Rectangle{10, 0, 3, 3}, 0},
}

func TestAreaOfIntersection(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if area := AreaOfIntersection(tc.r1, tc.r2); area != tc.area {
			t.Errorf("Expected %v, got %v", tc.area, area)
		}
	}
}

func BenchmarkAreaOfIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AreaOfIntersection(tc.r1, tc.r2)
		}
	}
}
