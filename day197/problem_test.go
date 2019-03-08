package day197

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	nums     []int
	k        int
	expected []int
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4}},
}

func TestRotateRightK(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		result := append([]int{}, tc.nums...)
		RotateRightK(result, tc.k)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkRotateRightK(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			result := append([]int{}, tc.nums...)
			RotateRightK(result, tc.k)
		}
	}
}
