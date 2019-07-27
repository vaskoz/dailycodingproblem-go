package day339

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	nums     []int
	k        int
	expected []int
}{
	{[]int{20, 303, 3, 4, 25}, 49, []int{20, 4, 25}},
	{[]int{20, 303, 3, 4, 25}, 1000, nil},
	{[]int{300}, 900, nil},
	{[]int{300, 301, 300}, 900, nil},
	{[]int{300, 300, 301}, 900, nil},
}

func TestThreeSum(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ThreeSum(tc.nums, tc.k); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkThreeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ThreeSum(tc.nums, tc.k)
		}
	}
}
