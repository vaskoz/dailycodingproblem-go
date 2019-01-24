package day154

import "testing"

var testcases = []struct {
	items []interface{}
}{
	{[]interface{}{"foo", "bar", 1, 3, 5}},
	{[]interface{}{10, 20, 30, 15, 30, 55}},
}

func TestStackHeap(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		var stack StackHeap
		for i := range tc.items {
			stack.Push(tc.items[i])
		}
		for i := range tc.items {
			if result := stack.Pop(); result != tc.items[len(tc.items)-1-i] {
				t.Errorf("Expected %v got %v", tc.items[len(tc.items)-1-i], result)
			}
		}
	}
}

func BenchmarkStackHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			var stack StackHeap
			for i := range tc.items {
				stack.Push(tc.items[i])
			}
			for range tc.items {
				stack.Pop()
			}
		}
	}
}
