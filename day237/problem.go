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
	if tree == nil {
		return true
	}
	level := tree.Children
	for {
		nextLevel := make([]*KaryTree, 0, len(level)*2)
		for _, l := range level {
			if l == nil {
				nextLevel = append(nextLevel, l)
			} else {
				nextLevel = append(nextLevel, l.Children...)
			}
		}
		for i := 0; i < len(nextLevel)/2; i++ {
			if !equal(nextLevel[i], nextLevel[len(nextLevel)-1-i]) {
				return false
			}
		}
		done := true
		for i := range nextLevel {
			if nextLevel[i] != nil {
				done = false
				break
			}
		}
		if done {
			break
		}
		level = nextLevel
	}
	return true
}

func equal(a, b *KaryTree) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil && a.Value == b.Value {
		return true
	}
	return false
}
