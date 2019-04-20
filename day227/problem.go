package day227

// Board is an NxN board of runes.
type Board [][]rune

type trie struct {
	word    bool
	letters map[rune]*trie
}

// BoggleSolver takes a boggle board (2D runes NxN) and a dictionary
// of valid words.
func BoggleSolver(board Board, dict []string) []string {
	t := buildTrie(dict)
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	result := make([]string, 0, len(dict))
	for i := range board {
		for j := range board[i] {
			result = append(result, boggleSolver(board, visited, t, []rune{}, i, j)...)
		}
	}
	// dedup
	m := make(map[string]struct{})
	for _, w := range result {
		m[w] = struct{}{}
	}
	result = make([]string, 0, len(m))
	for w := range m {
		result = append(result, w)
	}
	return result
}

func boggleSolver(board Board, visited [][]bool, t *trie, path []rune, i, j int) []string {
	if i < 0 || i >= len(board) ||
		j < 0 || j >= len(board[i]) ||
		visited[i][j] || t.letters[board[i][j]] == nil {
		return nil
	}
	var result []string
	path = append(path, board[i][j])
	t = t.letters[board[i][j]]
	if t.word {
		result = append(result, string(path))
	}
	visited[i][j] = true
	// visit 8 directions: up,down,left,right and the 4 diagonals
	result = append(result, boggleSolver(board, visited, t, path, i-1, j)...)   // up
	result = append(result, boggleSolver(board, visited, t, path, i+1, j)...)   // down
	result = append(result, boggleSolver(board, visited, t, path, i, j-1)...)   // left
	result = append(result, boggleSolver(board, visited, t, path, i, j+1)...)   // right
	result = append(result, boggleSolver(board, visited, t, path, i-1, j-1)...) // up-left
	result = append(result, boggleSolver(board, visited, t, path, i-1, j+1)...) // up-right
	result = append(result, boggleSolver(board, visited, t, path, i+1, j-1)...) // down-left
	result = append(result, boggleSolver(board, visited, t, path, i+1, j+1)...) // down-right
	visited[i][j] = false
	return result
}

func buildTrie(dict []string) *trie {
	root := &trie{false, make(map[rune]*trie)}
	for _, word := range dict {
		curr := root
		for _, r := range word {
			if next, found := curr.letters[r]; found {
				curr = next
			} else {
				next = &trie{false, make(map[rune]*trie)}
				curr.letters[r] = next
				curr = next
			}
		}
		curr.word = true
	}
	return root
}
