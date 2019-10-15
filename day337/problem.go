package day337

import (
	"math/rand"
	"time"
)

// SinglyLL is a singly linked list of any kind of value.
// Type management falls to the user.
type SinglyLL struct {
	Val  interface{}
	Next *SinglyLL
}

// nolint
var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// FastShuffle uniformly shuffles the singly linked list in
// O(N) time and O(N) space.
// Technically, O(2N) space. First N is for slice of values, second N
// is for new Linked List (not done in-place and doesn't mutate input linked list).
func FastShuffle(head *SinglyLL) *SinglyLL {
	data := convertToSlice(head)

	r.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	return convertToLL(data)
}

// MemoryOptimizedShuffle uniformly shuffles the singly linked list
// in O(N^2) time and O(1) space.
func MemoryOptimizedShuffle(head *SinglyLL) *SinglyLL {
	n := size(head)
	origHolder := &SinglyLL{nil, head}
	resultHolder := &SinglyLL{}
	resultTail := resultHolder

	for items := n; items > 0; items-- {
		curr := origHolder

		for i := 0; i < r.Intn(items); i++ {
			curr = curr.Next
		}

		resultTail.Next = curr.Next
		curr.Next = curr.Next.Next
		resultTail = resultTail.Next
		resultTail.Next = nil
	}

	return resultHolder.Next
}

func size(head *SinglyLL) int {
	var count int

	for ; head != nil; head = head.Next {
		count++
	}

	return count
}

func convertToLL(data []interface{}) *SinglyLL {
	holder := &SinglyLL{}
	current := holder

	for i := range data {
		current.Next = &SinglyLL{data[i], nil}
		current = current.Next
	}

	return holder.Next
}

func convertToSlice(head *SinglyLL) []interface{} {
	var result []interface{}

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}
