package day78

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	lists    [][]int
	expected []int
}{
	{[][]int{
		{1, 5, 10},
		{2, 6, 11},
		{3, 7, 12},
	},
		[]int{1, 2, 3, 5, 6, 7, 10, 11, 12},
	},
	{[][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

func TestMergeKSortedLists(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		var input []*SinglyLL
		for i := range tc.lists {
			input = append(input, createSinglyLL(tc.lists[i]))
		}
		head := MergeKSortedLists(input)
		result := convertSinglyLLToSlice(head)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v got %v", tc.expected, result)
		}
	}
}

func BenchmarkMergeKSortedLists(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			var input []*SinglyLL
			for i := range tc.lists {
				input = append(input, createSinglyLL(tc.lists[i]))
			}
			MergeKSortedLists(input)
		}
	}
}

func createSinglyLL(data []int) *SinglyLL {
	if len(data) == 0 {
		return nil
	}
	head := &SinglyLL{data[0], nil}
	current := head
	for i := 1; i < len(data); i++ {
		current.Next = &SinglyLL{data[i], nil}
		current = current.Next
	}
	return head
}

func convertSinglyLLToSlice(head *SinglyLL) []int {
	var result []int
	for head != nil {
		result = append(result, head.Value)
		head = head.Next
	}
	return result
}
