package mergeklist

import (
	"container/heap"
	"math"
)

// MergeKSortedLists merges K-sorted lists into a single list
// Runtime is O(K*(K*N)) = O(K^2*N) assuming each list of is length N
// Space is O(K) for the indicies without considering the answer list of O(K*N) size.
// Use this if K is very small
func MergeKSortedLists(lists [][]int) []int {
	idxs := make([]int64, len(lists))
	var result []int
	hasMore := true
	for hasMore {
		listID := -1
		smallest := math.MaxInt64
		// find the smallest next value
		for i, idx := range idxs {
			if idx < int64(len(lists[i])) && lists[i][idx] < smallest {
				smallest = lists[i][idx]
				listID = i
			}
		}
		if listID == -1 {
			hasMore = false
		} else {
			result = append(result, smallest)
			idxs[listID]++
		}
	}
	return result
}

type tuple struct {
	id, pos, value int
}

type priorityQueue []*tuple

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*tuple)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	tuple := old[n-1]
	*pq = old[0 : n-1]
	return tuple
}

// MergeKSortedListsUsingHeap merges K-sorted lists into a single list using a min-heap.
// Runtime is O(N*K log K) assuming each list of is length N
// Space is O(K) for the heap without considering the answer list of O(K*N) size.
// Use this if K is very large.
func MergeKSortedListsUsingHeap(lists [][]int) []int {
	pq := make(priorityQueue, len(lists))
	for i, lst := range lists {
		pq[i] = &tuple{id: i, pos: 0, value: lst[0]}
	}
	heap.Init(&pq)
	var result []int
	for pq.Len() > 0 {
		tup := heap.Pop(&pq).(*tuple)
		result = append(result, tup.value)
		id := tup.id
		nextpos := tup.pos + 1
		if nextpos < len(lists[id]) {
			heap.Push(&pq, &tuple{id: id, pos: nextpos, value: lists[id][nextpos]})
		}
	}
	return result
}
