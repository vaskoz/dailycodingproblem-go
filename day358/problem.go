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
	if e1 := bo.lookupEntryByKey[key]; e1 == nil { // entry doesn't exist
		if e2 := bo.lookupListByCount[1]; e2 == nil { // count of 1 doesn't exist either.
			newE := &entry{
				count:       1,
				values:      list.New(),
				lookupByKey: make(map[interface{}]*list.Element),
			}
			newElem := bo.entryList.PushFront(newE)
			bo.lookupEntryByKey[key] = newElem
			bo.lookupListByCount[newE.count] = newElem
			newE.lookupByKey[key] = newE.values.PushFront(key)
		} else { // count 1 does exist for new key
			e2s := e2.Value.(*entry)
			e2s.lookupByKey[key] = e2s.values.PushFront(key)
			bo.lookupEntryByKey[key] = e2
		}
	} else { // entry exists.
		e1s := e1.Value.(*entry)
		if e2 := bo.lookupListByCount[e1s.count+1]; e2 == nil { // count+1 doesn't exist.
			newE := &entry{
				count:       e1s.count + 1,
				values:      list.New(),
				lookupByKey: make(map[interface{}]*list.Element),
			}
			e1key := e1s.lookupByKey[key]
			e1s.values.Remove(e1key)
			newElem := bo.entryList.InsertAfter(newE, e1)
			bo.lookupEntryByKey[key] = newElem
			bo.lookupListByCount[newE.count] = newElem
			newE.lookupByKey[key] = newE.values.PushFront(key)
			if e1s.values.Len() == 0 {
				delete(bo.lookupListByCount, e1s.count)
				bo.entryList.Remove(e1)
			}
		} else { // next count exists just move
			oldE := e1.Value.(*entry)
			oldE1 := oldE.lookupByKey[key]
			oldE.values.Remove(oldE1)
			delete(oldE.lookupByKey, key)
			if oldE.values.Len() == 0 {
				delete(bo.lookupListByCount, oldE.count)
				bo.entryList.Remove(e1)
			}

			newE := e2.Value.(*entry)
			bo.lookupEntryByKey[key] = e2
			newE.lookupByKey[key] = newE.values.PushFront(key)
		}
	}
}

func (bo *bigOhOne) Minus(key interface{}) {
	if e1 := bo.lookupEntryByKey[key]; e1 != nil {
		e1s := e1.Value.(*entry)
		if e1s.count == 1 {
			delete(bo.lookupEntryByKey, key)
			e1skey := e1s.lookupByKey[key]
			e1s.values.Remove(e1skey)

			if e1s.values.Len() == 0 {
				bo.entryList.Remove(e1)
				delete(bo.lookupListByCount, e1s.count)
			}
		} else {
			if e2 := bo.lookupListByCount[e1s.count-1]; e2 == nil { // target doesn't exist
				newE := &entry{
					count:       e1s.count - 1,
					values:      list.New(),
					lookupByKey: make(map[interface{}]*list.Element),
				}
				newE.lookupByKey[key] = newE.values.PushFront(key)
				newElem := bo.entryList.InsertBefore(newE, e1)
				bo.lookupEntryByKey[key] = newElem
				bo.lookupListByCount[newE.count] = newElem
				e1skey := e1s.lookupByKey[key]
				e1s.values.Remove(e1skey)
				delete(e1s.lookupByKey, key)
				if e1s.values.Len() == 0 {
					bo.entryList.Remove(e1)
					delete(bo.lookupListByCount, e1s.count)
				}
			} else { // target does exist.
				e2s := e2.Value.(*entry)
				e2s.lookupByKey[key] = e2s.values.PushFront(key)
				bo.lookupEntryByKey[key] = e2
				e1skey := e1s.lookupByKey[key]
				e1s.values.Remove(e1skey)
				delete(e1s.lookupByKey, key)
				if e1s.values.Len() == 0 {
					bo.entryList.Remove(e1)
					delete(bo.lookupListByCount, e1s.count)
				}
			}
		}
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
