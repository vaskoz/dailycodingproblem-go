package day368

import "errors"

// KeyValueStore is a key value store for the problem.
type KeyValueStore interface {
	// Update updates the value at key to val, or sets it if doesn't exist
	Update(key, val int)
	// Get returns the value with key or an error if no such value exists
	Get(key int) (int, error)
	// MaxKey returns the largest key with value val or error
	// if no key with that value exists
	MaxKey(val int) (int, error)
}

var errNotFound = errors.New("not found")

type keyValueStore struct {
	keyToVals map[int]int
	valToKeys map[int]map[int]struct{}
}

func (kvs *keyValueStore) Update(k, v int) {
	if prevV, found := kvs.keyToVals[k]; found {
		delete(kvs.valToKeys[prevV], k) // delete previous key

		if len(kvs.valToKeys[prevV]) == 0 { // delete previous value
			delete(kvs.valToKeys, prevV)
		}
	}

	kvs.keyToVals[k] = v
	if m, exists := kvs.valToKeys[v]; exists {
		m[k] = struct{}{}
	} else {
		kvs.valToKeys[v] = make(map[int]struct{})
		kvs.valToKeys[v][k] = struct{}{}
	}
}

func (kvs *keyValueStore) Get(k int) (int, error) {
	if v, found := kvs.keyToVals[k]; found {
		return v, nil
	}

	return 0, errNotFound
}

func (kvs *keyValueStore) MaxKey(v int) (int, error) {
	if len(kvs.valToKeys[v]) == 0 {
		return 0, errNotFound
	}

	maxKey := -int(^uint(0)>>1) - 1 // most negative int
	for k := range kvs.valToKeys[v] {
		if k > maxKey {
			maxKey = k
		}
	}

	return maxKey, nil
}

// NewKeyValueStore returns a new instance of a KeyValueStore.
func NewKeyValueStore() KeyValueStore {
	return &keyValueStore{
		keyToVals: make(map[int]int),
		valToKeys: make(map[int]map[int]struct{}),
	}
}
