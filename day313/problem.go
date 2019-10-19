package day313

// MovesToUnlock returns the minimum number of moves to unlock a rotary
// lock avoid all dead ends.
// Returns -1 if it's not possible.
// Panics if the input data is inconsistent.
func MovesToUnlock(start, target [3]int, deadEnds [][3]int) int {
	if start == target {
		return 0
	}

	for _, v := range start {
		if v < 0 || v > 9 {
			panic("start values must all be 0-9")
		}
	}

	for _, v := range target {
		if v < 0 || v > 9 {
			panic("target values must all be 0-9")
		}
	}

	graph := make([][][]int, 10)
	for i := range graph {
		graph[i] = make([][]int, 10)
		for j := range graph[i] {
			graph[i][j] = make([]int, 10)
		}
	}

	for _, de := range deadEnds {
		for _, v := range de {
			if v < 0 || v > 9 {
				panic("dead end values must all be 0-9")
			}
		}

		x, y, z := de[0], de[1], de[2]
		graph[x][y][z] = -1
	}

	bfs(graph, start)

	return graph[target[0]][target[1]][target[2]] - 1
}

func bfs(g [][][]int, start [3]int) {
	var q [][3]int

	g[start[0]][start[1]][start[2]] = 1

	q = append(q, start)

	for len(q) != 0 {
		var v [3]int
		v, q = q[0], q[1:]
		dist := g[v[0]][v[1]][v[2]]

		for _, dir := range [][3]int{
			{-1, 0, 0}, {0, -1, 0}, {0, 0, -1},
			{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
		} {
			next := coord(v, dir)
			x, y, z := next[0], next[1], next[2]

			if g[x][y][z] == 0 {
				g[x][y][z] = dist + 1

				q = append(q, next)
			}
		}
	}
}

func coord(start, direction [3]int) [3]int {
	var result [3]int
	for i := range result {
		result[i] = start[i] + direction[i]
		if result[i] > 9 {
			result[i] = 0
		} else if result[i] < 0 {
			result[i] = 9
		}
	}

	return result
}
