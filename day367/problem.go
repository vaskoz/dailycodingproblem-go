package day367

import "sort"

// LessThanEqual returns true if the first argument is less or equal to the second.
// False otherwise.
type LessThanEqual func(a, b interface{}) bool

// SortedIterator is an iterator with sorted elements.
type SortedIterator interface {
	// Empty is true if there are no more elements in the iterator.
	Empty() bool

	// Peek returns the next value without advancing the iterator.
	Peek() interface{}

	// Next returns the next sorted iterator element.
	// Returns nil if Empty() is true.
	Next() interface{}
}

type singleSourceSortedIterator struct {
	sortedData []interface{}
	nextIndex  int
}

func (sssi *singleSourceSortedIterator) Empty() bool {
	return sssi.nextIndex == len(sssi.sortedData)
}

func (sssi *singleSourceSortedIterator) Peek() interface{} {
	if sssi.nextIndex == len(sssi.sortedData) {
		return nil
	}

	return sssi.sortedData[sssi.nextIndex]
}

func (sssi *singleSourceSortedIterator) Next() interface{} {
	if sssi.nextIndex == len(sssi.sortedData) {
		return nil
	}
	sssi.nextIndex++

	return sssi.sortedData[sssi.nextIndex-1]
}

// NewSortedIterator returns a new SortedIterator given a slice and a lessThan
// function.
// "lessThan" should answer consistently if 'a' is less than 'b' then true, else false.
func NewSortedIterator(vals []interface{}, lessThan LessThanEqual) SortedIterator {
	copied := make([]interface{}, len(vals))
	copy(copied, vals)
	sort.Slice(copied, func(i, j int) bool {
		return lessThan(copied[i], copied[j])
	})

	return &singleSourceSortedIterator{sortedData: copied, nextIndex: 0}
}

type multipleSourceSortedIterator struct {
	a, b     SortedIterator
	lessThan func(a, b interface{}) bool
}

func (mssi *multipleSourceSortedIterator) Empty() bool {
	return mssi.a.Empty() && mssi.b.Empty()
}

func (mssi *multipleSourceSortedIterator) Peek() interface{} {
	switch {
	case mssi.Empty():
		return nil
	case mssi.a.Empty():
		return mssi.b.Peek()
	case mssi.b.Empty():
		return mssi.a.Peek()
	default:
		if mssi.lessThan(mssi.a.Peek(), mssi.b.Peek()) {
			return mssi.a.Peek()
		}

		return mssi.b.Peek()
	}
}

func (mssi *multipleSourceSortedIterator) Next() interface{} {
	switch {
	case mssi.Empty():
		return nil
	case mssi.a.Empty():
		return mssi.b.Next()
	case mssi.b.Empty():
		return mssi.a.Next()
	default:
		if mssi.lessThan(mssi.a.Peek(), mssi.b.Peek()) {
			return mssi.a.Next()
		}

		return mssi.b.Next()
	}
}

// MergeSortedIterators returns a SortedIterator that is the result
// of merging two other SortedIterator's.
func MergeSortedIterators(a, b SortedIterator, lessThan LessThanEqual) SortedIterator {
	return &multipleSourceSortedIterator{a, b, lessThan}
}
