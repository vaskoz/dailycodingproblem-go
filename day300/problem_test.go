package day300

import (
	"bytes"
	"strings"
	"testing"
)

// nolint
var testcases = []struct {
	input, output string
}{
	{"1,1 2,2 3,3 4,3 5,3", "1,1 \n1,1 2,1 \n1,1 2,1 3,1 \n1,1 2,1 3,2 \n1,1 2,1 3,3 \n"},
	{"1,1 2,2 3,3 1,3", "1,1 \n1,1 2,1 \n1,1 2,1 3,1 \nvoter 1 already voted. this vote for candidate 3 will not count\n"},
	{"1,1 2,2 3,3 4,4 5,4",
		"1,1 \n1,1 2,1 \n1,1 2,1 3,1 \n1,1 2,1 3,1 \n2,1 3,1 4,2 \n"},
}

func TestMainVoterTop3(t *testing.T) {
	// don't run in parallel due to global var
	for _, tc := range testcases {
		buff := new(bytes.Buffer)
		out = buff
		in = strings.NewReader(tc.input)
		main()
		if result := buff.String(); result != tc.output {
			t.Errorf("Expected \n%s, got \n%s", tc.output, result)
		}
	}
}

func BenchmarkMainVoterTop3(b *testing.B) {
	// don't run in parallel due to global var
	buff := new(bytes.Buffer)
	out = buff
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			in = strings.NewReader(tc.input)
			main()
		}
	}
}
