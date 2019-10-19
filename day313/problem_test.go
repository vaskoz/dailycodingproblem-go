package day313

import "testing"

// nolint
var testcases = []struct {
	start, target [3]int
	deadEnds      [][3]int
	expected      int
}{
	{
		[3]int{0, 0, 0},
		[3]int{7, 1, 9},
		[][3]int{{1, 1, 1}},
		5,
	},
	{
		[3]int{0, 0, 0},
		[3]int{0, 0, 0},
		[][3]int{{1, 1, 1}},
		0,
	},
	{
		[3]int{0, 0, 0},
		[3]int{5, 5, 5},
		[][3]int{{1, 0, 0}, {0, 1, 0}, {9, 0, 0}, {0, 9, 0}, {0, 0, 1}, {0, 0, 9}},
		-1,
	},
	{
		[3]int{0, 0, 0},
		[3]int{5, 5, 5},
		[][3]int{{0, 1, 0}, {9, 0, 0}, {0, 9, 0}, {0, 0, 1}, {0, 0, 9}},
		15,
	},
}

func TestMovesToUnlock(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := MovesToUnlock(tc.start, tc.target, tc.deadEnds); result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func TestMovesToUnlockBadStartPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for bad start coordinates")
		}
	}()

	MovesToUnlock([3]int{-1, 10, 5}, [3]int{1, 1, 1}, [][3]int{})
}

func TestMovesToUnlockBadTargetPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for bad target coordinates")
		}
	}()

	MovesToUnlock([3]int{0, 0, 0}, [3]int{10, 1, 1}, [][3]int{})
}

func TestMovesToUnlockBadDeadEndPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for bad dead end coordinates")
		}
	}()

	MovesToUnlock([3]int{0, 0, 0}, [3]int{1, 1, 1}, [][3]int{{-1, 0, 0}})
}

func BenchmarkMovesToUnlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MovesToUnlock(tc.start, tc.target, tc.deadEnds)
		}
	}
}
