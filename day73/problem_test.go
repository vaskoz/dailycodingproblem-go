package day73

import "testing"

var testcases = []struct {
	input []interface{}
}{
	{[]interface{}{"first", 2, "third", 4, "fifth", 6}},
	{[]interface{}{"foo", "bar", "baz"}},
	{[]interface{}{1, 2, 3}},
}

func createSinglyLL(data []interface{}) *SinglyLL {
	head := &SinglyLL{}
	current := head
	for _, value := range data {
		current.Next = &SinglyLL{Value: value, Next: nil}
		current = current.Next
	}
	return head.Next
}

func wasReversed(head *SinglyLL, data []interface{}) bool {
	for i := range data {
		if head == nil || head.Value != data[len(data)-1-i] {
			return false
		}
		head = head.Next
	}
	return head == nil
}

func TestReverseSinglyLinkedList(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		sll := createSinglyLL(tc.input)
		if reversed := ReverseSinglyLinkedList(sll); !wasReversed(reversed, tc.input) {
			t.Errorf("This list was not reversed correctly: %v", tc.input)
		}
	}
}

func BenchmarkReverseSinglyLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			sll := createSinglyLL(tc.input)
			ReverseSinglyLinkedList(sll)
		}
	}
}
