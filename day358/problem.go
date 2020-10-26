package day358

import "container/list"

// BigOhOne is a data structure that performs all operations
// in O(1) time: constant time operations.
type BigOhOne interface {

	// Plus adds a key with value 1.
	// If the key already exists, increment its value by one.
	Plus(key interface{})

	// Minus decrements the value of a key.
	// If the key's value is currently 1, remove it.
	Minus(key interface{})

	// Max returns the key with the highest value.
	Max() (interface{}, int)

	// Min returns the key with the lowest value.
	Min() (interface{}, int)
}

type bigOhOne struct {
	entryList         *list.List
	lookupEntryByKey  map[interface{}]*list.Element
	lookupListByCount map[int]*list.Element
}

type entry struct {
	count       int
	values      *list.List
	lookupByKey map[interface{}]*list.Element
}

func (bo *bigOhOne) Plus(key interface{}) {
	eSource := bo.lookupEntryByKey[key]

	if eSource == nil {
		vals := list.New()
		lookup := make(map[interface{}]*list.Element)
		lookup[key] = vals.PushFront(key)

		newEntry := &entry{
			count:       0,
			values:      vals,
			lookupByKey: lookup,
		}
		eSource = bo.entryList.PushFront(newEntry)
	}

	sourceEntry := eSource.Value.(*entry)

	eTarget := bo.lookupListByCount[sourceEntry.count+1]

	if eTarget == nil {
		newEntry := &entry{
			count:       sourceEntry.count + 1,
			values:      list.New(),
			lookupByKey: make(map[interface{}]*list.Element),
		}
		eTarget = bo.entryList.InsertAfter(newEntry, eSource)
	}

	targetEntry := eTarget.Value.(*entry)

	bo.lookupEntryByKey[key] = eTarget
	bo.lookupListByCount[targetEntry.count] = eTarget
	targetEntry.lookupByKey[key] = targetEntry.values.PushFront(key)

	sourceEntryKey := sourceEntry.lookupByKey[key]
	sourceEntry.values.Remove(sourceEntryKey)
	delete(sourceEntry.lookupByKey, key)

	if len(sourceEntry.lookupByKey) == 0 {
		delete(bo.lookupListByCount, sourceEntry.count)
		bo.entryList.Remove(eSource)
	}
}

func (bo *bigOhOne) Minus(key interface{}) {
	eSource := bo.lookupEntryByKey[key]

	if eSource == nil {
		return
	}

	sourceEntry := eSource.Value.(*entry)

	eTarget := bo.lookupListByCount[sourceEntry.count-1]

	if eTarget == nil {
		newEntry := &entry{
			count:       sourceEntry.count - 1,
			values:      list.New(),
			lookupByKey: make(map[interface{}]*list.Element),
		}
		eTarget = bo.entryList.InsertBefore(newEntry, eSource)
	}

	targetEntry := eTarget.Value.(*entry)

	bo.lookupEntryByKey[key] = eTarget
	bo.lookupListByCount[targetEntry.count] = eTarget
	targetEntry.lookupByKey[key] = targetEntry.values.PushFront(key)

	sourceEntryKey := sourceEntry.lookupByKey[key]
	sourceEntry.values.Remove(sourceEntryKey)
	delete(sourceEntry.lookupByKey, key)

	if len(sourceEntry.lookupByKey) == 0 {
		delete(bo.lookupListByCount, sourceEntry.count)
		bo.entryList.Remove(eSource)
	}

	if targetEntry.count == 0 {
		delete(bo.lookupEntryByKey, key)
		delete(bo.lookupListByCount, targetEntry.count)
		bo.entryList.Remove(eTarget)
	}
}

func (bo *bigOhOne) Max() (interface{}, int) {
	if bo.entryList.Len() == 0 {
		return nil, 0
	}

	back := bo.entryList.Back().Value.(*entry)
	val := back.values.Back().Value

	return val, back.count
}

func (bo *bigOhOne) Min() (interface{}, int) {
	if bo.entryList.Len() == 0 {
		return nil, 0
	}

	front := bo.entryList.Front().Value.(*entry)
	val := front.values.Front().Value

	return val, front.count
}

// New returns a new instance of a BigOhOne data structure.
func New() BigOhOne {
	return &bigOhOne{
		entryList:         list.New(),
		lookupEntryByKey:  make(map[interface{}]*list.Element),
		lookupListByCount: make(map[int]*list.Element),
	}
}
