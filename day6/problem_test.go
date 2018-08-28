package day6

import (
	"reflect"
	"testing"
	"unsafe"
)

var testcases = []struct {
	data []interface{}
}{
	{[]interface{}{6, 10, 30, 40, 100, 101, 102}},
	{[]interface{}{102, 101, 33, 44, 101, 25, 29, 50, 76}},
	{[]interface{}{102, 101, 100, 99, 88, 77, 66, 55, 44, 33, 22, 11}},
}

func TestXorList(t *testing.T) {
	for _, tc := range testcases {
		xl := New(tc.data[0])
		for i := 1; i < len(tc.data); i++ {
			xl.Add(tc.data[i])
		}
		if result := forwardToSlice(xl); !reflect.DeepEqual(result, []interface{}(tc.data)) {
			t.Errorf("Expected %v but got %v", tc.data, result)
		}
	}
}

func BenchmarkXorList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			xl := New(tc.data[0])
			for i := 1; i < len(tc.data); i++ {
				xl.Add(tc.data[i])
			}
		}
	}
}

// forwardToSlice is a testing utility to iterate through this list and construct a slice.
// it's easier to assert on.
func forwardToSlice(xl *XorList) []interface{} {
	var result []interface{}
	prev := uintptr(0)
	this := unsafe.Pointer(xl)
	result = append(result, xl.val)
	for prev^xl.both != 0 {
		prev, this = uintptr(this), unsafe.Pointer(prev^xl.both)
		xl = (*XorList)(this)
		result = append(result, xl.val)
	}
	return result
}
