package day208

import (
	"log"
	"testing"
)

var testcases = []struct {
	head     *SinglyLL
	k        int
	expected *SinglyLL
}{
	{
		&SinglyLL{5, &SinglyLL{1, &SinglyLL{8, &SinglyLL{0, &SinglyLL{3, nil}}}}},
		3,
		&SinglyLL{1, &SinglyLL{0, &SinglyLL{5, &SinglyLL{8, &SinglyLL{3, nil}}}}},
	},
}

func TestPartitionSinglyLL(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := PartitionSinglyLL(tc.head, tc.k); !equal(result, tc.expected) {
			log.Println("Expected")
			printLL(tc.expected)
			log.Println("Received")
			printLL(result)
			t.Errorf("linked list isn't partitioned as expected for tcid%d", tcid)
		}
	}
}

func BenchmarkPartitionSinglyLL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			PartitionSinglyLL(tc.head, tc.k)
		}
	}
}

func printLL(a *SinglyLL) {
	for a != nil {
		log.Println(a)
		a = a.Next
	}
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
