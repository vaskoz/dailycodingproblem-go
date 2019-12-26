package day362

import (
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	digits   int
	expected []string
}{
	{0, nil},
	{1, []string{"0", "1", "8"}},
	{2, []string{"11", "69", "88", "96"}},
	{3, []string{"101", "111", "181", "609", "619", "689", "808", "818", "888", "906", "916", "986"}},
	{4, []string{"1001", "1111", "1691", "1881", "1961", "6009", "6119", "6699", "6889", "6969", "8008",
		"8118", "8698", "8888", "8968", "9006", "9116", "9696", "9886", "9966"}},
}

func TestStrobogrammaticNumber(t *testing.T) {
	t.Parallel()

	for _, tc := range testcases {
		if result := StrobogrammaticNumber(tc.digits); !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}

func BenchmarkStrobogrammaticNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			StrobogrammaticNumber(tc.digits)
		}
	}
}
