package day112

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	t.Parallel()
	lca := &Node{&Node{&Node{}}}
	n1 := &Node{lca}
	n2 := &Node{&Node{&Node{lca}}}
	if result := LowestCommonAncestor(n1, n2); result != lca {
		t.Errorf("LCA did not match correctly")
	}
	n3 := &Node{&Node{&Node{}}}
	if result := LowestCommonAncestor(n3, n2); result != nil {
		t.Errorf("LCA is nil because the nodes are from separate graphs")
	}
	cycle := &Node{&Node{}}
	cycle.Parent.Parent = cycle
	if result := LowestCommonAncestor(cycle, n2); result != nil {
		t.Errorf("LCA is nil because there is a cycle in first arg")
	}
	if result := LowestCommonAncestor(n2, cycle); result != nil {
		t.Errorf("LCA is nil because there is a cycle in second arg")
	}
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	lca := &Node{&Node{&Node{}}}
	n1 := &Node{lca}
	n2 := &Node{&Node{&Node{lca}}}
	n3 := &Node{&Node{&Node{}}}
	cycle := &Node{&Node{}}
	cycle.Parent.Parent = cycle
	for i := 0; i < b.N; i++ {
		LowestCommonAncestor(n1, n2)
		LowestCommonAncestor(n3, n2)
		LowestCommonAncestor(cycle, n2)
		LowestCommonAncestor(n2, cycle)
	}
}
