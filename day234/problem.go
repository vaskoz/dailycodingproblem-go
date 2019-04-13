package day234

import (
	"github.com/algds/kruskals"
)

// MaximumSpanningTree returns the subset of edges
// representing the maximum spanning tree.
// Just Kruskal's algorithm with negated weights.
func MaximumSpanningTree(edges []kruskals.SimpleWeightedEdge) []kruskals.SimpleWeightedEdge {
	copied := make([]kruskals.WeightedEdge, len(edges))
	for i, e := range edges {
		copied[i] = kruskals.SimpleWeightedEdge{F: e.F, T: e.T, W: -e.W}
	}
	mst := kruskals.MinimumSpanningTree(copied)
	result := make([]kruskals.SimpleWeightedEdge, len(mst))
	for i, e := range mst {
		result[i] = kruskals.SimpleWeightedEdge{F: e.From(), T: e.To(), W: -e.Weight()}
	}
	return result
}
