package day24

import "testing"

func createSmallTree() []*LockableBinaryTree {
	root := &LockableBinaryTree{}
	left := &LockableBinaryTree{Parent: root}
	right := &LockableBinaryTree{Parent: root}
	root.Left = left
	root.Right = right
	return []*LockableBinaryTree{root, left, right}
}

func TestLockableBinaryTree(t *testing.T) {
	t.Parallel()
	nodes := createSmallTree()
	if locked := nodes[2].Lock(); !locked {
		t.Error("Should be able to lock this node")
	}
	if locked := nodes[0].Lock(); locked {
		t.Error("Should be able not be lockable because a child is locked")
	}
	if locked := nodes[1].Lock(); !locked {
		t.Error("Locking other side of subtree should work")
	}
	if unlocked := nodes[2].Unlock(); !unlocked {
		t.Error("Should unlock")
	}
	if locked := nodes[0].Lock(); locked {
		t.Error("Should be able not be lockable because a child is locked")
	}
	if unlocked := nodes[1].Unlock(); !unlocked {
		t.Error("Should unlock")
	}
	if locked := nodes[0].Lock(); !locked {
		t.Error("Now the root is locked")
	}
	if locked := nodes[1].Lock(); locked {
		t.Error("If root is locked can't lock left child")
	}
	if locked := nodes[2].Lock(); locked {
		t.Error("If root is locked can't lock right child")
	}
}

func BenchmarkLockableBinaryTree(b *testing.B) {
	nodes := createSmallTree()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nodes[2].Lock()
		nodes[0].Lock()
		nodes[1].Lock()
		nodes[2].Unlock()
		nodes[0].Lock()
		nodes[1].Unlock()
		nodes[0].Lock()
		nodes[1].Lock()
		nodes[2].Lock()
		nodes[0].Unlock()
		nodes[1].Unlock()
		nodes[2].Unlock()
	}
}
