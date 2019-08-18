package day360

import "errors"

import "github.com/algds/dfstopo"

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
	dg := make(dfstopo.DirectedGraph)
	for _, rank := range ranks {
		for i := 1; i < len(rank); i++ {
			if dg[rank[i-1]] == nil {
				dg[rank[i-1]] = make(map[dfstopo.Node]struct{})
			}
			dg[rank[i-1]][rank[i]] = struct{}{}
		}
	}
	nodes, err := dfstopo.Sort(dg)
	if err != nil {
		return nil, errPlaylistNotPossible
	}
	songs := make([]Song, len(nodes))
	for i := range nodes {
		songs[i] = nodes[i].(Song)
	}
	return songs, nil
}
