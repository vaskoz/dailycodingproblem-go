package day242

import "sync"

// DailySubscriber is an interface that describes the two actions.
type DailySubscriber interface {
	Update(hour, value int)
	Query(start, end int) int
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
