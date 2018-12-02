package day102

import (
	"reflect"
	"testing"
)

var nonNegativeTestcases = []struct {
	input    []int
	k        int
	expected []int
}{
	{[]int{1, 2, 3, 4, 5}, 9, []int{2, 3, 4}},
	{[]int{1, 2, 3, 4, 5}, 16, nil},
	{[]int{1, 2, 3, 4, 5}, 15, []int{1, 2, 3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25},
		49, []int{4, 5, 6, 7, 8, 9, 10}},
}

var negativeTestcases = []struct {
	input    []int
	k        int
	expected []int
}{
	{[]int{-1, 2, 3, 4, 5}, 1, []int{-1, 2}},
	{[]int{1, -2, 3, 4, 5}, 5, []int{-2, 3, 4}},
	{[]int{1, 2, -3, 4, 5}, 6, []int{-3, 4, 5}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, -24, -25},
		-49, []int{-24, -25}},
}

func TestContiguousSumBrute(t *testing.T) {
	t.Parallel()
	for _, tc := range nonNegativeTestcases {
		if result := ContiguousSumBrute(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
	for _, tc := range negativeTestcases {
		if result := ContiguousSumBrute(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkContiguousSumBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range nonNegativeTestcases {
			ContiguousSumBrute(tc.input, tc.k)
		}
		for _, tc := range negativeTestcases {
			ContiguousSumBrute(tc.input, tc.k)
		}
	}
}

func TestContiguousSumNonNegative(t *testing.T) {
	t.Parallel()
	for _, tc := range nonNegativeTestcases {
		if result := ContiguousSumNonNegative(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func TestContiguousSumNonNegativePanic(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic from the negative input value")
		}
	}()
	ContiguousSumNonNegative([]int{1, -2, 3}, 10)
}

func BenchmarkContiguousSumNonNegative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range nonNegativeTestcases {
			ContiguousSumNonNegative(tc.input, tc.k)
		}
		for _, tc := range nonNegativeTestcases {
			ContiguousSumNonNegative(tc.input, tc.k)
		}
	}
}

func TestContiguousSumNegatives(t *testing.T) {
	t.Parallel()
	for _, tc := range nonNegativeTestcases {
		if result := ContiguousSumNegatives(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
	for _, tc := range negativeTestcases {
		if result := ContiguousSumNegatives(tc.input, tc.k); !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkContiguousSumNegatives(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range nonNegativeTestcases {
			ContiguousSumNegatives(tc.input, tc.k)
		}
		for _, tc := range negativeTestcases {
			ContiguousSumNegatives(tc.input, tc.k)
		}
	}
}
