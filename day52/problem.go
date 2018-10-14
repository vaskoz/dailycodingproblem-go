package day52

import (
	"container/list"
)

// Cache specifies a generic cache interface of key-value pairs.
type Cache interface {
	Set(key, value interface{})
	Get(key interface{}) interface{}
}

type kvPair struct {
	k, v interface{}
}

type lruCache struct {
	data     map[interface{}]*list.Element
	l        *list.List
	capacity int
}

func (lc *lruCache) Set(key interface{}, value interface{}) {
	if elem, found := lc.data[key]; found {
		lc.l.MoveToFront(elem)
		elem.Value = kvPair{k: key, v: value}
		return
	}
	if len(lc.data) == lc.capacity { // evict oldest
		oldest := lc.l.Back()
		delete(lc.data, oldest.Value.(kvPair).k)
		lc.l.Remove(oldest)
	}
	e := lc.l.PushFront(kvPair{k: key, v: value})
	lc.data[key] = e
}

func (lc *lruCache) Get(key interface{}) interface{} {
	if elem, found := lc.data[key]; found {
		lc.l.MoveToFront(elem)
		return elem.Value.(kvPair).v
	}
	return nil
}

// NewLRU returns a Cache implementation that behaves as a LRU.
func NewLRU(capacity int) Cache {
	return &lruCache{make(map[interface{}]*list.Element), list.New(), capacity}
}
