package day270

import (
	"container/heap"
)

// Edge represents a directed edge from a start to an end with a time cost.
type Edge struct {
	Start, End, Time int
}

type nodeVisit struct {
	id   int
	time int
}
type priorityQueue []nodeVisit

// TimeAllNodesReceiveMessage returns the time required for all edges to
// receive the message originating from node '0'.
// Nodes are from 0 to N inclusive.
func TimeAllNodesReceiveMessage(n int, edges []Edge) int {
	graph := make(map[int]map[int]int)
	for _, edge := range edges {
		if _, exists := graph[edge.Start]; !exists {
			graph[edge.Start] = make(map[int]int)
		}
		graph[edge.Start][edge.End] = edge.Time
	}
	totalTime := 0
	visited := make(map[int]struct{})
	visited[0] = struct{}{}
	pq := make(priorityQueue, n+1)
	heap.Push(&pq, nodeVisit{0, 0})
	for len(visited) != n+1 {
		next := heap.Pop(&pq).(nodeVisit)
		visited[next.id] = struct{}{}
		totalTime = next.time
		for node, time := range graph[next.id] {
			if _, seen := visited[node]; seen {
				continue
			}
			heap.Push(&pq, nodeVisit{node, next.time + time})
		}
	}
	return totalTime
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].time < pq[j].time
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(nodeVisit)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
