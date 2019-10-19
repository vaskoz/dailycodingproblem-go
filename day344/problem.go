package day344

// Tree is an undirected tree.
type Tree map[int]map[int]struct{}

// MaxRemovedEdgesEvenSubtrees returns the maximum number of edges
// that can be removed to result in subtrees that all have even
// number of nodes.
func MaxRemovedEdgesEvenSubtrees(t Tree) int {
	nt := make(Tree)

	var firstStart int
	for start, toMap := range t {
		firstStart = start

		if _, exists := nt[start]; !exists {
			nt[start] = make(map[int]struct{})
		}

		for to := range toMap {
			if _, exists := nt[to]; !exists {
				nt[to] = make(map[int]struct{})
			}

			nt[start][to] = struct{}{}
			nt[to][start] = struct{}{}
		}
	}

	visited := make(map[int]struct{}, len(nt)+1)
	_, result := dfs(nt, firstStart, visited, 0)

	return result
}

func dfs(t Tree, start int, visited map[int]struct{}, result int) (int, int) {
	var nodes, subtreeCount int

	visited[start] = struct{}{}

	for to := range t[start] {
		if _, seen := visited[to]; !seen {
			subtreeCount, result = dfs(t, to, visited, result)
			if subtreeCount%2 == 0 {
				result++
			} else {
				nodes += subtreeCount
			}
		}
	}

	return nodes + 1, result
}
