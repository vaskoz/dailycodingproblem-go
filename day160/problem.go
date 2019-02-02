package day160

// WeightedEdges is the weight of an undirected edge.
type WeightedEdges struct {
	start, end string
	weight     int
}

// Graph is an undirected weighted graph.
type Graph map[string]map[string]int

// LongestPath returns the weight of the longest path
// in an undirected tree.
func LongestPath(weights []WeightedEdges) int {
	g, start := buildAdjacencyMatrix(weights)
	if len(g) == 0 {
		return 0
	}
	for n := range g {
		start = n
		break
	}
	newStart, _ := farthestNode(start, g)
	_, dist := farthestNode(newStart, g)
	return dist
}

func farthestNode(start string, g Graph) (string, int) {
	var farthest string
	var dist int
	visited := make(map[string]int)
	visited[start] = 0
	n := len(g)
	queue := make([]string, 0, n)
	queue = append(queue, start)
	for len(queue) != 0 {
		nextQ := make([]string, 0, n)
		for _, curr := range queue {
			for node, d := range g[curr] {
				if _, seen := visited[node]; !seen {
					visited[node] = visited[curr] + d
					if visited[node] > dist {
						farthest = node
						dist = visited[node]
					}
					nextQ = append(nextQ, node)
				}
			}
		}
		queue = nextQ
	}
	return farthest, dist
}

func buildAdjacencyMatrix(weights []WeightedEdges) (Graph, string) {
	if len(weights) == 0 {
		return nil, ""
	}
	g := make(Graph)
	for _, w := range weights {
		if g[w.start] == nil {
			g[w.start] = make(map[string]int)
		}
		if g[w.end] == nil {
			g[w.end] = make(map[string]int)
		}
		g[w.start][w.end] = w.weight
		g[w.end][w.start] = w.weight
	}
	return g, weights[0].start
}
