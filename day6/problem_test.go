package day6

import (
	"reflect"
	"testing"
	"unsafe"
)

// XorList is an xor doubly linked list
type XorList struct {
	val  interface{}
	both uintptr // prev xor next
}

// New returns a pointer to an XorList
func New(val interface{}) *XorList {
	return &XorList{val: val, both: 0}
}

// Add adds a new value
func (xl *XorList) Add(val interface{}) {
	prev := uintptr(0)
	this := unsafe.Pointer(xl) // nolint: gosec

	for prev^xl.both != 0 {
		prev, this = uintptr(this), unsafe.Pointer(prev^xl.both) // nolint: vet, gosec
		xl = (*XorList)(this)
	}

	nxl := &XorList{val: val, both: uintptr(this)}
	xl.both = prev ^ uintptr(unsafe.Pointer(nxl)) // nolint: gosec
}

// nolint
var testcases = []struct {
	data []interface{}
}{
	{[]interface{}{6, 10, 30, 40, 100, 101, 102}},
	{[]interface{}{102, 101, 33, 44, 101, 25, 29, 50, 76}},
	{[]interface{}{102, 101, 100, 99, 88, 77, 66, 55, 44, 33, 22, 11}},
}

func TestXorList(t *testing.T) {
	t.Parallel()
	t.Skip("No longer works in Go 1.13 due to GC memory safety")

	for _, tc := range testcases {
		xl := New(tc.data[0])

		for i := 1; i < len(tc.data); i++ {
			xl.Add(tc.data[i])
		}

		if result := forwardToSlice(xl); !reflect.DeepEqual(result, tc.data) {
			t.Errorf("Expected %v but got %v", tc.data, result)
		}

		if reversed, result := reverse(tc.data), backwardToSlice(xl); !reflect.DeepEqual(result, reversed) {
			t.Errorf("Expected %v but got %v", reversed, result)
		}
	}
}

func BenchmarkXorList(b *testing.B) {
	b.Skip("No longer works in Go 1.13 due to GC memory safety")

	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			xl := New(tc.data[0])

			for i := 1; i < len(tc.data); i++ {
				xl.Add(tc.data[i])
			}
		}
	}
}

// reverses a slice of elements
func reverse(data []interface{}) []interface{} {
	reversed := make([]interface{}, len(data))
	for i, v := range data {
		reversed[len(data)-i-1] = v
	}

	return reversed
}

// backwardToSlice is a testing utility to iterate through this list and construct a slice.
// it's easier to assert on.
func backwardToSlice(xl *XorList) []interface{} {
	// go forward from the head to the tail
	prev := uintptr(0)
	this := unsafe.Pointer(xl)

	for prev^xl.both != 0 {
		prev, this = uintptr(this), unsafe.Pointer(prev^xl.both) // nolint: vet
		xl = (*XorList)(this)
	}
	// now collect items from tail to the head
	var result []interface{}

	prev = uintptr(0)
	this = unsafe.Pointer(xl)
	result = append(result, xl.val)

	for prev^xl.both != 0 {
		prev, this = uintptr(this), unsafe.Pointer(prev^xl.both) // nolint: vet
		xl = (*XorList)(this)
		result = append(result, xl.val)
	}

	return result
}

// forwardToSlice is a testing utility to iterate through this list and construct a slice.
// it's easier to assert on.
func forwardToSlice(xl *XorList) []interface{} {
	var result []interface{}

	prev := uintptr(0)
	this := unsafe.Pointer(xl)
	result = append(result, xl.val)

	for prev^xl.both != 0 {
		prev, this = uintptr(this), unsafe.Pointer(prev^xl.both) // nolint: vet
		xl = (*XorList)(this)
		result = append(result, xl.val)
	}

	return result
}
