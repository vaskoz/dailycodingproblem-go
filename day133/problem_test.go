package day133

import "testing"

var tree = createBST()

var testcases = []struct {
	givenNode, inorderSuccessor *BST
}{
	{nil, nil},
	{tree.Right.Left, tree.Right},
	{tree, tree.Right.Left},
	{tree.Left, tree},
	{tree.Right, tree.Right.Right},
}

func TestInorderSuccessor(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := InorderSuccessor(tc.givenNode); result != tc.inorderSuccessor {
			t.Errorf("Nodes don't match for TCID%d", tcid)
		}
	}
}

func BenchmarkInorderSuccessor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			InorderSuccessor(tc.givenNode)
		}
	}
}

func createBST() *BST {
	bst := &BST{
		Value:  10,
		Parent: nil,
		Left:   &BST{5, nil, nil, nil},
		Right: &BST{
			Value: 30,
			Left:  &BST{22, nil, nil, nil},
			Right: &BST{35, nil, nil, nil},
		},
	}
	bst.Left.Parent = bst
	bst.Right.Parent = bst
	bst.Right.Left.Parent = bst.Right
	bst.Right.Right.Parent = bst.Right
	return bst
}
