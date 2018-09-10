package day20

import (
	"testing"
)

func createTwoSharedLists() (*Node, *Node, *Node) {
	common := &Node{1015, &Node{444, nil}}
	one := &Node{"abc", &Node{"def", &Node{"ghi", common}}}
	two := &Node{111, common}
	return one, two, common
}

func createTwoIndependentLists() (*Node, *Node, *Node) {
	one := &Node{"abc", &Node{"def", &Node{"ghi", nil}}}
	two := &Node{111, nil}
	return one, two, nil
}

func TestFindCommonNode(t *testing.T) {
	one, two, common := createTwoSharedLists()
	if result := FindCommonNode(one, two); result != common {
		t.Error("Expected these two be the same pointer")
	}
	one, two = two, one // swap
	if result := FindCommonNode(one, two); result != common {
		t.Error("Expected these two be the same pointer")
	}
	one, two, common = createTwoIndependentLists()
	if result := FindCommonNode(one, two); result != common {
		t.Error("Expected these two be nil")
	}
	one, two = two, one // swap
	if result := FindCommonNode(one, two); result != common {
		t.Error("Expected these two be nil")
	}
}

func BenchmarkFindCommonNode(b *testing.B) {
	one, two, _ := createTwoSharedLists()
	onei, twoi, _ := createTwoIndependentLists()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindCommonNode(one, two)
		FindCommonNode(onei, twoi)
	}
}
