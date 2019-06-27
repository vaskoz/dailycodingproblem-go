package day307

import "errors"

var errNoAnswer = errors.New("no answer exists")

// ErrNoAnswer returns the error returned for no answer.
func ErrNoAnswer() error {
	return errNoAnswer
}

// BinaryTree is a binary tree of integers.
type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// FloorRecursive returns the highest element in the tree less
// than or equal to the target.
// Returns an error if no answer exists.
func FloorRecursive(root *BinaryTree, target int) (int, error) {
	if root == nil {
		return 0, errNoAnswer
	}
	if target == root.Value {
		return target, nil
	}
	if target < root.Value {
		return FloorRecursive(root.Left, target)
	}
	ans, err := FloorRecursive(root.Right, target)
	if err != nil && root.Value <= target {
		return root.Value, nil
	}
	return ans, err
}

// FloorFaster returns the highest element in the tree less
// than or equal to the target.
// Returns an error if no answer exists.
func FloorFaster(root *BinaryTree, target int) (int, error) {
	curr := root
	var candidate *BinaryTree
	for {
		if curr.Value == target {
			return curr.Value, nil
		}
		if target < curr.Value {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				return 0, errNoAnswer
			}
		} else {
			candidate = curr
			if curr.Right != nil {
				curr = curr.Right
			} else if candidate != nil {
				return candidate.Value, nil
			}
		}
	}
}

// CeilingRecursive returns the lowest element in the tree greater
// than or equal to the target.
// Returns an error if no answer exists.
func CeilingRecursive(root *BinaryTree, target int) (int, error) {
	if root == nil {
		return 0, errNoAnswer
	}
	if target == root.Value {
		return target, nil
	}
	if target < root.Value {
		ans, err := CeilingRecursive(root.Left, target)
		if err != nil && root.Value >= target {
			return root.Value, nil
		}
		return ans, err
	}
	return CeilingRecursive(root.Right, target)
}

// CeilingFaster returns the lowest element in the tree greater
// than or equal to the target.
// Returns an error if no answer exists.
func CeilingFaster(root *BinaryTree, target int) (int, error) {
	curr := root
	var candidate *BinaryTree
	for {
		if curr.Value == target {
			return curr.Value, nil
		}
		if target < curr.Value {
			candidate = curr
			if curr.Left != nil {
				curr = curr.Left
			} else if candidate != nil {
				return candidate.Value, nil
			}
		} else {
			switch {
			case curr.Right != nil:
				curr = curr.Right
			case candidate != nil:
				return candidate.Value, nil
			default:
				return 0, errNoAnswer
			}
		}
	}
}
