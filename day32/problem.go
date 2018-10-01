package day32

import "math"

// Arbitrage calculates how much money is possibly made using arbitrage.
// Assuming no transactional costs.
// Uses Bellman-Ford Shortest Path which runs in O(N^3) time and O(N^2) space
// due to the transformed rates table.
func Arbitrage(rates [][]float64) bool {
	graph := make([][]float64, len(rates))
	for i := range graph {
		graph[i] = make([]float64, len(rates[i]))
		for j := range graph[i] {
			graph[i][j] = -math.Log(rates[i][j])
		}
	}
	source := 0
	minDist := make([]float64, len(rates))
	for i := range minDist {
		minDist[i] = math.Inf(1)
	}
	minDist[source] = 0

	for i := 0; i < len(graph)-1; i++ {
		for v := range graph {
			for w := range graph[v] {
				if minDist[w] > minDist[v]+graph[v][w] {
					minDist[w] = minDist[v] + graph[v][w]
				}
			}
		}
	}
	for v := range graph {
		for w := range graph[v] {
			if minDist[w] > minDist[v]+graph[v][w] {
				return true
			}
		}
	}
	return false
}
