package day215

import (
	"sort"
)

// BinaryTree is a binary tree that can hold any value.
type BinaryTree struct {
	Value       interface{}
	Left, Right *BinaryTree
}

type value struct {
	height int
	value  interface{}
}

// BottomView returns the bottom-view of the binary
// tree.
func BottomView(root *BinaryTree) []interface{} {
	pos := make(map[int]value)
	bottomView(root, 0, 0, pos)
	result := make([]interface{}, len(pos))
	keys := make([]int, 0, len(pos))
	for k := range pos {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for i, key := range keys {
		result[i] = pos[key].value
	}
	return result
}

func bottomView(node *BinaryTree, dist, height int, pos map[int]value) {
	if node == nil {
		return
	}
	if val, found := pos[dist]; found {
		if height >= val.height {
			pos[dist] = value{height, node.Value}
		}
	} else {
		pos[dist] = value{height, node.Value}
	}
	bottomView(node.Left, dist-1, height+1, pos)
	bottomView(node.Right, dist+1, height+1, pos)
}
