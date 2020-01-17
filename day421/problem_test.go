package day421

import (
	"reflect"
	"sort"
	"testing"
)

// nolint
var testcases = []struct {
	input, sorted []int32
}{
	{[]int32{170, 45, 75, 90, 802, 24, 2, 66}, []int32{2, 24, 45, 66, 75, 90, 170, 802}},
}

func TestRadixSortInt32(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		copied := make([]int32, len(tc.input))
		copy(copied, tc.input)
		RadixSortInt32(copied)

		if !reflect.DeepEqual(copied, tc.sorted) {
			t.Errorf("Expected %v, got %v", tc.sorted, copied)
		}
	}
}

func BenchmarkRadixSortInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := make([]int32, len(tc.input))
			copy(copied, tc.input)
			RadixSortInt32(copied)
		}
	}
}

func BenchmarkStdLibSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := make([]int32, len(tc.input))
			copy(copied, tc.input)
			sort.Slice(copied, func(i, j int) bool {
				return copied[i] < copied[j]
			})
		}
	}
}
