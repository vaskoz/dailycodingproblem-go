package day258

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	root     *IntBinaryTree
	expected []int
}{
	{nil, nil},
	{
		&IntBinaryTree{
			1,
			&IntBinaryTree{
				2,
				&IntBinaryTree{4, nil, nil},
				&IntBinaryTree{5, nil, nil},
			},
			&IntBinaryTree{
				3,
				&IntBinaryTree{6, nil, nil},
				&IntBinaryTree{7, nil, nil},
			},
		},
		[]int{1, 3, 2, 4, 5, 6, 7},
	},
}

func TestBoustrophedonOrder(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := BoustrophedonOrder(tc.root); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkBoustrophedonOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BoustrophedonOrder(tc.root)
		}
	}
}
