package day173

import (
	"reflect"
	"testing"
)

var testcases = []struct {
	nested, flattened map[string]interface{}
}{
	{map[string]interface{}{
		"key": 3,
		"foo": map[string]interface{}{
			"a": 5,
			"bar": map[string]interface{}{
				"baz": 8,
			},
		},
	},
		map[string]interface{}{
			"key":         3,
			"foo.a":       5,
			"foo.bar.baz": 8,
		}},
}

func TestFlattenMap(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		if result := FlattenMap(tc.nested); !reflect.DeepEqual(result, tc.flattened) {
			t.Errorf("Expected %v got %v", tc.flattened, result)
		}
	}
}

func BenchmarkFlattenMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			FlattenMap(tc.nested)
		}
	}
}
