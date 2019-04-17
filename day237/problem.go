// Package day237 has a note.
// NOTE: I didn't provide an API for a k-ary
// tree. It's up to the user to build a valid
// k-ary tree.
package day237

// KaryTree is a k-ary tree of integers.
type KaryTree struct {
	Value    int
	Children []*KaryTree
}

// IsTreeSymmetric answers if its data and shape remain
// unchanged when it is reflected about the root node.
func IsTreeSymmetric(tree *KaryTree) bool {
	return equal(tree, tree)
}

func equal(a, b *KaryTree) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil && a.Value == b.Value {
		for i := range a.Children {
			if !equal(a.Children[i], b.Children[len(a.Children)-1-i]) {
				return false
			}
		}
		return true
	}
	return false
}
