package day319

import (
	"container/heap"
	"errors"
)

// nolint
var (
	errNotSolvable = errors.New("this puzzle is unsolvable")
	answer         = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 0}
)

// Puzzle represents and 8-puzzle in row-major form.
type Puzzle [9]int

type node struct {
	parent *node
	grid   Puzzle
	r, c   int
	cost   int
	level  int
}

type priorityQueue []*node

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*node))
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]

	return item
}

// EightPuzzle solves an 8-puzzle or returns
// an error if unsolvable.
func EightPuzzle(puzzle Puzzle) ([]Puzzle, error) {
	if inv := inversions(puzzle); inv%2 == 1 {
		return nil, errNotSolvable
	}

	seen := make(map[Puzzle]struct{})

	var r, c int

	for i, v := range puzzle {
		if v == 0 {
			r = i / 3
			c = i % 3
		}
	}

	var min *node

	root := &node{
		parent: nil,
		grid:   puzzle,
		cost:   cost(puzzle),
		r:      r,
		c:      c,
	}

	seen[puzzle] = struct{}{}

	pq := make(priorityQueue, 0)
	rows := []int{1, 0, -1, 0}
	cols := []int{0, -1, 0, 1}

	heap.Push(&pq, root)

	for pq.Len() > 0 {
		min = heap.Pop(&pq).(*node)

		if min.cost == 0 {
			break
		}

		for i := range rows {
			newR := min.r + rows[i]
			newC := min.c + cols[i]

			if inBounds(newR, newC) {
				g := min.grid
				posA := min.r*3 + min.c
				posB := newR*3 + newC

				g[posA], g[posB] = g[posB], g[posA]

				if _, found := seen[g]; !found {
					seen[g] = struct{}{}

					child := &node{
						parent: min,
						grid:   g,
						cost:   cost(g),
						level:  min.level + 1,
						r:      newR,
						c:      newC,
					}

					heap.Push(&pq, child)
				}
			}
		}
	}

	return path(min), nil
}

func path(n *node) []Puzzle {
	var res []Puzzle

	for n != nil {
		res = append(res, n.grid)
		n = n.parent
	}

	for i := 0; i < len(res)/2; i++ {
		j := len(res) - 1 - i
		res[i], res[j] = res[j], res[i]
	}

	return res
}

func inversions(puz Puzzle) int {
	inv := 0

	for i := range puz {
		for j := i + 1; j < len(puz); j++ {
			if puz[i] != 0 && puz[j] != 0 && puz[i] > puz[j] {
				inv++
			}
		}
	}

	return inv
}

func cost(puz Puzzle) int {
	total := 0

	for i, v := range puz {
		if v != answer[i] {
			total++
		}
	}

	return total
}

func inBounds(r, c int) bool {
	return r >= 0 && r < 3 && c >= 0 && c < 3
}
