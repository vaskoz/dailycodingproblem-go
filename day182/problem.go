package day182

// Node represents a node in an undirected graph.
type Node string

// UndirectedGraph is an undirected graph of nodes.
// Edges go in both directions since it's undirected.
type UndirectedGraph map[Node]map[Node]struct{}

// IsMinimallyConnected answers if the graph is minimally connected.
// Basically runs DFS.
func IsMinimallyConnected(g UndirectedGraph) bool {
	visited := make(map[Node]struct{})
	if edges := countEdges(g); edges != (len(g)-1)*2 {
		return false
	}
	start := randomNode(g)
	isConnected(g, start, visited)
	return len(visited) == len(g)
}

func randomNode(g UndirectedGraph) Node {
	var n Node
	for x := range g {
		n = x
		break
	}
	return n
}

func isConnected(g UndirectedGraph, start Node, visited map[Node]struct{}) {
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
