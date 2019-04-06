package day226

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	alienWords []string
	order      []rune
}{
	{[]string{"xww", "wxyz", "wxyw", "ywx", "ywz"}, []rune{'x', 'z', 'w', 'y'}},
}

func TestAlienLanguageOrder(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if order := AlienLanguageOrder(tc.alienWords); !reflect.DeepEqual(order, tc.order) {
			t.Errorf("Expected %v, got %v", tc.order, order)
		}
	}
}

func BenchmarkAlienLanguageOrder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			AlienLanguageOrder(tc.alienWords)
		}
	}
}
