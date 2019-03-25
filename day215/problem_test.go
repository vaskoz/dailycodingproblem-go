package day215

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	tree     *BinaryTree
	expected []interface{}
}{
	{&BinaryTree{
		5,
		&BinaryTree{
			3,
			&BinaryTree{
				1,
				&BinaryTree{0, nil, nil},
				nil,
			},
			&BinaryTree{4, nil, nil},
		},
		&BinaryTree{
			7,
			&BinaryTree{6, nil, nil},
			&BinaryTree{
				9,
				&BinaryTree{8, nil, nil},
				nil,
			},
		},
	},
		[]interface{}{0, 1, 3, 6, 8, 9},
	},
}

func TestBottomView(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := BottomView(tc.tree); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkBottomView(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BottomView(tc.tree)
		}
	}
}
