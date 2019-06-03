package day285

import "testing"

// nolint
var testcases = []struct {
	buildingHeights []int
	sunsetViews     int
}{
	{[]int{3, 7, 8, 3, 6, 1}, 3},
	{nil, 0},
	{[]int{}, 0},
}

func TestSunsetViews(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if sunsets := SunsetViews(tc.buildingHeights); sunsets != tc.sunsetViews {
			t.Errorf("Expected %v, got %v", tc.sunsetViews, sunsets)
		}
	}
}

func BenchmarkSunsetViews(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SunsetViews(tc.buildingHeights)
		}
	}
}
