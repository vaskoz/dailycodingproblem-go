package day56

// Graph represents an adjacency list.
type Graph map[int]map[int]struct{}

// CanColor returns true if the graph can be colored with at
// most k colors.
func CanColor(v int, g Graph, color []int, k int) bool {
	if v == len(g) {
		return true
	}
	for c := 1; c <= k; c++ {
		if isSafe(v, g, color, c) {
			color[v] = c
			if CanColor(v+1, g, color, k) {
				return true
			}
			color[v] = 0
		}
	}
	return false
}

func isSafe(v int, g Graph, color []int, c int) bool {
	for i := 0; i < len(g); i++ {
		if _, found := g[v][i]; found && c == color[i] {
			return false
		}
	}
	return true
}
