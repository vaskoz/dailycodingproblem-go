package day360

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	ranks       []Playlist
	expected    []Song
	expectedErr error
}{
	{
		[]Playlist{
			{1, 7, 3},
			{2, 1, 6, 7, 9},
			{3, 9, 5},
		},
		[]Song{2, 1, 6, 7, 3, 9, 5},
		nil,
	},
	{
		[]Playlist{
			{1, 7, 3},
			{3, 1, 6, 7, 9},
			{3, 9, 5},
		},
		nil,
		ErrPlaylistNotPossible(),
	},
}

func TestInterleaveSongsByRank(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result, err := InterleaveSongsByRank(tc.ranks); !reflect.DeepEqual(result, tc.expected) ||
			err != tc.expectedErr {
			t.Errorf("Expected (%v,%v), got (%v,%v)", tc.expected, tc.expectedErr, result, err)
		}
	}
}

func BenchmarkInterleaveSongsByRank(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			InterleaveSongsByRank(tc.ranks) // nolint
		}
	}
}
