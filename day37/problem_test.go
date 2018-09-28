package day37

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	input    []int
	expected [][]int
}{
	{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	{[]int{4, 2, 3, 6}, [][]int{{}, {4}, {2}, {4, 2}, {3}, {4, 3}, {2, 3},
		{4, 2, 3}, {6}, {4, 6}, {2, 6}, {4, 2, 6}, {3, 6}, {4, 3, 6}, {2, 3, 6}, {4, 2, 3, 6}}},
}

func TestPowerSet(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		result := PowerSet(tc.input)
		if len(result) != len(tc.expected) {
			t.Errorf("Length of result doesn't match. Expected %v got %v", len(tc.expected), len(result))
			continue
		}
		for i := range result {
			if !reflect.DeepEqual(result[i], tc.expected[i]) {
				t.Errorf("Expected %v got %v", tc.expected[i], result[i])
			}
		}
	}
}

func BenchmarkPowerSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PowerSet(tc.input)
		}
	}
}
