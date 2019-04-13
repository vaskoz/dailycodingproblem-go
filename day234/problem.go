package day234

import "sort"

import "github.com/algds/uf"

// Edge represents a weighted edge in a graph.
type Edge struct {
	From, To int
	Weight   int
}

// MaximumSpanningTree returns the subset of edges
// representing the maximum spanning tree.
// Just Kruskal's algorithm with negated weights.
func MaximumSpanningTree(edges []Edge) []Edge {
	copied := make([]Edge, len(edges))
	vertices := make(map[int]struct{})
	for i, e := range edges {
		copied[i] = e
		copied[i].Weight = -copied[i].Weight
		vertices[e.From] = struct{}{}
		vertices[e.To] = struct{}{}
	}
	sort.Slice(copied, func(i, j int) bool {
		return copied[i].Weight < copied[j].Weight
	})
	u := uf.New(len(vertices))
	result := make([]Edge, 0, len(edges))
	for _, e := range copied {
		if !u.Connected(e.From, e.To) {
			u.Union(e.From, e.To)
			e.Weight = -e.Weight
			result = append(result, e)
		}
	}
	return result
}
