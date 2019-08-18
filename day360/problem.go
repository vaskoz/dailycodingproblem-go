package day360

import "errors"

// Song represents a particular song.
type Song int

// Playlist is a ranked priority list of songs.
type Playlist []Song

// ErrPlaylistNotPossible is the error returned
// if a playlist is not possible respecting the ranks.
func ErrPlaylistNotPossible() error {
	return errPlaylistNotPossible
}

var errPlaylistNotPossible = errors.New("not possible to produce a playlist with constraints")

// InterleaveSongsByRank uses a topological sort
// to return a satisfactory playlist or an error if
// not possible.
func InterleaveSongsByRank(ranks []Playlist) ([]Song, error) {
	g := make(directedGraph)
	for _, rank := range ranks {
		for i := 1; i < len(rank); i++ {
			if g[rank[i-1]] == nil {
				g[rank[i-1]] = make(map[node]struct{})
			}
			g[rank[i-1]][rank[i]] = struct{}{}
		}
	}
	nodes, err := dfsTopo(g)
	if err != nil {
		return nil, errPlaylistNotPossible
	}
	songs := make([]Song, len(nodes))
	for i := range nodes {
		songs[i] = nodes[i].(Song)
	}
	return songs, nil
}

type node interface{}

type directedGraph map[node]map[node]struct{}

func dfsTopo(g directedGraph) ([]node, error) {
	var result []node
	marks := make(map[node]int)
	for from, to := range g {
		marks[from] = 0
		for t := range to {
			marks[t] = 0
		}
	}
	for n, mark := range marks {
		if mark == 0 { // unmarked
			tmp := make([]node, 0, len(marks))
			nodes, err := visit(n, g, marks)
			if err == nil {
				tmp = append(tmp, nodes...)
				tmp = append(tmp, result...)
				result = tmp
			} else {
				return nil, err
			}
		}
	}
	return result, nil
}

func visit(n node, g directedGraph, marks map[node]int) ([]node, error) {
	if marks[n] == 2 {
		return nil, nil
	} else if marks[n] == 1 {
		return nil, errPlaylistNotPossible
	}
	marks[n] = 1
	result := make([]node, 0, len(marks))
	for m := range g[n] {
		tmp := make([]node, 0, len(marks))
		nodes, err := visit(m, g, marks)
		if err == nil {
			tmp = append(tmp, nodes...)
			tmp = append(tmp, result...)
			result = tmp
		} else {
			return nil, err
		}
	}
	marks[n] = 2
	tmp := make([]node, 0, len(marks))
	tmp = append(tmp, n)
	tmp = append(tmp, result...)
	result = tmp
	return result, nil
}
