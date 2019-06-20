package day299

import "github.com/algds/kruskals"

// LowestCostPipeConfiguration is a MST for the given pipes.
func LowestCostPipeConfiguration(pipes []kruskals.WeightedEdge) []kruskals.WeightedEdge {
	return kruskals.MinimumSpanningTree(pipes)
}
