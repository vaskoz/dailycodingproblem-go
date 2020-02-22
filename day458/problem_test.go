package day458

import "testing"

// nolint
var testcases = []struct {
	rules   []string
	isValid bool
}{
	{[]string{"A N B", "B NE C", "C N A"}, false},
	{[]string{"A NW B", "A N B"}, true},
	{[]string{"A N B", "C N A", "A S C"}, true},
	{[]string{"A NW B", "C NE A", "A S C"}, true},
	{[]string{"A NW B", "C NE A", "A SE C"}, false},
	{[]string{"A NW B", "C NE A", "A SW C"}, true},
	{[]string{"A NW B", "C NE A", "A W C"}, true},
	{[]string{"A NW B", "C NE A", "A E C"}, false},
	{[]string{"A NW B", "C NE A", "B S C"}, true},
	{[]string{"A NW B", "C NE A", "C E B"}, false},
	{[]string{"A NW B", "C NE A", "B SE A"}, true},
	{[]string{"A NW B", "C NE A", "A SE B"}, false},
}

func TestIsValidRules(t *testing.T) {
	t.Parallel()

	for tcid, tc := range testcases {
		if result := IsValidRules(tc.rules); result != tc.isValid {
			t.Errorf("TC%d Expected %v got %v", tcid, tc.isValid, result)
		}
	}
}

func BenchmarkIsValidRules(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			IsValidRules(tc.rules)
		}
	}
}
