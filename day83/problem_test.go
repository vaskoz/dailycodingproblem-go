package day83

import "testing"

func TestInvertBinaryTree(t *testing.T) {
	t.Parallel()
	input := &BinaryTree{"a",
		&BinaryTree{"b", &BinaryTree{"d", nil, nil}, &BinaryTree{"e", nil, nil}},
		&BinaryTree{"c", &BinaryTree{"f", nil, nil}, nil}}
	original := &BinaryTree{"a",
		&BinaryTree{"b", &BinaryTree{"d", nil, nil}, &BinaryTree{"e", nil, nil}},
		&BinaryTree{"c", &BinaryTree{"f", nil, nil}, nil}}
	inverted := &BinaryTree{"a",
		&BinaryTree{"c", nil, &BinaryTree{"f", nil, nil}},
		&BinaryTree{"b", &BinaryTree{"e", nil, nil}, &BinaryTree{"d", nil, nil}}}
	InvertBinaryTree(input)
	if !equalTrees(input, inverted) {
		t.Errorf("Expected it to equal the inverted tree")
	}
	InvertBinaryTree(input)
	if !equalTrees(input, original) {
		t.Errorf("Invert an invert should match the original")
	}
}

func BenchmarkInvertBinaryTree(b *testing.B) {
	input := &BinaryTree{"a",
		&BinaryTree{"b", &BinaryTree{"d", nil, nil}, &BinaryTree{"e", nil, nil}},
		&BinaryTree{"c", &BinaryTree{"f", nil, nil}, nil}}
	for i := 0; i < b.N; i++ {
		InvertBinaryTree(input)
	}
}

func equalTrees(one, two *BinaryTree) bool {
	if one == nil && two == nil {
		return true
	} else if (one == nil && two != nil) || (one != nil && two == nil) || (one.Value != two.Value) {
		return false
	}
	return equalTrees(one.Left, two.Left) || equalTrees(one.Right, two.Right)
}
