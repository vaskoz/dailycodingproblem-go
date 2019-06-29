package day311

import "testing"

// nolint
var testcases = []struct {
	nums      []int
	peakIndex int
}{
	{[]int{1, 3, 20, 4, 1, 0}, 2},
	{nil, -1},
	{[]int{100}, 0},
	{[]int{5, 3, 2, 1, 0}, 0},
	{[]int{1, 3, 5, 8, 10, 100}, 5},
}

func TestPeakIndexNR(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if peak := PeakIndexNR(tc.nums); peak != tc.peakIndex {
			t.Errorf("Expected %v, got %v", tc.peakIndex, peak)
		}
	}
}

func BenchmarkPeakIndexNR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PeakIndex(tc.nums)
		}
	}
}

func TestPeakIndex(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if peak := PeakIndex(tc.nums); peak != tc.peakIndex {
			t.Errorf("Expected %v, got %v", tc.peakIndex, peak)
		}
	}
}

func BenchmarkPeakIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PeakIndex(tc.nums)
		}
	}
}

func TestPeakIndexBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if peak := PeakIndexBrute(tc.nums); peak != tc.peakIndex {
			t.Errorf("Expected %v, got %v", tc.peakIndex, peak)
		}
	}
}

func BenchmarkPeakIndexBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PeakIndexBrute(tc.nums)
		}
	}
}
