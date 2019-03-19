package day207

type color int

const (
	blue color = iota
	green
)

// Node is a node in the graph.
type Node string

// UndirectedGraph has an edge going in each direction.
type UndirectedGraph map[Node]map[Node]struct{}

// IsBipartiteGraph answer if an UndirectedGraph
// is actually bipartite.
func IsBipartiteGraph(g UndirectedGraph) bool {
	colors := make(map[Node]color, len(g))
	var n Node
	for n1 := range g {
		n = n1
		break
	}
	result := isBipartiteGraph(g, n, colors)
	return result && len(colors) == len(g)
}

func isBipartiteGraph(g UndirectedGraph, n Node, colors map[Node]color) bool {
	if len(g) == len(colors) {
		return true
	}
	nextColor := blue
	if colors[n] == blue {
		nextColor = green
	}
	for node := range g[n] {
		if c, found := colors[node]; !found {
			colors[node] = nextColor
			if !isBipartiteGraph(g, node, colors) {
				return false
			}
		} else if c != nextColor {
			return false
		}
	}
	return true
}
