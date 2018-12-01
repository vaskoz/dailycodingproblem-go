package day93

// BinaryTree is a binary tree: left & right children.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// TreeInfo contains information from a given subtree.
type TreeInfo struct {
	IsBST  bool
	Min    int
	Max    int
	Size   int
	Ans    int
	AnsPtr *BinaryTree
}

// MaxInt is the largest int.
const MaxInt = int(^uint(0) >> 1)

// MinInt is the smallest int.
const MinInt = -MaxInt - 1

// FindLargestBST will return a pointer to the root of the
// largest BST in the binary tree.
// Runs in O(N) time with no extra space.
func FindLargestBST(tree *BinaryTree) *BinaryTree {
	res := sizeOfLargestBST(tree)
	return res.AnsPtr
}

// SizeOfLargestBST will return the number of nodes in the
// largest BST in the binary tree.
// Runs in O(N) time with no extra space.
func SizeOfLargestBST(tree *BinaryTree) int {
	res := sizeOfLargestBST(tree)
	return res.Ans
}

func sizeOfLargestBST(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{IsBST: true, Min: MaxInt, Max: MinInt, Size: 0, Ans: 0}
	}
	if tree.Left == nil && tree.Right == nil {
		return TreeInfo{IsBST: true, Min: tree.Value, Max: tree.Value, Size: 1, Ans: 1}
	}
	left := sizeOfLargestBST(tree.Left)
	right := sizeOfLargestBST(tree.Right)
	var info TreeInfo
	info.Size = 1 + left.Size + right.Size
	if left.IsBST && right.IsBST && left.Max < tree.Value && right.Min > tree.Value {
		info.Min = min(left.Min, min(right.Min, tree.Value))
		info.Max = max(right.Max, max(left.Max, tree.Value))
		info.Ans = info.Size
		info.AnsPtr = tree
		info.IsBST = true
		return info
	}
	if left.Ans > right.Ans {
		info.Ans = left.Ans
		info.AnsPtr = left.AnsPtr
	} else {
		info.Ans = right.Ans
		info.AnsPtr = right.AnsPtr
	}
	info.Ans = max(left.Ans, right.Ans)
	info.IsBST = false
	return info
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
