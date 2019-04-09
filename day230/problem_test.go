package day230

import "testing"

var testcases = []struct {
	eggs, floors int
	minimum      int
}{
	{1, 5, 5},
	{2, 10, 4},
}

func TestEggDropDP(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := EggDropDP(tc.eggs, tc.floors); result != tc.minimum {
			t.Errorf("Expected %d, got %d", tc.minimum, result)
		}
	}
}

func TestEggDropDPPanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected a panic for negative values")
		}
	}()
	EggDropDP(-1, 0)
}

func BenchmarkEggDropDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EggDropDP(tc.eggs, tc.floors)
		}
	}
}

func TestEggDropRecursive(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := EggDropRecursive(tc.eggs, tc.floors); result != tc.minimum {
			t.Errorf("Expected %d, got %d", tc.minimum, result)
		}
	}
}

func TestEggDropRecursivePanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("expected a panic for negative values")
		}
	}()
	EggDropRecursive(-1, 0)
}

func BenchmarkEggDropRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			EggDropRecursive(tc.eggs, tc.floors)
		}
	}
}
