package day132

import "github.com/algds/llrb"

// NewFasterHitCounter returns a new instance of a HitCounter.
func NewFasterHitCounter() HitCounter {
	return &fasterHitCounter{timestamps: llrb.New(func(a, b llrb.Key) int {
		aui := a.(uint64)
		bui := b.(uint64)
		if aui < bui {
			return -1
		} else if aui == bui {
			return 0
		}
		return 1
	})}
}

type fasterHitCounter struct {
	timestamps llrb.Tree
}

func (fhc *fasterHitCounter) Record(ts uint64) {
	llrb.Insert(fhc.timestamps, ts, ts)
}

func (fhc *fasterHitCounter) Total() uint64 {
	return uint64(fhc.timestamps.Size())
}

func (fhc *fasterHitCounter) Range(lower, upper uint64) uint64 {
	return uint64(llrb.RangeCount(fhc.timestamps, lower, upper))
}
