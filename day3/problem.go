package day3

import (
	"fmt"
	"strconv"
	"strings"
)

// Node represents a general binary tree.
type Node struct {
	val         int
	left, right *Node
}

func preorderSerialize(node *Node, serialized []string) []string {
	if node == nil {
		serialized = append(serialized, "-1")
	} else {
		serialized = append(serialized, fmt.Sprintf("%d", node.val))
		serialized = preorderSerialize(node.left, serialized)
		serialized = preorderSerialize(node.right, serialized)
	}
	return serialized
}

// Serialize persists a general binary tree into a string.
func Serialize(root *Node) string {
	var serialized []string
	serialized = preorderSerialize(root, serialized)
	return strings.Join(serialized, " ")
}

func preorderDeserialize(serialized []string) (*Node, int) {
	if len(serialized) == 0 || serialized[0] == "-1" {
		return nil, 1
	}
	v, err := strconv.Atoi(serialized[0])
	if err != nil {
		panic("bad serialized data cannot be deserialized")
	}
	node := &Node{val: v}
	advanceLeft, advanceRight := 0, 0
	node.left, advanceLeft = preorderDeserialize(serialized[1:])
	node.right, advanceRight = preorderDeserialize(serialized[1+advanceLeft:])
	return node, 1 + advanceLeft + advanceRight
}

// Deserialize converts a string into a general binary tree.
func Deserialize(data string) *Node {
	serialized := strings.Split(data, " ")
	result, _ := preorderDeserialize(serialized)
	return result
}
