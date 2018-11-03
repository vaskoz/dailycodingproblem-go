package day72

import (
	"errors"
)

var errInfiniteLoop = errors.New("loop detected")

// AdjacencyMatrix represents a sparse adjacency matrix.
type AdjacencyMatrix map[int]map[int]struct{}

// Edge represents a directed edge going 'From' and 'To' a different node.
// Nodes are represented by integers from 0..N.
type Edge struct {
	From, To int
}

// ErrInfiniteLoop returns the error if an infinite loop is detected.
func ErrInfiniteLoop() error {
	return errInfiniteLoop
}

// LargestPathValue takes a slice of runes representing nodes.
// Edges represents all the directed edges.
// Returns the largest path value or an error if an infinite loop is detected.
func LargestPathValue(nodes []rune, edges []Edge) (int, error) {
	graph := buildAdjacencyMatrix(edges)
	largest := 0
	for from := range graph {
		visited := make([]bool, len(nodes))
		counts := make(map[rune]int)
		stack := make([]int, 0, len(nodes))
		stack = append(stack, from)
		for len(stack) != 0 {
			var v int
			v, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if !visited[v] {
				visited[v] = true
				counts[nodes[v]]++
				for k := range graph[v] {
					stack = append(stack, k)
				}
				for _, count := range counts {
					if count > largest {
						largest = count
					}
				}
			} else {
				return 0, ErrInfiniteLoop() // cycle found
			}
		}
	}
	return largest, nil
}

func buildAdjacencyMatrix(edges []Edge) AdjacencyMatrix {
	graph := make(AdjacencyMatrix)
	for _, edge := range edges {
		if _, found := graph[edge.From]; found {
			graph[edge.From][edge.To] = struct{}{}
		} else {
			graph[edge.From] = make(map[int]struct{})
			graph[edge.From][edge.To] = struct{}{}
		}
	}
	return graph
}
