package day132

// HitCounter implements 3 methods.
type HitCounter interface {
	// Record stores a timestamp.
	Record(uint64)
	// Total returns number of recorded events.
	Total() int
	// Range returns the total counts inclusive.
	Range(lower, upper uint64) int
}

// NewHitCounter returns a new instance of a HitCounter.
func NewHitCounter() HitCounter {
	return &hitCounter{}
}

type hitCounter struct {
	total      int
	timestamps []uint64
}

func (hc *hitCounter) Record(ts uint64) {
	hc.total++
	hc.timestamps = append(hc.timestamps, ts)
}

func (hc *hitCounter) Total() int {
	return hc.total
}

func (hc *hitCounter) Range(lower, upper uint64) int {
	var count int
	for _, ts := range hc.timestamps {
		if ts >= lower && ts <= upper {
			count++
		}
	}
	return count
}
