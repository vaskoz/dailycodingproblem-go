package day230

import "testing"

var testcases = []struct {
	eggs, floors int
	minimum      int
}{
	{2, 10, 4},
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
