package day182

// UndirectedGraph is an undirected graph of nodes.
// Edges go in both directions since it's undirected.
type UndirectedGraph map[string]map[string]struct{}

// IsMinimallyConnected answers if the graph is minimally connected.
// Basically runs DFS.
func IsMinimallyConnected(g UndirectedGraph) bool {
	visited := make(map[string]struct{})
	if edges := countEdges(g); edges != (len(g)-1)*2 {
		return false
	}
	start := randomNode(g)
	isConnected(g, start, visited)
	return len(visited) == len(g)
}

func randomNode(g UndirectedGraph) string {
	var n string
	for x := range g {
		n = x
		break
	}
	return n
}

func isConnected(g UndirectedGraph, start string, visited map[string]struct{}) {
	visited[start] = struct{}{}
	for dest := range g[start] {
		if _, found := visited[dest]; !found {
			isConnected(g, dest, visited)
		}
	}
}

func countEdges(g UndirectedGraph) int {
	count := 0
	for _, targets := range g {
		for range targets {
			count++
		}
	}
	return count
}
