package day302

import "testing"

// nolint
var testcases = []struct {
	grid    []string
	regions int
}{
	{
		[]string{
			`\    /`,
			` \  / `,
			`  \/  `,
		},
		3,
	},
}

func TestDistinctRegions(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if regions := DistinctRegions(tc.grid); regions != tc.regions {
			t.Errorf("Expected %v, got %v", tc.regions, regions)
		}
	}
}

func BenchmarkDistinctRegions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			DistinctRegions(tc.grid)
		}
	}
}
