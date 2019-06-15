package day296

import "testing"

// nolint
var testcases = []struct {
	sorted   []int
	expected *BST
}{
	{
		[]int{1, 2, 3, 4, 5, 6, 7},
		&BST{
			4,
			&BST{
				2,
				&BST{Value: 1},
				&BST{Value: 3},
			},
			&BST{
				6,
				&BST{Value: 5},
				&BST{Value: 7},
			},
		},
	},
}

func TestSortedSliceToBST(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := SortedSliceToBST(tc.sorted); !equal(tc.expected, result) {
			t.Errorf("BSTs don't match in TCID%d", tcid)
		}
	}
}

func BenchmarkSortedSliceToBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SortedSliceToBST(tc.sorted)
		}
	}
}

func equal(a, b *BST) bool {
	switch {
	case a == nil && b != nil:
		return false
	case a != nil && b == nil:
		return false
	case a == nil && b == nil:
		return true
	case a.Value != b.Value:
		return false
	}
	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
