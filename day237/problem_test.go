package day237

import "testing"

var testcases = []struct {
	root      *KaryTree
	symmetric bool
}{
	{&KaryTree{
		4,
		[]*KaryTree{
			{
				3,
				[]*KaryTree{{9, nil}, nil, nil},
			},
			{
				5,
				[]*KaryTree{nil, nil, nil},
			},
			{
				3,
				[]*KaryTree{nil, nil, {9, nil}},
			},
		},
	}, true},
	{nil, true},
	{&KaryTree{
		4,
		[]*KaryTree{
			{
				3,
				[]*KaryTree{nil, {9, nil}, nil},
			},
			{
				5,
				[]*KaryTree{nil, nil, nil},
			},
			{
				3,
				[]*KaryTree{nil, nil, {9, nil}},
			},
		},
	}, false},
}

func TestIsTreeSymmetric(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := IsTreeSymmetric(tc.root); result != tc.symmetric {
			t.Errorf("Expected %v, got %v", tc.symmetric, result)
		}
	}
}

func BenchmarkIsTreeSymmetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsTreeSymmetric(tc.root)
		}
	}
}
