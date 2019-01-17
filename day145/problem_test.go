package day145

import "testing"

var testcases = []struct {
	input, expected *LL
}{
	{&LL{1, &LL{2, &LL{3, &LL{4, nil}}}},
		&LL{2, &LL{1, &LL{4, &LL{3, nil}}}},
	},
	{&LL{1, &LL{2, &LL{3, &LL{4, &LL{5, nil}}}}},
		&LL{2, &LL{1, &LL{4, &LL{3, &LL{5, nil}}}}},
	},
}

func TestSwapEveryTwo(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := SwapEveryTwo(tc.input); !equal(result, tc.expected) {
			t.Errorf("TCID%d do not match", tcid)
		}
	}
}

func BenchmarkSwapEveryTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SwapEveryTwo(tc.input)
		}
	}
}

func equal(one, two *LL) bool {
	for one != nil && two != nil {
		if one.Value != two.Value {
			return false
		}
		one = one.Next
		two = two.Next
	}
	return one == nil && two == nil
}
