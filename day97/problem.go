package day97

import (
	"sort"
)

// TimestampedValue represents the value set at a particular time.
type TimestampedValue struct {
	Time  int
	Value interface{}
}

// TimeMap is a map that is aware of changes over time.
type TimeMap struct {
	data map[interface{}][]*TimestampedValue
}

// NewTimeMap returns a new timestamp aware map.
func NewTimeMap() *TimeMap {
	return &TimeMap{make(map[interface{}][]*TimestampedValue)}
}

// Set a key-value pair with a particular timestamp.
func (tm *TimeMap) Set(key, value interface{}, time int) {
	tm.data[key] = append(tm.data[key], &TimestampedValue{time, value})
	sort.Slice(tm.data[key], func(i, j int) bool {
		return tm.data[key][i].Time < tm.data[key][j].Time
	})
}

// Get returns the value for a key at a particular time.
func (tm *TimeMap) Get(key interface{}, time int) interface{} {
	pos := sort.Search(len(tm.data[key]), func(i int) bool {
		return tm.data[key][i].Time > time
	})
	if pos <= 0 {
		if len(tm.data[key]) == 0 || tm.data[key][0].Time > time {
			return nil
		}
	}
	return tm.data[key][pos-1].Value
}
