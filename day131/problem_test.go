package day131

import (
	"testing"
)

var testcases = []struct {
	head *RandomSinglyLL
}{
	{nil},
	{&RandomSinglyLL{}},
	{&RandomSinglyLL{Value: 10, Next: &RandomSinglyLL{Value: 20, Next: &RandomSinglyLL{30, nil, nil}}}},
	{generateRandomSinglyLL()},
}

func TestDeepCloneRandomSinglyLL(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := DeepCloneRandomSinglyLL(tc.head); !equalButNotIdentical(tc.head, result) {
			t.Errorf("DeepClone didn't work for TCID%d", tcid)
		}
	}
}

func BenchmarkDeepCloneRandomSinglyLL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DeepCloneRandomSinglyLL(tc.head)
		}
	}
}

func equalButNotIdentical(one, two *RandomSinglyLL) bool {
	for one != nil && two != nil {
		if checks(one, two) {
			return false
		}
		one = one.Next
		two = two.Next
	}
	return one == nil && two == nil
}

func checks(one, two *RandomSinglyLL) bool {
	return one == two || one.Value != two.Value || (one.Random != nil && one.Random == two.Random) ||
		(one.Random != nil && two.Random != nil && one.Random.Value != two.Random.Value)
}

func generateRandomSinglyLL() *RandomSinglyLL {
	head := &RandomSinglyLL{Value: 10, Next: &RandomSinglyLL{Value: 20, Next: &RandomSinglyLL{30, nil, nil}}}
	head.Random = head.Next.Next
	head.Next.Random = head
	return head
}
