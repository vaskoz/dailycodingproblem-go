package day6

import (
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
