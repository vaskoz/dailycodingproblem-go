package day218

// AdjacencyMatrix is a directed graph adjacency matrix.
type AdjacencyMatrix map[rune]map[rune]struct{}

// ReverseDirectedGraph reverses all the edges in a given
// directed graph.
// Runs in O(N) time and produces a new graph using O(N) space.
func ReverseDirectedGraph(g AdjacencyMatrix) AdjacencyMatrix {
	reversed := make(AdjacencyMatrix, len(g))
	for a := range g {
		for b := range g[a] {
			if _, found := reversed[b]; !found {
				reversed[b] = make(map[rune]struct{})
			}
			reversed[b][a] = struct{}{}
		}
	}
	return reversed
}
