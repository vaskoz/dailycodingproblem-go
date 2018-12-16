package day116

// UnboundBinaryTree is a tree that lazily grows.
type UnboundBinaryTree interface {
	Value() interface{}
	Left() UnboundBinaryTree
	Right() UnboundBinaryTree
}

type ubtTree struct {
	value       interface{}
	left, right *ubtTree
}

// Left returns the left child of this node.
// Returns nil if the variable is set to nil.
func (ubt *ubtTree) Left() UnboundBinaryTree {
	if ubt.left == nil {
		ubt.left = &ubtTree{}
	}
	return ubt.left
}

// Right returns the left child of this node.
// Returns nil if the variable is set to nil.
func (ubt *ubtTree) Right() UnboundBinaryTree {
	if ubt.right == nil {
		ubt.right = &ubtTree{}
	}
	return ubt.right
}

// Value returns the value stored at this node.
// Returns nil if the variable is set to nil.
func (ubt *ubtTree) Value() interface{} {
	return ubt.value
}

// GenerateUnboundBinaryTree runs in O(1) time.
func GenerateUnboundBinaryTree() UnboundBinaryTree {
	return &ubtTree{}
}
