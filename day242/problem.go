package day242

import "sync"

// DailySubscriber is an interface that describes the two actions.
type DailySubscriber interface {
	Update(hour, value int)
	Query(start, end int) int
}

type readOptimized struct {
	sync.RWMutex
	data [24]int
}

// NewReadOptimized returns a read-optimized version.
func NewReadOptimized(data [24]int) DailySubscriber {
	sum := data[0]
	for i := 1; i < 24; i++ {
		sum += data[i]
		data[i] = sum
	}
	return &readOptimized{data: data}
}

func (r *readOptimized) Update(hour, value int) {
	if hour < 0 || hour > 23 {
		panic("invalid hour")
	}
	r.Lock()
	for i := hour; i < 24; i++ {
		r.data[i] += value
	}
	r.Unlock()
}

func (r *readOptimized) Query(start, end int) int {
	if start < 0 || start > 23 || end < 0 || end > 23 {
		panic("start and end must be between 0-23 inclusive")
	}
	if start > end {
		panic("start must preceed or equal end")
	}
	r.RLock()
	var sum int
	if start == 0 {
		sum = r.data[end]
	} else {
		sum = r.data[end] - r.data[start-1]
	}
	r.RUnlock()
	return sum
}

type writeOptimized struct {
	sync.RWMutex
	data [24]int
}

// NewWriteOptimized returns a write-optimized version.
func NewWriteOptimized(data [24]int) DailySubscriber {
	return &writeOptimized{data: data}
}

func (w *writeOptimized) Update(hour, value int) {
	if hour < 0 || hour > 23 {
		panic("invalid hour")
	}
	w.Lock()
	w.data[hour] += value
	w.Unlock()
}

func (w *writeOptimized) Query(start, end int) int {
	if start < 0 || start > 23 || end < 0 || end > 23 {
		panic("start and end must be between 0-23 inclusive")
	}
	if start > end {
		panic("start must preceed or equal end")
	}
	w.RLock()
	var sum int
	for i := start; i <= end; i++ {
		sum += w.data[i]
	}
	w.RUnlock()
	return sum
}
