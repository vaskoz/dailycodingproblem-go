package day312

import "testing"

// nolint
var testcases = []struct {
	n          int
	tilingWays int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{5, 8},
}

func TestTilingBoardWaysFaster(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if ways := TilingBoardWaysFaster(tc.n); ways != tc.tilingWays {
			t.Errorf("Expected %v, got %v", tc.tilingWays, ways)
		}
	}
}

func TestTilingBoardWaysFasterBadInput(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for a value less than 1")
		}
	}()
	TilingBoardWaysFaster(0)
}

func BenchmarkTilingBoardWaysFaster(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TilingBoardWaysFaster(tc.n)
		}
	}
}

func TestTilingBoardWays(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if ways := TilingBoardWays(tc.n); ways != tc.tilingWays {
			t.Errorf("Expected %v, got %v", tc.tilingWays, ways)
		}
	}
}

func TestTilingBoardWaysBadInput(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for a value less than 1")
		}
	}()
	TilingBoardWays(0)
}

func BenchmarkTilingBoardWays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TilingBoardWays(tc.n)
		}
	}
}
