package day329

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	males, females map[string][]string
	matches        Matches
}{
	{
		map[string][]string{
			"andrew":  {"caroline", "abigail", "betty"},
			"bill":    {"caroline", "betty", "abigail"},
			"chester": {"betty", "caroline", "abigail"},
		},
		map[string][]string{
			"abigail":  {"andrew", "bill", "chester"},
			"betty":    {"bill", "andrew", "chester"},
			"caroline": {"bill", "chester", "andrew"},
		},
		Matches{
			"andrew":  "abigail",
			"bill":    "caroline",
			"chester": "betty",
		},
	},
	{
		map[string][]string{
			"0": {"7", "5", "6", "4"},
			"1": {"5", "4", "6", "7"},
			"2": {"4", "5", "6", "7"},
			"3": {"4", "5", "6", "7"},
		},
		map[string][]string{
			"4": {"0", "1", "2", "3"},
			"5": {"0", "1", "2", "3"},
			"6": {"0", "1", "2", "3"},
			"7": {"0", "1", "2", "3"},
		},
		Matches{
			"2": "4",
			"1": "5",
			"3": "6",
			"0": "7",
		},
	},
	{
		map[string][]string{
			"0": {"2", "3"},
			"1": {"3", "2"},
		},
		map[string][]string{
			"2": {"1", "0"},
			"3": {"1", "0"},
		},
		Matches{
			"0": "2",
			"1": "3",
		},
	},
}

func TestStableMarriage(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if res := StableMarriage(tc.males, tc.females); !reflect.DeepEqual(res, tc.matches) {
			t.Errorf("Expected %v, got %v", tc.matches, res)
		}
	}
}

func BenchmarkStableMarriage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			StableMarriage(tc.males, tc.females)
		}
	}
}
