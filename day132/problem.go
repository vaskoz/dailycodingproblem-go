package day132

// HitCounter implements 3 methods.
type HitCounter interface {
	// Record stores a timestamp.
	Record(uint64)
	// Total returns number of recorded events.
	Total() uint64
	// Range returns the total counts inclusive.
	Range(lower, upper uint64) uint64
}

// NewHitCounter returns a new instance of a HitCounter.
func NewHitCounter() HitCounter {
	return &hitCounter{}
}

type hitCounter struct {
	total      uint64
	timestamps []uint64
}

func (hc *hitCounter) Record(ts uint64) {
	hc.total++
	hc.timestamps = append(hc.timestamps, ts)
}

func (hc *hitCounter) Total() uint64 {
	return hc.total
}

func (hc *hitCounter) Range(lower, upper uint64) uint64 {
	var count uint64
	for _, ts := range hc.timestamps {
		if ts >= lower && ts <= upper {
			count++
		}
	}
	return count
}
