package day333

import "testing"

// nolint
var testcases = []struct {
	knows  *Knowing
	people []interface{}
	celeb  int
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
		[]interface{}{"a", "b", "c", "d", "e"},
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
		[]interface{}{"a", "b", "c", "d", "e"},
		-1,
	},
}

func TestFindCelebrity(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if celeb := FindCelebrity(tc.knows, tc.people); celeb != tc.celeb {
			t.Errorf("Expected %v, got %v", tc.celeb, celeb)
		}
	}
}
