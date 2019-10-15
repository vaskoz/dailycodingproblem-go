package day337

import (
	"math/rand"
	"reflect"
	"testing"
)

// nolint
var testcases = []struct {
	input                          []interface{}
	expectedFastShuffle            []interface{}
	expectedMemoryOptimizedShuffle []interface{}
}{
	{
		[]interface{}{"a", 1, 3.5, "foobar", nil, nil, nil},
		[]interface{}{"foobar", 1, nil, "a", nil, nil, 3.5},
		[]interface{}{"a", 1, "foobar", nil, nil, nil, 3.5},
	},
}

func TestFastShuffle(t *testing.T) {
	// NOTE: Don't run tests in parallel. Setting a global variable is a race.
	// t.Parallel()
	r = rand.New(rand.NewSource(314159))

	for _, tc := range testcases {
		head := convertToLL(tc.input)
		ll := FastShuffle(head)
		result := convertToSlice(ll)

		if !reflect.DeepEqual(result, tc.expectedFastShuffle) {
			t.Errorf("Expected %v, got %v", tc.expectedFastShuffle, result)
		}
	}
}

func TestMemoryOptimizedShuffle(t *testing.T) {
	// NOTE: Don't run tests in parallel. Setting a global variable is a race.
	// t.Parallel()
	r = rand.New(rand.NewSource(314159))

	for _, tc := range testcases {
		head := convertToLL(tc.input)
		ll := MemoryOptimizedShuffle(head)
		result := convertToSlice(ll)

		if !reflect.DeepEqual(result, tc.expectedMemoryOptimizedShuffle) {
			t.Errorf("Expected %v, got %v", tc.expectedMemoryOptimizedShuffle, result)
		}
	}
}

func BenchmarkFastShuffle(b *testing.B) {
	large := make([]interface{}, 10_000)

	for i := range large {
		large[i] = i
	}

	head := convertToLL(large)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		head = FastShuffle(head)
	}
}

func BenchmarkMemoryOptimizedShuffle(b *testing.B) {
	large := make([]interface{}, 10_000)
	for i := range large {
		large[i] = i
	}

	head := convertToLL(large)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		head = MemoryOptimizedShuffle(head)
	}
}
