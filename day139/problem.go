package day139

// Iterator is a simple iterator interface.
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// PeekableIterator wraps an iterator and provides the ability to peek.
type PeekableIterator struct {
	peek   interface{}
	peeked bool
	Iter   Iterator
}

// Peek returns the next element without moving the iterator.
func (p *PeekableIterator) Peek() interface{} {
	has := p.Iter.HasNext()
	if !p.peeked && !has {
		return nil
	} else if !p.peeked {
		p.peeked = true
		p.peek = p.Iter.Next()
	}
	return p.peek
}

// Next returns the next element and advances the iterator.
func (p *PeekableIterator) Next() interface{} {
	has := p.Iter.HasNext()
	var result interface{}
	if p.peeked {
		result = p.peek
	} else if has {
		result = p.Iter.Next()
	}
	p.peeked = false
	return result
}

// HasNext returns true if there is more to read from the iterator
// and false otherwise.
func (p *PeekableIterator) HasNext() bool {
	has := p.Iter.HasNext()
	if has || p.peeked {
		return true
	}
	return false
}
