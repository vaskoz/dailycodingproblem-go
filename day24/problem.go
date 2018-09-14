package day24

// LockableBinaryTree implements a binary tree where nodes
// can be locked.
type LockableBinaryTree struct {
	Locked              bool
	LockedDescendants   int
	Left, Right, Parent *LockableBinaryTree
}

// IsLocked returns true if locked, false otherwise.
func (lbt *LockableBinaryTree) IsLocked() bool {
	return lbt.Locked
}

// Lock acquires a lock on this node if possible.
// Returns true if lock is acquired, false if it wasn't.
func (lbt *LockableBinaryTree) Lock() bool {
	if lbt.canModify() {
		lbt.Locked = true
		cur := lbt.Parent
		for cur != nil {
			cur.LockedDescendants++
			cur = cur.Parent
		}
		return true
	}
	return false
}

// Unlock frees this node if possible.
// Returns true if unlocked, false if it couldn't be.
func (lbt *LockableBinaryTree) Unlock() bool {
	if lbt.canModify() {
		lbt.Locked = false
		cur := lbt.Parent
		for cur != nil {
			cur.LockedDescendants--
			cur = cur.Parent
		}
		return true
	}
	return false
}

func (lbt *LockableBinaryTree) canModify() bool {
	if lbt.LockedDescendants > 0 {
		return false
	}
	cur := lbt.Parent
	for cur != nil {
		if cur.IsLocked() {
			return false
		}
		cur = cur.Parent
	}
	return true
}
