package day399

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input       []int
	expected    [][]int
	expectedErr error
}{
	{
		[]int{3, 5, 8, 0, 8},
		[][]int{
			{3, 5},
			{8},
			{0, 8},
		},
		nil,
	},
	{
		[]int{3, 5, 8, 0, 9},
		nil,
		errNotPossible,
	},
	{
		[]int{5, 5, 8, 0, 9},
		nil,
		errNotPossible,
	},
}

func TestPartitionIntoThreeEqualSum(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res, err := PartitionIntoThreeEqualSum(tc.input); !reflect.DeepEqual(res, tc.expected) ||
			tc.expectedErr != err {
			t.Errorf("Expected(%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, res, err)
		}
	}
}

func BenchmarkPartitionIntoThreeEqualSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PartitionIntoThreeEqualSum(tc.input) // nolint
		}
	}
}
