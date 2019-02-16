package day177

import "testing"

var testcases = []struct {
	input    *LL
	k        int
	expected *LL
}{
	{&LL{7, &LL{7, &LL{3, &LL{5, nil}}}}, 2, &LL{3, &LL{5, &LL{7, &LL{7, nil}}}}},
	{&LL{1, &LL{2, &LL{3, &LL{4, &LL{5, nil}}}}}, 3, &LL{4, &LL{5, &LL{1, &LL{2, &LL{3, nil}}}}}},
	{nil, 10, nil},
	{&LL{5, nil}, 10000, &LL{5, nil}},
	{&LL{1, &LL{2, &LL{3, &LL{4, &LL{5, nil}}}}}, 503, &LL{4, &LL{5, &LL{1, &LL{2, &LL{3, nil}}}}}},
}

func TestRotateRightLL(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		copied := cloneLL(tc.input)
		if result := RotateRightLL(copied, tc.k); !equal(tc.expected, result) {
			t.Errorf("TCID%d doesn't match expected", tcid)
		}
	}
}

func BenchmarkRotateRightLL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			copied := cloneLL(tc.input)
			RotateRightLL(copied, tc.k)
		}
	}
}

func cloneLL(a *LL) *LL {
	head := &LL{}
	for ptr := head; a != nil; ptr, a = ptr.Next, a.Next {
		ptr.Next = &LL{a.Value, nil}
	}
	return head.Next
}

func equal(a, b *LL) bool {
	for ; a != nil && b != nil; a, b = a.Next, b.Next {
		if a.Value != b.Value {
			return false
		}
	}
	return a == nil && b == nil
}
