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
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25,
		26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
		51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
		76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
		25,
		[]int{25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
			51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
			76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}},
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
