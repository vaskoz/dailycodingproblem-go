package day424

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	nums, singleOccurrence []int
}{
	{[]int{2, 4, 6, 8, 10, 2, 6, 10}, []int{4, 8}},
	{[]int{}, []int{}},
	{[]int{5}, []int{5}},
}

func TestSingleOccurrenceBrute(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		copied := append([]int{}, tc.nums...)
		if result := SingleOccurrenceBrute(copied); !reflect.DeepEqual(result, tc.singleOccurrence) {
			t.Errorf("Expected %v got %v", tc.singleOccurrence, result)
		}
	}
}

func BenchmarkSingleOccurrenceBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := append([]int{}, tc.nums...)
			SingleOccurrenceBrute(copied)
		}
	}
}

func TestSingleOccurrenceLinear(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := SingleOccurrenceLinear(tc.nums); !reflect.DeepEqual(result, tc.singleOccurrence) {
			t.Errorf("Expected %v got %v", tc.singleOccurrence, result)
		}
	}
}

func BenchmarkSingleOccurrenceLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SingleOccurrenceBrute(tc.nums)
		}
	}
}
