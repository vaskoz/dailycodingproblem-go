package day93

import "testing"

var testcases = []struct {
	tree           *BinaryTree
	largestBSTSize int
	largestBST     *BinaryTree
}{
	{&BinaryTree{
		Value: 60,
		Left: &BinaryTree{
			Value: 65,
			Left: &BinaryTree{
				Value: 50,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &BinaryTree{
			Value: 70,
			Left:  nil,
			Right: nil,
		},
	}, 2, nil},
	{&BinaryTree{
		Value: 60,
		Left: &BinaryTree{
			Value: 90,
			Left:  nil,
			Right: nil,
		},
		Right: &BinaryTree{
			Value: 65,
			Left:  nil,
			Right: &BinaryTree{
				Value: 70,
				Left:  nil,
				Right: nil,
			},
		},
	}, 2, nil},
	{&BinaryTree{
		Value: 60,
		Left: &BinaryTree{
			Value: 40,
			Left:  nil,
			Right: nil,
		},
		Right: &BinaryTree{
			Value: 65,
			Left:  nil,
			Right: &BinaryTree{
				Value: 70,
				Left:  nil,
				Right: nil,
			},
		},
	}, 4, nil},
}

func answerKey() {
	testcases[0].largestBST = testcases[0].tree.Left
	testcases[1].largestBST = testcases[1].tree.Right
	testcases[2].largestBST = testcases[2].tree
}

func TestSizeOfLargestBST(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := SizeOfLargestBST(tc.tree); result != tc.largestBSTSize {
			t.Errorf("Expected %v got %v", tc.largestBSTSize, result)
		}
	}
}

func BenchmarkSizeOfLargestBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SizeOfLargestBST(tc.tree)
		}
	}
}

func TestFindLargestBST(t *testing.T) {
	answerKey()
	t.Parallel()
	for _, tc := range testcases {
		if result := FindLargestBST(tc.tree); result != tc.largestBST {
			t.Errorf("Expected %v got %v", tc.largestBST, result)
		}
	}
}

func BenchmarkFindLargestBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindLargestBST(tc.tree)
		}
	}
}
