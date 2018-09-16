package day26

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	list     []int
	kth      int
	expected []int
}{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, []int{1, 2, 3, 4, 5, 6, 7, 9, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, []int{1, 3, 4, 5, 6, 7, 8, 9, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}},
}

func TestRemoveFromEnd(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		head := createList(tc.list)
		newhead := RemoveFromEnd(head, tc.kth)
		result := backToSlice(newhead)
		if !reflect.DeepEqual(tc.expected, result) {
			t.Errorf("Expected %v but got %v", tc.expected, result)
		}
	}
}

func BenchmarkRemoveFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			head := createList(tc.list)
			RemoveFromEnd(head, tc.kth)
		}
	}
}

func createList(data []int) *Node {
	head := &Node{data[0], nil}
	current := head
	for i := 1; i < len(data); i++ {
		current.Next = &Node{data[i], nil}
		current = current.Next
	}
	return head
}

func backToSlice(head *Node) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
