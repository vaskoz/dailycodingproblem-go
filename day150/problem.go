package day150

import (
	"container/heap"
	"math"
)

// Point has an X and Y coordinate.
type Point struct {
	X, Y int
}

type item struct {
	dist float64
	pt   Point
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].dist > pq[j].dist
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

// KNearestPoints returns the k-nearest points closest to a center point.
// Runs in O(N log K) time.
func KNearestPoints(pts []Point, center Point, k int) []Point {
	pq := make(priorityQueue, 0, k)
	for _, pt := range pts {
		dist := math.Sqrt(
			math.Pow(float64(pt.X-center.X), 2.0) +
				math.Pow(float64(pt.Y-center.Y), 2.0))
		if pq.Len() == k {
			if dist < pq[0].dist {
				heap.Pop(&pq)
				heap.Push(&pq, &item{dist, pt})
			}
		} else if pq.Len() < k {
			heap.Push(&pq, &item{dist, pt})
		}
	}
	result := make([]Point, 0, k)
	for pq.Len() > 0 {
		it := heap.Pop(&pq).(*item)
		result = append(result, it.pt)
	}
	return result
}
