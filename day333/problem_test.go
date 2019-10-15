package day333

import "testing"

// nolint
var testcases = []struct {
	knows     *Knowing
	numPeople int
	celeb     int
}{
	{
		&Knowing{
			0: {
				1: true,
			},
			2: {
				1: true,
			},
			3: {
				1: true,
			},
			4: {
				1: true,
			},
		},
		5,
		1,
	},
	{
		&Knowing{
			0: {
				1: true,
			},
			1: {
				2: true,
			},
			2: {
				3: true,
			},
			3: {
				4: true,
			},
		},
		5,
		-1,
	},
}

func TestFindCelebrity(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if celeb := FindCelebrity(tc.knows, tc.numPeople); celeb != tc.celeb {
			t.Errorf("Expected %v, got %v", tc.celeb, celeb)
		}
	}
}

func BenchmarkFindCelebrity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FindCelebrity(tc.knows, tc.numPeople)
		}
	}
}
