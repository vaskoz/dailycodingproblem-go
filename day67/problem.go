package day67

import (
	"container/list"
)

type freqNode struct {
	accessCount int
	siblings    *list.List // is a list of keyValueNodes
}

type keyValueNode struct {
	key      interface{}
	value    interface{}
	parent   *list.List
	freqElem *list.Element
}

// LFUCache is a least-frequently-used cache.
// Every operation runs in O(1) time.
type LFUCache struct {
	freqList *list.List // is a list of freqNodes
	lookup   map[interface{}]*list.Element
	capacity int
}

// Set sets key to value. If there are already n items in the cache
// and we are adding a new item, then it should also remove the least frequently used item.
// If there is a tie, then the least recently used key should be removed.
func (lfu *LFUCache) Set(key, value interface{}) {
	if keyValNode, found := lfu.lookup[key]; found {
		// update existing
		kvn := keyValNode.Value.(*keyValueNode)
		kvn.parent.Remove(keyValNode)
		kvn.value = value
		fqnNext := kvn.freqElem.Next()
		fqn := kvn.freqElem.Value.(*freqNode)
		if fqnNext == nil || fqnNext.Value.(*freqNode).accessCount != fqn.accessCount+1 {
			newFreq := &freqNode{accessCount: fqn.accessCount + 1, siblings: list.New()}
			e := newFreq.siblings.PushBack(kvn)
			lfu.lookup[key] = e
			lfu.freqList.InsertAfter(newFreq, kvn.freqElem)
		} else {
			e := fqnNext.Value.(*freqNode).siblings.PushBack(kvn)
			lfu.lookup[key] = e
		}
		if kvn.parent.Len() == 0 {
			lfu.freqList.Remove(kvn.freqElem)
		}
		return
	} else if len(lfu.lookup) == lfu.capacity {
		// evict the LFU first.
		lowestAccess := lfu.freqList.Front().Value.(*freqNode)
		leastRecently := lowestAccess.siblings.Front()
		kv := leastRecently.Value.(*keyValueNode)
		delete(lfu.lookup, kv.key)
		lowestAccess.siblings.Remove(leastRecently)
		if lowestAccess.siblings.Len() == 0 {
			lfu.freqList.Remove(lfu.freqList.Front())
		}
	}
	// add a new value into the cache.
	if front := lfu.freqList.Front(); front == nil {
		fqn := &freqNode{accessCount: 1, siblings: list.New()}
		lfu.freqList.PushFront(fqn)
		fqnElem := lfu.freqList.Front()
		kvn := &keyValueNode{key: key, value: value, parent: fqn.siblings, freqElem: fqnElem}
		e := fqn.siblings.PushBack(kvn)
		lfu.lookup[key] = e
	} else {
		fqn := front.Value.(*freqNode)
		if fqn.accessCount == 1 {
			kvn := &keyValueNode{key: key, value: value, parent: fqn.siblings, freqElem: front}
			e := fqn.siblings.PushBack(kvn)
			lfu.lookup[key] = e
		} else {
			fqn := &freqNode{accessCount: 1, siblings: list.New()}
			lfu.freqList.PushFront(fqn)
			fqnElem := lfu.freqList.Front()
			kvn := &keyValueNode{key: key, value: value, parent: fqn.siblings, freqElem: fqnElem}
			e := fqn.siblings.PushBack(kvn)
			lfu.lookup[key] = e
		}
	}
}

// Get gets the value at key. If no such key exists, return null.
func (lfu *LFUCache) Get(key interface{}) interface{} {
	// TODO: update freq counts.
	if _, found := lfu.lookup[key]; !found {
		return nil
	}
	keyValNode := lfu.lookup[key]
	kvn := keyValNode.Value.(*keyValueNode)
	kvn.parent.Remove(keyValNode)
	fqnNext := kvn.freqElem.Next()
	fqn := kvn.freqElem.Value.(*freqNode)
	if fqnNext == nil || fqnNext.Value.(*freqNode).accessCount != fqn.accessCount+1 {
		newFreq := &freqNode{accessCount: fqn.accessCount + 1, siblings: list.New()}
		e := newFreq.siblings.PushBack(kvn)
		lfu.lookup[key] = e
		lfu.freqList.InsertAfter(newFreq, kvn.freqElem)
	} else {
		e := fqnNext.Value.(*freqNode).siblings.PushBack(kvn)
		lfu.lookup[key] = e
	}
	if kvn.parent.Len() == 0 {
		lfu.freqList.Remove(kvn.freqElem)
	}
	return kvn.value
}

// NewLFUCache returns a new instance of LFUCache.
func NewLFUCache(n int) *LFUCache {
	return &LFUCache{capacity: n,
		lookup:   make(map[interface{}]*list.Element), // list.Element is *keyValueNode
		freqList: list.New()}
}
