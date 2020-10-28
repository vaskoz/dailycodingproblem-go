package day358

import "github.com/algds/minmaxcnt"

type Wrapper struct {
	delegate minmaxcnt.Interface
}

func (w Wrapper) Plus(key interface{}) {
	w.delegate.Increment(key)
}

func (w Wrapper) Minus(key interface{}) {
	w.delegate.Decrement(key)
}

func (w Wrapper) Max() (interface{}, int) {
	val, count := w.delegate.Max()

	return val, int(count)
}

func (w Wrapper) Min() (interface{}, int) {
	val, count := w.delegate.Min()

	return val, int(count)
}

// New returns a new instance of a Wrapper data structure.
func New() Wrapper {
	return Wrapper{
		minmaxcnt.New(),
	}
}
