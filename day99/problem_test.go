package day99

import "testing"

var testcases = []struct {
	input  []int
	length int
}{
	{[]int{100, 4, 200, 1, 3, 2}, 4},
	{[]int{100, 4, 101, 103, 1, 102, 3, 104, 2, 105}, 6},
}

func TestLongestConsecutiveSequenceBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		data := append([]int{}, tc.input...)
		if result := LongestConsecutiveSequenceBrute(data); result != tc.length {
			t.Errorf("Expected %v got %v", tc.length, result)
		}
	}
}

func BenchmarkLongestConsecutiveSequenceBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			data := append([]int{}, tc.input...)
			LongestConsecutiveSequenceBrute(data)
		}
	}
}
