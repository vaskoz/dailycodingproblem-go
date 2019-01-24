package day154

import "container/heap"

// StackHeap is a stack implemented using only a heap.
type StackHeap struct {
	heap    maxHeap
	counter uint64
}

type item struct {
	index uint64
	value interface{}
}

type maxHeap []item

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].index > h[j].index }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1].value
	*h = old[:n-1]
	return x
}

// Push adds an element to the stack.
func (sh *StackHeap) Push(v interface{}) {
	heap.Push(&sh.heap, item{sh.counter, v})
	sh.counter++
}

// Pop removes an element from the stack.
func (sh *StackHeap) Pop() interface{} {
	return heap.Pop(&sh.heap)
}
