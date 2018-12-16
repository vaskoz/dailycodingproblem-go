package day112

// Node has a pointer to its parent.
type Node struct {
	Parent *Node
}

// LowestCommonAncestor returns the lowest common ancestor.
// Returns nil if the two nodes don't share a common node.
// Assumes the graph is a DAG. Any cycles returns nil.
// Runs in O(N) time and stores 1 path so technically O(N) memory.
func LowestCommonAncestor(n1, n2 *Node) *Node {
	path := make(map[*Node]struct{})
	for n1 != nil {
		if _, seen := path[n1]; seen {
			return nil // there is a cycle
		}
		path[n1] = struct{}{}
		n1 = n1.Parent
	}
	path2 := make(map[*Node]struct{})
	for n2 != nil {
		if _, seen := path2[n2]; seen {
			return nil // there is a cycle
		}
		path2[n2] = struct{}{}
		if _, found := path[n2]; found {
			return n2
		}
		n2 = n2.Parent
	}
	return nil
}
