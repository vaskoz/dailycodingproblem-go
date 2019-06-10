package day291

import "testing"

// nolint
var testcases = []struct {
	weights []int
	k       int
	boats   int
}{
	{[]int{100, 200, 150, 80}, 200, 3},
}

func TestBoatsNeededForEvacuation(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if boats := BoatsNeededForEvacuation(tc.weights, tc.k); boats != tc.boats {
			t.Errorf("Expected %v, got %v", tc.boats, boats)
		}
	}
}

func TestBoatsNeededForEvacuationCant(t *testing.T) {
	t.Parallel()
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic for individuals that weigh more than boat can hold")
		}
	}()
	BoatsNeededForEvacuation([]int{200, 250, 300, 301}, 300)
}

func BenchmarkBoatsNeededForEvacuation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BoatsNeededForEvacuation(tc.weights, tc.k)
		}
	}
}
