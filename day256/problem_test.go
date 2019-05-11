package day256

import "testing"

// nolint
var testcases = []struct {
	input, expected *SinglyLL
}{
	{nil, nil},
	{
		&SinglyLL{1, &SinglyLL{2, &SinglyLL{3, &SinglyLL{4, &SinglyLL{5, nil}}}}},
		&SinglyLL{1, &SinglyLL{3, &SinglyLL{2, &SinglyLL{5, &SinglyLL{4, nil}}}}},
	},
}

func equal(a, b *SinglyLL) bool {
	for a != nil && b != nil {
		if a.Value != b.Value {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == b
}

func copyLL(a *SinglyLL) *SinglyLL {
	newhead := &SinglyLL{}
	curr := newhead
	for a != nil {
		curr.Next = &SinglyLL{a.Value, nil}
		a = a.Next
		curr = curr.Next
	}
	return newhead.Next
}

func TestRearrangeAlternating(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		input := copyLL(tc.input)
		if result := RearrangeAlternating(input); !equal(result, tc.expected) {
			t.Errorf("testcase #%d does not match", tcid)
		}
	}
}

func BenchmarkRearrangeAlternating(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			input := copyLL(tc.input)
			RearrangeAlternating(input)
		}
	}
}
