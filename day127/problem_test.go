package day127

import (
	"strconv"
	"strings"
	"testing"
)

var testcases = []struct {
	first, second, sum *RevLLNum
}{
	{&RevLLNum{9, &RevLLNum{9, nil}},
		&RevLLNum{5, &RevLLNum{2, nil}},
		&RevLLNum{4, &RevLLNum{2, &RevLLNum{1, nil}}}},
	{&RevLLNum{9, &RevLLNum{9, nil}},
		nil,
		&RevLLNum{9, &RevLLNum{9, nil}}},
	{nil,
		&RevLLNum{9, &RevLLNum{9, nil}},
		&RevLLNum{9, &RevLLNum{9, nil}}},
	{&RevLLNum{9, &RevLLNum{9, &RevLLNum{9, &RevLLNum{9, nil}}}},
		&RevLLNum{1, nil},
		&RevLLNum{0, &RevLLNum{0, &RevLLNum{0, &RevLLNum{0, &RevLLNum{1, nil}}}}}},
}

func TestSumRevLLNum(t *testing.T) {
	t.Parallel()
	for tcid, tc := range testcases {
		if result := SumRevLLNum(tc.first, tc.second); !equal(result, tc.sum) {
			t.Errorf("TCID%d does not match expected. Got %v expected %v",
				tcid, String(result), String(tc.sum))
		}
	}
}

func BenchmarkSumRevLLNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			SumRevLLNum(tc.first, tc.second)
		}
	}
}

func equal(a, b *RevLLNum) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil && b != nil {
		return false
	} else if a != nil && b == nil {
		return false
	}
	for a != nil && b != nil {
		if a.Digit != b.Digit {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == b
}

func String(a *RevLLNum) string {
	var sb strings.Builder
	for a != nil {
		sb.WriteString(strconv.Itoa(a.Digit))
		a = a.Next
	}
	return sb.String()
}
