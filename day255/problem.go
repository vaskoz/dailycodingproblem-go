package day255

// Graph is an integer based ajacency matrix.
type Graph map[int][]int

// TransitiveClosure returns the transitive closure
// matrix.
func TransitiveClosure(g Graph) [][]int {
	var largest int
	for start, ends := range g {
		if start > largest {
			largest = start
		}
		for _, node := range ends {
			if node > largest {
				largest = node
			}
		}
	}
	result := make([][]int, largest+1)
	for i := range result {
		result[i] = make([]int, largest+1)
	}
	visited := make(map[int]struct{}, largest+1)
	result = transitiveClosure(-1, g, result, visited)
	return result
}

func transitiveClosure(node int, g Graph, result [][]int, visited map[int]struct{}) [][]int {
	if node == -1 {
		for start := range g {
			result = transitiveClosure(start, g, result, visited)
		}
	} else {
		for start := range visited {
			result[start][node] = 1
		}
		visited[node] = struct{}{}
		for _, n := range g[node] {
			if n == node {
				result[node][n] = 1
			} else if _, found := visited[n]; !found {
				result = transitiveClosure(n, g, result, visited)
			}
		}
		delete(visited, node)
	}
	return result
}
