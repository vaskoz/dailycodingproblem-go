package day3

import (
	//	"log"
	"testing"
)

func TestSerializeAndDeserialize(t *testing.T) {
	t.Parallel()
	root := &Node{val: 20,
		left: nil,
		right: &Node{val: 8, left: nil,
			right: &Node{val: 10, left: nil,
				right: &Node{val: 5}}},
	}
	expected := "20 -1 8 -1 10 -1 5 -1 -1"
	var result string
	if result = Serialize(root); result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
	if root2 := Deserialize(result); !TreeEquality(root, root2) {
		t.Error("Trees are not equal")
	}
}

func TestDeserializeBadString(t *testing.T) {
	t.Parallel()
	defer func() {
		if r := recover(); r == nil {
			t.Error("I expected a panic on bad serialization data")
		}
	}()
	badSerializedData := "foobar"
	Deserialize(badSerializedData)
}

// TreeEquality tests for binary tree equality
func TreeEquality(t1, t2 *Node) bool {
	switch {
	case t1 == nil && t2 == nil:
		return true
	case t1 == nil || t2 == nil:
		return false
	case t1.val != t2.val:
		return false
	default:
		return TreeEquality(t1.left, t2.left) && TreeEquality(t1.right, t2.right)
	}
}

// inorder is a method to debug your inorder traversal of the tree
//func inorder(n *Node) {
//if n == nil {
//return
//}
//if n.left != nil {
//log.Println("going left")
//} else {
//log.Println("nothing to the left")
//}
//inorder(n.left)
//log.Println(n.val)
//if n.right != nil {
//log.Println("going right")
//} else {
//log.Println("nothing to the right")
//}
//inorder(n.right)
//}
