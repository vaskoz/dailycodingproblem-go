package day18

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    []int
	k        int
	expected []int
}{
	{[]int{10, 5, 2, 7, 8, 7}, 3, []int{10, 7, 8, 8}},
}

func TestMaxSubarraysBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxSubarraysBrute(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v bot got %v", tc.expected, result)
		}
	}
}

func BenchmarkMaxSubarraysBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxSubarraysBrute(tc.input, tc.k)
		}
	}
}

func TestMaxSubarraysOnePass(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := MaxSubarraysOnePass(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v bot got %v", tc.expected, result)
		}
	}
}

func BenchmarkMaxSubarraysOnePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			MaxSubarraysOnePass(tc.input, tc.k)
		}
	}
}
