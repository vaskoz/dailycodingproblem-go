package day179

import "testing"

var testcases = []struct {
	postorder []int
	expected  *BinaryTree
}{
	{[]int{2, 4, 3, 8, 7, 5}, &BinaryTree{
		5,
		&BinaryTree{
			3,
			&BinaryTree{2, nil, nil},
			&BinaryTree{4, nil, nil},
		},
		&BinaryTree{
			7,
			nil,
			&BinaryTree{8, nil, nil},
		},
	}},
	{[]int{1, 7, 5, 50, 40, 10}, &BinaryTree{
		10,
		&BinaryTree{
			5,
			&BinaryTree{1, nil, nil},
			&BinaryTree{7, nil, nil},
		},
		&BinaryTree{
			40,
			nil,
			&BinaryTree{50, nil, nil},
		},
	}},
	{[]int{}, nil},
}

func TestBuildTreeFromPostorder(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := BuildTreeFromPostorder(tc.postorder); !equal(result, tc.expected) {
			t.Errorf("Trees don't match for tcid%d", tcid)
		}
	}
}

func BenchmarkBuildTreeFromPostorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BuildTreeFromPostorder(tc.postorder)
		}
	}
}

func TestBuildTreeFromPostorderLinear(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := BuildTreeFromPostorderLinear(tc.postorder); !equal(result, tc.expected) {
			t.Errorf("Trees don't match for tcid%d", tcid)
		}
	}
}

func BenchmarkBuildTreeFromPostorderLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			BuildTreeFromPostorderLinear(tc.postorder)
		}
	}
}

func equal(a, b *BinaryTree) bool {
	if a == nil && b != nil {
		return false
	} else if a != nil && b == nil {
		return false
	} else if a == nil && b == nil {
		return true
	} else if a.Value != b.Value {
		return false
	}
	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}
