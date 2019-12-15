package day391

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	user1, user2 []string
	expected     []string
}{
	{
		[]string{"/home", "/register", "/login", "/user", "/one", "/two"},
		[]string{"/home", "/red", "/login", "/user", "/one", "/pink"},
		[]string{"/login", "/user", "/one"},
	},
	{
		[]string{"/a", "/b", "/a", "/b", "/c", "/d"},
		[]string{"/e", "/a", "/b", "/c", "/d", "/e"},
		[]string{"/a", "/b", "/c", "/d"},
	},
}

func TestLongestContiguousBrowsingHistory(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := LongestContiguousBrowsingHistory(tc.user1, tc.user2); !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkLongestContiguousBrowsingHistory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LongestContiguousBrowsingHistory(tc.user1, tc.user2)
		}
	}
}
