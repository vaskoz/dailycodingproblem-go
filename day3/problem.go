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

func preorderDeserialize(serialized []string) *Node {
	if len(serialized) == 0 || serialized[0] == "-1" {
		return nil
	}
	v, err := strconv.Atoi(serialized[0])
	if err != nil {
		panic("bad serialized data cannot be deserialized")
	}
	node := &Node{val: v}
	node.left = preorderDeserialize(serialized[1:])
	node.right = preorderDeserialize(serialized[2:])
	return node
}

// Deserialize converts a string into a general binary tree.
func Deserialize(data string) *Node {
	serialized := strings.Split(data, " ")
	return preorderDeserialize(serialized)
}
