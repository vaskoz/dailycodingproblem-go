package day126

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	original, rotated []int
	k                 int
}{
	{[]int{1, 2, 3, 4, 5, 6}, []int{3, 4, 5, 6, 1, 2}, 2},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 1, 2, 3}, 3},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{7, 8, 9, 10, 1, 2, 3, 4, 5, 6}, 6},
}

func TestRotateNew(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := RotateNew(tc.original, tc.k); !reflect.DeepEqual(result, tc.rotated) {
			t.Errorf("Expected %v got %v", tc.rotated, result)
		}
	}
}

func BenchmarkRotateNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RotateNew(tc.original, tc.k)
		}
	}
}

func TestRotateKSwaps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		copied := make([]int, len(tc.original))
		copy(copied, tc.original)
		RotateKSwaps(copied, tc.k)
		if !reflect.DeepEqual(copied, tc.rotated) {
			t.Errorf("Expected %v got %v", tc.rotated, copied)
		}
	}
}

func BenchmarkRotateKSwaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := make([]int, len(tc.original))
			copy(copied, tc.original)
			RotateKSwaps(copied, tc.k)
		}
	}
}

func TestRotateJuggleSwaps(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		copied := make([]int, len(tc.original))
		copy(copied, tc.original)
		RotateJuggleSwaps(copied, tc.k)
		if !reflect.DeepEqual(copied, tc.rotated) {
			t.Errorf("Expected %v got %v", tc.rotated, copied)
		}
	}
}

func BenchmarkRotateJuggleSwaps(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := make([]int, len(tc.original))
			copy(copied, tc.original)
			RotateJuggleSwaps(copied, tc.k)
		}
	}
}
