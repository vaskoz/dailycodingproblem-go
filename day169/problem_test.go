package day169

import "testing"

var testcases = []struct {
	input    *SinglyLL
	expected *SinglyLL
}{
	{&SinglyLL{4, &SinglyLL{1, &SinglyLL{-3, &SinglyLL{99, nil}}}},
		&SinglyLL{-3, &SinglyLL{1, &SinglyLL{4, &SinglyLL{99, nil}}}}},
}

func TestMergesortSinglyLinkedList(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		input := copySinglyLL(tc.input)
		if result := MergesortSinglyLinkedList(input); !equal(result, tc.expected) {
			t.Errorf("Lists don't match for tcid%d", tcid)
		}
	}
}

func BenchmarkMergesortSinglyLinkedList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			input := copySinglyLL(tc.input)
			MergesortSinglyLinkedList(input)
		}
	}
}

func copySinglyLL(a *SinglyLL) *SinglyLL {
	result := &SinglyLL{}
	ptr := result
	for a != nil {
		ptr.Next = &SinglyLL{a.Value, nil}
		ptr = ptr.Next
		a = a.Next
	}
	return result.Next
}

func equal(a, b *SinglyLL) bool {
	for a != nil && b != nil {
		if a.Value != b.Value {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}
