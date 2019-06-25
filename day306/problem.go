package day306

import "container/heap"

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// SortK sorts the input that is shifted by at most
// k positions and returns a new output.
// Runs in O(N log K) time.
func SortK(nums []int, k int) []int {
	size := 2*k + 1
	h := make(intHeap, 0, size)
	heap.Init(&h)
	sorted := make([]int, 0, len(nums))
	for i, num := range nums {
		if i >= size {
			smallest := heap.Pop(&h).(int)
			sorted = append(sorted, smallest)
		}
		heap.Push(&h, num)
	}
	for h.Len() > 0 {
		smallest := heap.Pop(&h).(int)
		sorted = append(sorted, smallest)
	}
	return sorted
}
