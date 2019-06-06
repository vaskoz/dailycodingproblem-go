package day287

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	visits   []WebsiteVisitor
	k        int
	expected []WebsitePair
}{
	{[]WebsiteVisitor{
		{"a", 1}, {"a", 3}, {"a", 5},
		{"b", 2}, {"b", 6},
		{"c", 1}, {"c", 2}, {"c", 3}, {"c", 4}, {"c", 5},
		{"d", 4}, {"d", 5}, {"d", 6}, {"d", 7},
		{"e", 1}, {"e", 3}, {"e", 5}, {"e", 6},
	},
		1,
		[]WebsitePair{
			{"a", "e", 0.75},
		},
	},
	{
		[]WebsiteVisitor{{"a", 1}},
		1,
		nil,
	},
	{[]WebsiteVisitor{
		{"a", 1}, {"a", 3}, {"a", 5},
		{"b", 2}, {"b", 6},
		{"c", 1}, {"c", 2}, {"c", 3}, {"c", 4}, {"c", 5},
		{"d", 4}, {"d", 5}, {"d", 6}, {"d", 7},
		{"e", 1}, {"e", 3}, {"e", 5}, {"e", 6},
	},
		3,
		[]WebsitePair{
			{"c", "e", 0.5}, {"a", "c", 0.6}, {"a", "e", 0.75},
		},
	},
}

func TestTopKMostSimilarPairs(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := TopKMostSimilarPairs(tc.visits, tc.k); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkTopKMostSimilarPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			TopKMostSimilarPairs(tc.visits, tc.k)
		}
	}
}
