package day398

import "testing"

// nolint
var testcases = []struct {
	head     *SinglyLL
	kth      int
	expected *SinglyLL
}{
	{
		&SinglyLL{1, &SinglyLL{2, &SinglyLL{3, &SinglyLL{4, &SinglyLL{5, nil}}}}},
		2,
		&SinglyLL{1, &SinglyLL{2, &SinglyLL{3, &SinglyLL{5, nil}}}},
	},
	{
		&SinglyLL{1, &SinglyLL{2, &SinglyLL{3, &SinglyLL{4, &SinglyLL{5, nil}}}}},
		5,
		&SinglyLL{2, &SinglyLL{3, &SinglyLL{4, &SinglyLL{5, nil}}}},
	},
}

func TestRemoveKthFromEnd(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if res := RemoveKthFromEnd(tc.head, tc.kth); !equal(res, tc.expected) {
			t.Errorf("TCID%d singly linked lists do not match", tcid)
		}
	}
}

func TestRemoveKthFromEndPanic(t *testing.T) {
	t.Parallel()

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected a panic when k is larger than length")
		}
	}()

	head := &SinglyLL{1, &SinglyLL{2, &SinglyLL{3, &SinglyLL{4, &SinglyLL{5, nil}}}}}
	RemoveKthFromEnd(head, 6)
}

func BenchmarkRemoveKthFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RemoveKthFromEnd(tc.head, tc.kth)
		}
	}
}

func equal(a, b *SinglyLL) bool {
	for ; a != nil && b != nil; a, b = a.Next, b.Next {
		if a.Value != b.Value {
			return false
		}
	}

	return a == nil && b == nil
}
