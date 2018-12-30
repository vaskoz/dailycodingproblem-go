package day128

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	disks int
	moves []string
}{
	{3, []string{
		"Move 1 to 3",
		"Move 1 to 2",
		"Move 3 to 2",
		"Move 1 to 3",
		"Move 2 to 1",
		"Move 2 to 3",
		"Move 1 to 3",
	}},
	{4, []string{
		"Move 1 to 2",
		"Move 1 to 3",
		"Move 2 to 3",
		"Move 1 to 2",
		"Move 3 to 1",
		"Move 3 to 2",
		"Move 1 to 2",
		"Move 1 to 3",
		"Move 2 to 3",
		"Move 2 to 1",
		"Move 3 to 1",
		"Move 2 to 3",
		"Move 1 to 2",
		"Move 1 to 3",
		"Move 2 to 3",
	}},
}

func TestTowerOfHanoiMoves(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if moves := TowerOfHanoiMoves(tc.disks); !reflect.DeepEqual(moves, tc.moves) {
			t.Errorf("Expected %v got %v", tc.moves, moves)
		}
	}
}

func BenchmarkTowerOfHanoiMoves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TowerOfHanoiMoves(tc.disks)
		}
	}
}
