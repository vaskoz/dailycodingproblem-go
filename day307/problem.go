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

// Floor returns the highest element in the tree less
// than or equal to the target.
// Returns an error if no answer exists.
func Floor(root *BinaryTree, target int) (int, error) {
	if root == nil {
		return 0, errNoAnswer
	}
	if target == root.Value {
		return target, nil
	}
	if target < root.Value {
		return Floor(root.Left, target)
	}
	ans, err := Floor(root.Right, target)
	if err != nil && root.Value <= target {
		return root.Value, nil
	}
	return ans, err
}

// Ceiling returns the lowest element in the tree greater
// than or equal to the target.
// Returns an error if no answer exists.
func Ceiling(root *BinaryTree, target int) (int, error) {
	if root == nil {
		return 0, errNoAnswer
	}
	if target == root.Value {
		return target, nil
	}
	if target < root.Value {
		ans, err := Ceiling(root.Left, target)
		if err != nil && root.Value >= target {
			return root.Value, nil
		}
		return ans, err
	}
	return Ceiling(root.Right, target)
}
