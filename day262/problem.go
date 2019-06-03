package day262

// Edge is an undirected edge from a Start node to an End node.
type Edge struct {
	Start, End int
}

// undirectedGraph is an undirected graph.
// Edges will be inserted as 2-directed edges in opposite directions.
type undirectedGraph map[int]map[int]struct{}

// FindAllBridges returns all bridges in a given UndirectedGraph.
func FindAllBridges(edges []Edge) []Edge {
	graph := make(undirectedGraph)
	for _, e := range edges {
		if _, found := graph[e.Start]; !found {
			graph[e.Start] = make(map[int]struct{})
		}
		if _, found := graph[e.End]; !found {
			graph[e.End] = make(map[int]struct{})
		}
		graph[e.Start][e.End] = struct{}{}
		graph[e.End][e.Start] = struct{}{}
	}
	visited := make([]bool, len(graph))
	disc := make([]int, len(graph))
	low := make([]int, len(graph))
	parent := make([]int, len(graph))
	for i := range disc {
		disc[i] = int(^uint(0) >> 1)
		low[i] = int(^uint(0) >> 1)
		parent[i] = -1
	}
	var result []Edge
	for i := range visited {
		if !visited[i] {
			result = findAllBridges(graph, i, 0, visited, parent, low, disc, result)
		}
	}
	return result
}

func findAllBridges(graph undirectedGraph, u, time int, visited []bool, parent, low, disc []int, result []Edge) []Edge {
	visited[u] = true
	disc[u] = time
	low[u] = time
	time++
	for v := range graph[u] {
		if !visited[v] {
			parent[v] = u
			result = findAllBridges(graph, v, time, visited, parent, low, disc, result)
			low[u] = min(low[u], low[v])
			if low[v] > disc[u] {
				result = append(result, Edge{u, v})
			}
		} else if v != parent[u] {
			low[u] = min(low[u], disc[v])
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
