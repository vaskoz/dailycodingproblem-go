package day299

import "github.com/algds/kruskals"

type Pipe kruskals.WeightedEdge

// LowestCostPipeConfiguration is a MST for the given pipes.
func LowestCostPipeConfiguration(pipes []Pipe) []Pipe {
	input := make([]kruskals.WeightedEdge, len(pipes))
	for i := range input {
		input[i] = kruskals.WeightedEdge(pipes[i])
	}
	answer := kruskals.MinimumSpanningTree(input)
	result := make([]Pipe, len(answer))
	for i := range result {
		result[i] = Pipe(answer[i])
	}
	return result
}
