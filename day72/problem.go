package day72

import (
	"errors"
)

var errInfiniteLoop = errors.New("loop detected")

// AdjacencyMatrix represents a sparse adjacency matrix.
type AdjacencyMatrix map[int]map[int]struct{}

// Node is a character representation of a node in a graph.
type Node rune

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
func LargestPathValue(nodes []Node, edges []Edge) (int, error) {
	graph := buildAdjacencyMatrix(edges)
	visited := make(map[int]struct{}, len(nodes))
	freq := make(map[Node]int)
	var largest int
	for start := range graph {
		if count, err := dfsBacktracking(graph, nodes, start, visited, freq); err != nil {
			return 0, err
		} else if count > largest {
			largest = count
		}
	}
	return largest, nil
}

func dfsBacktracking(g AdjacencyMatrix, nodes []Node, n int, visited map[int]struct{}, freq map[Node]int) (int, error) {
	visited[n] = struct{}{}
	freq[nodes[n]]++
	var max int
	var err error
	if len(g[n]) == 0 {
		for _, count := range freq {
			if count > max {
				max = count
			}
		}
	} else {
		for next := range g[n] {
			if _, seen := visited[next]; seen {
				err = ErrInfiniteLoop()
				break
			}
			var v int
			if v, err = dfsBacktracking(g, nodes, next, visited, freq); err != nil {
				err = ErrInfiniteLoop()
				break
			}
			if v > max {
				max = v
			}
		}
	}
	freq[nodes[n]]--
	delete(visited, n)
	return max, err
}

func buildAdjacencyMatrix(edges []Edge) AdjacencyMatrix {
	graph := make(AdjacencyMatrix)
	for _, edge := range edges {
		if _, found := graph[edge.From]; !found {
			graph[edge.From] = make(map[int]struct{})
		}
		graph[edge.From][edge.To] = struct{}{}
	}
	return graph
}
