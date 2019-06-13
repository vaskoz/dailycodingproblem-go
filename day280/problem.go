package day280

// UndirectedGraph is a connected undirected adjacency list.
type UndirectedGraph map[int]map[int]struct{}

// HasCycle answers if the connected undirected graph is cyclic.
func HasCycle(g UndirectedGraph, start int) bool {
	visits := make(map[int]bool)
	return hasCycle(g, start, visits)
}

func hasCycle(g UndirectedGraph, start int, visits map[int]bool) bool {
	visits[start] = true
	for node := range g[start] {
		if visited := visits[node]; visited || hasCycle(g, node, visits) {
			return true
		}
	}
	return false
}
