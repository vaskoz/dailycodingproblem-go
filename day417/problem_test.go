package day417

import "testing"

// nolint
var testcases = []struct {
	input, expected *LL
}{
	{
		&LL{3, &LL{4, &LL{-7, &LL{5, &LL{-6, &LL{6, nil}}}}}},
		&LL{5, nil},
	},
	{
		&LL{3, &LL{4, &LL{5, &LL{-12, nil}}}},
		nil,
	},
	{
		&LL{3, &LL{4, &LL{-7, &LL{5, &LL{-5, &LL{6, nil}}}}}},
		&LL{6, nil},
	},
	{
		&LL{3, &LL{4, &LL{-7, &LL{3, &LL{5, &LL{-6, &LL{6, &LL{-5, nil}}}}}}}},
		&LL{3, nil},
	},
	{
		&LL{1, &LL{1, &LL{2, &LL{3, &LL{4, &LL{5, &LL{6, &LL{-5, nil}}}}}}}},
		&LL{1, &LL{1, &LL{2, &LL{3, &LL{4, &LL{5, &LL{6, &LL{-5, nil}}}}}}}},
	},
}

func equal(a, b *LL) bool {
	for a != nil && b != nil {
		if a.Value != b.Value {
			return false
		}

		a = a.Next
		b = b.Next
	}

	return a == nil && b == nil
}

func TestRemoveConsecutiveSumZero(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result := RemoveConsecutiveSumZero(tc.input); !equal(result, tc.expected) {
			t.Errorf("TCID%d doesn't match", tcid)
		}
	}
}

func BenchmarkRemoveConsecutiveSumZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			RemoveConsecutiveSumZero(tc.input)
		}
	}
}
