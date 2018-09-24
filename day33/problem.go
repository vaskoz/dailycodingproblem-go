package day33

import "container/heap"

type MinPQ []int
type MaxPQ []int

type RunningMedian struct {
	low  *MaxPQ
	high *MinPQ
}

func (pq MinPQ) Len() int { return len(pq) }

func (pq MinPQ) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq MinPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinPQ) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *MinPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq MaxPQ) Len() int { return len(pq) }

func (pq MaxPQ) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq MaxPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxPQ) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

// Pop is not called, but needed to meet the interface.
func (pq *MaxPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// NewRunningMedian returns a new instance of a RunningMedian processor.
func NewRunningMedian() *RunningMedian {
	return &RunningMedian{
		&MaxPQ{},
		&MinPQ{},
	}
}

// Median returns the running median after adding this value.
func (rm *RunningMedian) Median(val int) float64 {
	heap.Push(rm.high, val)
	if rm.high.Len() > rm.low.Len()+1 {
		v := heap.Pop(rm.high)
		heap.Push(rm.low, v)
	}
	if rm.high.Len() == rm.low.Len() {
		return float64((*rm.high)[0]+(*rm.low)[0]) / 2
	}
	return float64((*rm.high)[0])
}
