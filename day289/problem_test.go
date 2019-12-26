package day289

import "testing"

// nolint
var testcases = []struct {
	heaps     []int
	forcedWin bool
}{
	{[]int{3, 4, 5}, true},
	{[]int{1, 2, 3}, false},
	{[]int{1, 4, 5}, false},
	{[]int{1, 6, 7}, false},
	{[]int{5, 9, 12}, false},
	{[]int{1, 2, 4, 7}, false},
	{[]int{1, 2, 4, 6}, true},
}

func TestCanFirstPlayerForceWinNimGame(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := CanFirstPlayerForceWinNimGame(tc.heaps); res != tc.forcedWin {
			t.Errorf("Expected %v, got %v", tc.forcedWin, res)
		}
	}
}

func BenchmarkCanFirstPlayerForceWinNimGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			CanFirstPlayerForceWinNimGame(tc.heaps)
		}
	}
}
