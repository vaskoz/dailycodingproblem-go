package day397

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	jobs, expected []Job
}{
	{
		[]Job{
			{0, 6},
			{1, 4},
			{3, 5},
			{3, 8},
			{4, 7},
			{5, 9},
			{6, 10},
			{8, 11},
		},
		[]Job{
			{1, 4},
			{4, 7},
			{8, 11},
		},
	},
	{
		[]Job{
			{0, 100},
			{1, 3},
			{3, 5},
			{0, 10},
			{6, 8},
			{9, 10},
			{6, 10},
			{8, 11},
		},
		[]Job{
			{1, 3},
			{3, 5},
			{6, 8},
			{9, 10},
		},
	},
}

func TestLargestSubsetCompatibleJobs(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := LargestSubsetCompatibleJobs(tc.jobs); !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, res)
		}
	}
}

func BenchmarkLargestSubsetCompatibleJobs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			LargestSubsetCompatibleJobs(tc.jobs)
		}
	}
}
