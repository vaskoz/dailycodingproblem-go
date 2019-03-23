package day213

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	digits string
	ips    []string
}{
	{"2542540123", []string{"254.25.40.123", "254.254.0.123"}},
}

func TestValidIPAddresses(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := ValidIPAddresses(tc.digits); !reflect.DeepEqual(tc.ips, result) {
			t.Errorf("Expected %v, got %v", tc.ips, result)
		}
	}
}

func BenchmarkValidIPAddresses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			ValidIPAddresses(tc.digits)
		}
	}
}
