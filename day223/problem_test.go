package day223

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	root    *BinaryTree
	inorder []interface{}
}{
	{&BinaryTree{
		1,
		&BinaryTree{
			2,
			&BinaryTree{4, nil, nil},
			&BinaryTree{5, nil, nil},
		},
		&BinaryTree{3, nil, nil},
	},
		[]interface{}{4, 2, 5, 1, 3},
	},
}

func TestInorderMorrisTraversal(t *testing.T) {
	for _, tc := range testcases {
		if result := InorderMorrisTraversal(tc.root); !reflect.DeepEqual(tc.inorder, result) {
			t.Errorf("Expected %v, got %v", tc.inorder, result)
		}
	}
}

func BenchmarkInorderMorrisTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			InorderMorrisTraversal(tc.root)
		}
	}
}
